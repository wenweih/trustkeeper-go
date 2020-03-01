package repository

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"trustkeeper-go/app/service/chains_query/pkg/model"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/ybbus/jsonrpc"
	"golang.org/x/crypto/sha3"
)

func (repo *repo) ConstructTxETH(ctx context.Context, from, to, amount string) (string, string, error) {
	balanceAmountDecimal, transferAmountDecimal, withdrawLockAmountDecimal, gasPrice, value, balance, err := repo.constructTxParamsValidate(ctx, from, to, amount, model.ETHSymbol)
	if err != nil {
		return "", "", err
	}
	// tx fee
	var (
		data           []byte
		txPoolInspect  *model.TxPoolInspect
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
			model.ToEther(balanceAmountDecimal.Coefficient()).String(),
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
			if tx.Kind() == reflect.Map && strings.ToLower(from) == strings.ToLower(address) {
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
	if pendingNonaceAt != 0 && txPoolMaxCount+1 > pendingNonaceAt {
		pendingNonce = txPoolMaxCount + 1
	}

	tx := types.NewTransaction(pendingNonce, common.HexToAddress(to), value, gasLimit, gasPrice, data)
	rawTxHex, err := model.EncodeETHTx(tx)
	if err != nil {
		return "", "", fmt.Errorf("Encode raw tx %s", err)
	}

	chainID, err := repo.ethClient.ChainID(ctx)
	if err != nil {
		return "", "", err
	}
	if err := repo.db.Model(&balance).
		UpdateColumn("withdraw_lock", withdrawLockAmountDecimal.Add(decimal.NewFromBigInt(value, 0)).String()).Error; err != nil {
		return "", "", fmt.Errorf("Fail to update withdrawLock row %s", err.Error())
	}
	return rawTxHex, chainID.String(), nil
}

func (repo *repo) ConstructTxERC20(ctx context.Context, from, to, amount, contract string) (unsignedTxHex, chainID string, err error) {
	_, transferAmountDecimal, withdrawLockAmountDecimal, gasPrice, value, tokenBalance, err := repo.constructTxParamsValidate(ctx, from, to, amount, contract)
	if err != nil {
		return "", "", err
	}
	// ethbalance data
	ethBalance := model.Balance{}
	err = repo.db.Where("address = ? AND symbol = ?", from, model.ETHSymbol).First(&ethBalance).Error
	if err != nil {
		return "", "", fmt.Errorf("Fail to query eth balance %s", err.Error())
	}
	ethAmountDecimal, err := decimal.NewFromString(ethBalance.Amount)
	if err != nil {
		return "", "", fmt.Errorf("Fail to extract fee %s", err.Error())
	}

	// construct data
	transferFunSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFunSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(common.HexToAddress(to).Bytes(), 32)
	tokenAmount, ok := new(big.Int).SetString(transferAmountDecimal.String(), 10)
	if !ok {
		return "", "", fmt.Errorf("Set amount error")
	}
	paddedAmount := common.LeftPadBytes(tokenAmount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	tokenAddress := common.HexToAddress(tokenBalance.Identify)
	gasLimit, err := repo.ethClient.EstimateGas(ctx, ethereum.CallMsg{
		From: common.HexToAddress(from), // https://github.com/paritytech/parity-ethereum/issues/10147#issuecomment-462177568
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		return "", "", fmt.Errorf("EstimateGas %s", err)
	}
	value = big.NewInt(0)
	// tx fee
	txFee := new(big.Int)
	txFee = txFee.Mul(gasPrice, big.NewInt(int64(gasLimit)))
	feeDecimal, err := decimal.NewFromString(txFee.String())
	if err != nil {
		return "", "", fmt.Errorf("Fail to extract fee %s", err.Error())
	}
	if ethAmountDecimal.LessThan(feeDecimal) {
		return "", "", fmt.Errorf("Insufficient ETH balance for fee %s : %s", ethAmountDecimal.String(), feeDecimal.String())
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
	var (
		txPoolInspect  *model.TxPoolInspect
		txPoolMaxCount uint64
	)
	if err = response.GetObject(&txPoolInspect); err != nil {
		return "", "", err
	}
	pending := reflect.ValueOf(txPoolInspect.Pending)
	if pending.Kind() == reflect.Map {
		for _, key := range pending.MapKeys() {
			address := key.Interface().(string)
			tx := reflect.ValueOf(pending.MapIndex(key).Interface())
			if tx.Kind() == reflect.Map && strings.ToLower(from) == strings.ToLower(address) {
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
	if pendingNonaceAt != 0 && txPoolMaxCount+1 > pendingNonaceAt {
		pendingNonce = txPoolMaxCount + 1
	}

	tx := types.NewTransaction(pendingNonce, tokenAddress, value, gasLimit, gasPrice, data)
	rawTxHex, err := model.EncodeETHTx(tx)
	if err != nil {
		return "", "", fmt.Errorf("Encode raw tx %s", err)
	}

	chainIDBig, err := repo.ethClient.ChainID(ctx)
	if err != nil {
		return "", "", err
	}
	if err := repo.db.Model(&tokenBalance).
		UpdateColumn("withdraw_lock", withdrawLockAmountDecimal.Add(transferAmountDecimal).String()).Error; err != nil {
		return "", "", fmt.Errorf("Fail to update withdrawLock row %s", err.Error())
	}
	return rawTxHex, chainIDBig.String(), nil
}

func (repo *repo) SendETHTx(ctx context.Context, signedTxHex string) (txID string, err error) {
	tx, err := model.DecodeETHTx(signedTxHex)
	if err != nil {
		return "", fmt.Errorf("Fail to SendETHTx %s", err)
	}
	ms, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
	if err != nil {
		return "", fmt.Errorf("Fail to extract tx ms %s", err.Error())
	}
	if err := repo.ethClient.SendTransaction(ctx, tx); err != nil {
		e := repo.rollbackETHTx(ms)
		repo.logger.Log("Fail to rollbackETHTx", e.Error())
		return "", fmt.Errorf("Fail to SendETHTx %s", err)
	}

	amount := new(big.Int)
	sender := ms.From().String()
	balance := model.Balance{}
	inputData := "0x" + common.Bytes2Hex(ms.Data())
	if strings.Contains(inputData, ERC20TransferMethodHex) {
		amountData := inputData[74:138]
		amountBytes, err := hex.DecodeString(amountData)
		if err != nil {
			return "", err
		}
		amount = amount.SetBytes(amountBytes)
		tokenIdendify := strings.ToLower(ms.To().String())
		if err := repo.db.Where("address = ? AND identify = ?", sender, tokenIdendify).First(&balance).Error; err != nil {
			return "", fmt.Errorf("Fail to query erc20 sender balance record, address:" + sender + "identify:" + tokenIdendify + err.Error())
		}
	} else {
		if err := repo.db.Where("address = ? AND symbol = ?", sender, model.ETHSymbol).First(&balance).Error; err != nil {
			return "", fmt.Errorf("Fail to find sender address in balance table %s", err.Error())
		}
		amount = tx.Value()
	}
	txRecord := model.Tx{}
	if err := repo.db.FirstOrCreate(&txRecord,
		model.Tx{
			TxID:      tx.Hash().String(),
			TxType:    model.TxTypeWithdraw,
			Address:   balance.Address,
			Asset:     balance.Symbol,
			Amount:    amount.String(),
			BalanceID: balance.ID,
			ChainName: model.ChainEthereum,
		}).Error; err != nil {
		return "", fmt.Errorf("Fail to create withdraw tx record %s", err.Error())
	}
	return tx.Hash().String(), nil
}

func (repo *repo) rollbackETHTx(ms types.Message) error {
	sender := ms.From().String()
	balance := model.Balance{}
	if err := repo.db.Where("address = ? AND symbol = ?", sender, model.ETHSymbol).First(&balance).Error; err != nil {
		return fmt.Errorf("Fail to find sender address in balance table %s", err.Error())
	}
	withdrawLockDecimal, err := decimal.NewFromString(balance.WithdrawLock)
	if err != nil {
		return fmt.Errorf("Extract originWithdrawLock error %s", err.Error())
	}
	if err := repo.db.Model(&balance).UpdateColumn("withdraw_lock",
		withdrawLockDecimal.Sub(decimal.NewFromBigInt(ms.Value(), 0)).String()).Error; err != nil {
		return fmt.Errorf("Fail to update withdrawLock row %s", err.Error())
	}
	return nil
}

func (repo *repo) constructTxParamsValidate(ctx context.Context, from, to, amount, symbol string) (balanceAmountDecimal,
	transferAmountDecimal, withdrawLockAmountDecimal decimal.Decimal, gasPrice *big.Int, transferValue *big.Int, balance model.Balance, rr error) {
	// withdraw and deposit address validate
	if !common.IsHexAddress(from) {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("Invalid address: %s", from)
	}
	if !common.IsHexAddress(to) {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("Invalid address: %s", to)
	}

	// balance data
	balance = model.Balance{}
	err := repo.db.Where("address = ? AND symbol = ?", from, symbol).First(&balance).Error
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("Fail to query balance %s", err.Error())
	}
	balanceAmount, result := new(big.Int).SetString(balance.Amount, 10)
	if !result {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("fail to extract amount from balance")
	}
	withdrawLockAmount, result := new(big.Int).SetString(balance.WithdrawLock, 10)
	if !result {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("fail to extract withdrawLock from balance")
	}
	withdrawLockAmountDecimal = decimal.NewFromBigInt(withdrawLockAmount, 0)
	balanceAmount = balanceAmount.Sub(balanceAmount, withdrawLockAmount)
	balanceAmountDecimal = decimal.NewFromBigInt(balanceAmount, 0)

	// transfer amount data
	transferAmountDecimal, err = decimal.NewFromString(amount)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("Fail to extract transfer amount %s", err.Error())
	}
	decimalBig, result := new(big.Int).SetString(strconv.FormatUint(balance.Decimal, 10), 10)
	if !result {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("Fail to convert decimal")
	}
	transferAmountDecimal = transferAmountDecimal.Mul(decimal.NewFromBigInt(decimalBig, 0))
	transferValue, ok := new(big.Int).SetString(transferAmountDecimal.String(), 10)
	if !ok {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("Set amount error")
	}

	// check balance is whether GreaterThan transfer amount
	if !balanceAmountDecimal.GreaterThan(transferAmountDecimal) {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("Insufficient balance")
	}

	// gas price
	gasPrice, err = repo.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, nil, nil, model.Balance{}, fmt.Errorf("Fail to query gas price %s", err.Error())
	}
	return
}
