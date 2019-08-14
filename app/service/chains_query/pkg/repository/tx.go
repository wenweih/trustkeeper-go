package repository


import (
  "fmt"
  "bytes"
  "strconv"
  "math/big"
  "context"
  "encoding/hex"
  "github.com/btcsuite/btcutil"
  "github.com/btcsuite/btcd/chaincfg"
  "github.com/btcsuite/btcutil/coinset"
  "github.com/btcsuite/btcd/wire"
  "github.com/btcsuite/btcd/txscript"
  "github.com/btcsuite/btcd/mempool"
  "trustkeeper-go/app/service/chains_query/pkg/model"
)

func (repo *repo) ConstructTxBTC(ctx context.Context, from, to, amount string) (unsignedTxHex string, err error) {
  fromAddress, err := btcutil.DecodeAddress(from, &chaincfg.RegressionNetParams)
  if err != nil {
    return "", fmt.Errorf("Invalid From address:" + err.Error())
  }
  fromPkScript, err := txscript.PayToAddrScript(fromAddress)
  if err != nil {
    return "", err
  }

  toAddress, err := btcutil.DecodeAddress(from, &chaincfg.RegressionNetParams)
  if err != nil {
    return "", fmt.Errorf("Invalid To address:" + err.Error())
  }
  toPkScript, err := txscript.PayToAddrScript(toAddress)
  if err != nil {
    return "", err
  }

  balance := model.Balance{}
  if err := repo.db.Where("address = ? AND symbol = ?", from, model.BTCSymbol).First(&balance).Error; err != nil {
    return "", err
  }

  balanceAmount, result := new(big.Int).SetString(balance.Amount, 10)
  if !result {
    return "", fmt.Errorf("fail to extract amount from balance")
  }
  amountFloat, err := strconv.ParseFloat(amount, 64)
  if err != nil {
    return "", fmt.Errorf("fail to parse float %s", err.Error())
  }
  transferAmount, err := btcutil.NewAmount(amountFloat)
  if err != nil {
    return "", fmt.Errorf("Transfer amount %s", err.Error())
  }
  transferAmountBig := big.NewInt(int64(transferAmount))
  if balanceAmount.Cmp(transferAmountBig) <= 0 {
    return "", fmt.Errorf("Insufficient balance")
  }
  var txids []string
  if err := repo.db.Model(&model.Tx{State: model.StateSuccess,
    ChainName: model.ChainBitcoin,
    BalanceID: balance.ID,
    Asset: model.BTCSymbol}).Pluck("tx_id", &txids).Error;
    err != nil {
      return "", fmt.Errorf("Fail to query comfirmations btc tx:" + err.Error())
    }
  var matureUTXOs []model.BtcUtxo
  if err := repo.db.Where("txid IN (?) AND balance_id = ?", txids, balance.ID).Find(&matureUTXOs).Error; err != nil{
    return "", fmt.Errorf("fail to query mature utxos:" + err.Error())
  }

  ledgerInfo, err := repo.QueryBTCLedgerInfo(ctx)
  if err != nil {
    return "", err
  }
  feeKB, err := repo.bitcoinClient.EstimateSmartFee(int64(6))
  if err != nil {
    return "", err
  }
  feeRate := mempool.SatoshiPerByte(feeKB.FeeRate)

  if feeKB.FeeRate <= 0 {
    feeRate = mempool.SatoshiPerByte(100)
  }

  selectedutxos, unselectedutxos, selectedCoins, err := CoinSelect(int64(ledgerInfo.Headers), btcutil.Amount(transferAmountBig.Int64()), matureUTXOs)
  if err != nil {
    return "", fmt.Errorf("fail to select UTXO %s", err.Error())
  }
  var vinAmount int64
  for _, coin := range selectedCoins.Coins() {
    vinAmount += int64(coin.Value())
  }
  msgTx := coinset.NewMsgTxWithInputCoins(wire.TxVersion, selectedCoins)
  txOutTo := wire.NewTxOut(int64(transferAmountBig.Int64()), toPkScript)
  msgTx.AddTxOut(txOutTo)

  // recharge
  // 181, 34: https://bitcoin.stackexchange.com/questions/1195/how-to-calculate-transaction-size-before-sending-legacy-non-segwit-p2pkh-p2sh
  fee := feeRate.Fee(uint32(msgTx.SerializeSize() + 181 + 34))
  if (vinAmount - transferAmountBig.Int64() - int64(fee)) > 0 {
    txOutReCharge := wire.NewTxOut((vinAmount - transferAmountBig.Int64() - int64(fee)), fromPkScript)
    msgTx.AddTxOut(txOutReCharge)
  }else {
    selectedutxoForFee, _, selectedCoinsForFee, err := CoinSelect(int64(ledgerInfo.Headers), fee, unselectedutxos)
    if err != nil {
      return "", fmt.Errorf("Select UTXO for fee %s", err)
    }
    for _, coin := range selectedCoinsForFee.Coins() {
      vinAmount += int64(coin.Value())
    }
    txOutReCharge := wire.NewTxOut((vinAmount - transferAmountBig.Int64() - int64(fee)), fromPkScript)
    msgTx.AddTxOut(txOutReCharge)
    selectedutxos = append(selectedutxos, selectedutxoForFee...)
  }

  buf := bytes.NewBuffer(make([]byte, 0, msgTx.SerializeSize()))
  msgTx.Serialize(buf)
  rawTxHex := hex.EncodeToString(buf.Bytes())
  // c.Wallet.SelectedUTXO = selectedutxos

  return rawTxHex, nil
}
