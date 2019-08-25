package repository

import(
  "fmt"
  "strconv"
  "reflect"
  "context"
  "strings"
  "math/big"
  "github.com/shopspring/decimal"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ybbus/jsonrpc"
  "trustkeeper-go/app/service/chains_query/pkg/model"
  "github.com/ethereum/go-ethereum/core/types"
)

func (repo *repo) ConstructTxETH(ctx context.Context, from, to, amount string) (string, string, error) {
  // withdraw and deposit address validate
  if !common.IsHexAddress(from) {
    return "", "", fmt.Errorf("Invalid address: %s", from)
  }
  if !common.IsHexAddress(to) {
    return "", "", fmt.Errorf("Invalid address: %s", to)
  }

  // balance data
  balance := model.Balance{}
  if err := repo.db.Where("address = ? AND symbol = ?", from, model.ETHSymbol).First(&balance).Error; err != nil {
    return "", "", fmt.Errorf("Fail to query balance %s", err.Error())
  }
  balanceAmount, result := new(big.Int).SetString(balance.Amount, 10)
  if !result {
    return "", "", fmt.Errorf("fail to extract amount from balance")
  }
  withdrawLockAmount, result := new(big.Int).SetString(balance.WithdrawLock, 10)
  if !result {
    return "", "", fmt.Errorf("fail to extract withdrawLock from balance")
  }
  balanceAmount = balanceAmount.Sub(balanceAmount, withdrawLockAmount)
  balanceAmountDecimal := decimal.NewFromBigInt(balanceAmount, 0)

  // transfer amount data
  transferAmountDecimal, err :=  decimal.NewFromString(amount)
  if err != nil {
    return "", "", fmt.Errorf("Fail to extract transfer amount %s", err.Error())
  }
  decimalBig, result := new(big.Int).SetString(strconv.FormatUint(balance.Decimal, 10), 10)
  if !result {
    return "", "", fmt.Errorf("Fail to convert decimal")
  }
  transferAmountDecimal = transferAmountDecimal.Mul(decimal.NewFromBigInt(decimalBig, 0))
  value, ok := new(big.Int).SetString(transferAmountDecimal.String(), 10)
  if !ok {
    return "", "", fmt.Errorf("Set amount error")
  }

  // check balance is whether GreaterThan transfer amount
  if !balanceAmountDecimal.GreaterThan(transferAmountDecimal) {
    return "", "", fmt.Errorf("Insufficient balance")
  }

  // gas price
  gasPrice, err := repo.ethClient.SuggestGasPrice(ctx)
  if err != nil {
    return "", "", fmt.Errorf("Fail to query gas price %s", err.Error())
  }
  // tx fee
  var (
    data []byte
    txPoolInspect *model.TxPoolInspect
    txPoolMaxCount uint64
  )
  txFee := new(big.Int)
  gasLimit := uint64(21000) // in units
  txFee = txFee.Mul(gasPrice, big.NewInt(int64(gasLimit)))
  feeDecimal, _ := decimal.NewFromString(txFee.String())

  // ETH transfer
  // if totalCost > balance then return
  totalCost := transferAmountDecimal.Add(feeDecimal)
  if balanceAmountDecimal.LessThan(totalCost) {
    totalCostBig, _ := new(big.Int).SetString(totalCost.String(), 10)
    return "", "", fmt.Errorf("Insufficient ETH balance %s : %s",
      model.ToEther(balanceAmount).String(),
      model.ToEther(totalCostBig).String(),
    )
  }

  // pendingNonceAt account
  pendingNonaceAt, err := repo.ethClient.PendingNonceAt(ctx, common.HexToAddress(from))
  if err != nil {
    return "", "", fmt.Errorf("Fail to query pending nonce for address %s", err.Error())
  }
  // get real nonce in mempool
  rpcClient := jsonrpc.NewClient(repo.conf.ETHRPC)
  response, err := rpcClient.Call("txpool_inspect")
  if err != nil {
    return "", "", fmt.Errorf("Fail to query txpool inspect %s", err.Error())
  }
  if response.Error != nil {
    return "", "", response.Error
  }

  // nonce for account
  if err = response.GetObject(&txPoolInspect); err != nil {
    return "", "", err
  }
  pending := reflect.ValueOf(txPoolInspect.Pending)
  if pending.Kind() == reflect.Map {
    for _, key := range pending.MapKeys() {
      address := key.Interface().(string)
      tx := reflect.ValueOf(pending.MapIndex(key).Interface())
      if tx.Kind() == reflect.Map && strings.ToLower(from) == strings.ToLower(address){
        for _, key := range tx.MapKeys() {
          count := key.Interface().(uint64)
          if count > txPoolMaxCount {
            txPoolMaxCount = count
          }
        }
      }
    }
  }
  pendingNonce := pendingNonaceAt
  if pendingNonaceAt !=0 && txPoolMaxCount + 1 > pendingNonaceAt {
    pendingNonce = txPoolMaxCount + 1
  }

  tx := types.NewTransaction(pendingNonce, common.HexToAddress(to), value, gasLimit, gasPrice, data)
  rawTxHex, err := model.EncodeETHTx(tx)
  if err != nil {
    return "", "", fmt.Errorf("Encode raw tx %s", err)
  }

  chainID, err :=  repo.ethClient.ChainID(ctx)
  if err != nil {
    return "", "", err
  }
  if err := repo.db.Model(&balance).
  UpdateColumn("withdraw_lock", decimal.NewFromBigInt(withdrawLockAmount, 0).Add(decimal.NewFromBigInt(value, 0)).String()).Error;
  err != nil {
    return "", "", fmt.Errorf("Fail to update withdrawLock row %s", err.Error())
  }
  return rawTxHex, chainID.String(), nil
}
