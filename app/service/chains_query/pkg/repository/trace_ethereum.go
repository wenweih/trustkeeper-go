package repository

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"trustkeeper-go/app/service/chains_query/pkg/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

func (repo *repo) EthererumDBBestBlock(ctx context.Context) (*model.EthBlock, error) {
	block := model.EthBlock{}
	if err := repo.db.Order("height desc").First(&block).Error; err != nil {
		return nil, fmt.Errorf("fail to query ethereum best in db:" + err.Error())
	}
	return &block, nil
}

func (repo *repo) CreateETHBlockWithTx(ctx context.Context, height int64) (*model.EthBlock, error) {
	number := new(big.Int)
	number.SetInt64(height)
	block, err := repo.EthereumBlock(ctx, number)
	if err != nil {
		return nil, err
	}
	blockBytes, err := model.EncodeETHBlock(*block)
	if err != nil {
		return nil, err
	}
	simpleBlock := model.EthereumBlock{}
	buf := bytes.NewBuffer(blockBytes)
	dc := gob.NewDecoder(buf)
	if err := dc.Decode(&simpleBlock); err != nil {
		return nil, err
	}
	dbBlock := model.EthBlock{}
	ts := repo.db.Begin()
	if err := ts.FirstOrCreate(&dbBlock, model.EthBlock{
		Hash:   simpleBlock.Hash.String(),
		Height: simpleBlock.Header.Number.Int64(),
	}).Error; err != nil {
		ts.Rollback()
		return nil, fmt.Errorf("create eth block error: %s", err)
	}
	for _, tx := range simpleBlock.Tx {
		// 0xa9059cbb is first 4 bytes of the resulting hash is the methodID as token transfer tx [0:10]
		// token transfer to address is 32 bytes in inpute data [10:74]
		// token transfer amount is 32 bytes in input data [74:138]
		if strings.Contains(tx.Data, "0xa9059cbb") {
			toData := tx.Data[10:74]
			receiver := common.HexToAddress(toData)

			amountData := tx.Data[74:138]
			amountBytes, err := hex.DecodeString(amountData)
			if err != nil {
				return nil, err
			}
			amount := new(big.Int).SetBytes(amountBytes)
			var (
				txRecord model.Tx
				balance  model.Balance
			)
			if err := ts.Where("address = ? AND identify = ?",
				receiver.Hex(), strings.ToLower(tx.To)).First(&balance).Error; err != nil && err.Error() != "record not found" {
				ts.Rollback()
				return nil, fmt.Errorf("Fail to query erc20 balance record, address:" + receiver.Hex() + "identify:" + tx.To + err.Error())
			} else if err != nil && err.Error() == "record not found" {
				continue
			}
			ts.FirstOrCreate(&txRecord, model.Tx{
				TxID:      tx.Txid,
				TxType:    model.TxTypeDeposit,
				Address:   receiver.Hex(),
				Asset:     balance.Symbol,
				Amount:    amount.String(),
				BalanceID: balance.ID,
				ChainName: model.ChainEthereum})
		} else {
			var (
				txRecord model.Tx
				balance  model.Balance
			)
			if err := ts.Where("address = ? AND Symbol = ?", tx.To, "ETH").First(&balance).Error; err != nil && err.Error() != "record not found" {
				ts.Rollback()
				return nil, fmt.Errorf("Fail to query eth balance record, address:" + tx.To + " " + err.Error())
			} else if err != nil && err.Error() == "record not found" {
				continue
			}
			amountBig, err := hexutil.DecodeBig(tx.ValueHex)
			if err != nil {
				return nil, err
			}

			ts.FirstOrCreate(&txRecord, model.Tx{
				TxID:      tx.Txid,
				TxType:    model.TxTypeDeposit,
				Address:   tx.To,
				Asset:     balance.Symbol,
				Amount:    amountBig.String(),
				BalanceID: balance.ID,
				ChainName: model.ChainEthereum})
		}
	}
	if err := ts.Commit().Error; err != nil {
		return nil, err
	}
	return &dbBlock, nil
}

func (repo *repo) UpdateEthereumTx(ctx context.Context) {
	txes := make([]model.Tx, 0)
	err := repo.db.Where("chain_name = ?", model.ChainEthereum).
		Not("state", []string{model.StateSuccess, model.StateFail}).Find(&txes).Error
	if err != nil {
		repo.logger.Log("UpdateEthereumTx fail to query txed:", err.Error())
		return
	}
	bestBlockHead, err := repo.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		repo.logger.Log("UpdateEthereumTx fail to query bestBlockHead:", err.Error())
		return
	}
	for _, tx := range txes {
		receipt, err := repo.QueryEthereumTxReceipt(ctx, tx.TxID)
		if err != nil {
			repo.logger.Log("UpdateEthereumTx fail to query ethereum tx receipt:", err.Error())
			return
		}
		// pending tx: update tx record state row only
		if receipt.BlockNumber == nil {
			// blockNumber field will be null until the transaction is included into a mined block
			repo.logger.Log("UpdateEthereumTx", tx.TxID, "StateFrom", tx.State, "To", model.StatePending)
			if err := repo.db.Model(&tx).UpdateColumn("state", model.StatePending).Error; err != nil {
				repo.logger.Log("Fail to tx state txid", tx.TxID, "err", err.Error())
			}
		} else if receipt.Status == 1 {
			// valid tx
			// Since block 4370000 (Byzantium), a status indicator has been added to receipts. 1 mean success, 0 mean fail
			confirmations := bestBlockHead.Number.Int64() - receipt.BlockNumber.Int64() + 1
			// confirmations account greater than min setting confirmations, update tx state to success and balance record amount and withdrawLock rows
			if confirmations >= DepositEthereumComfirmation {
				// transfer amount
				transferAmount, result := new(big.Int).SetString(tx.Amount, 10)
				if !result {
					repo.logger.Log("Fail to extract transfer amount")
					return
				}
				transferAmountDecimal := decimal.NewFromBigInt(transferAmount, 0)

				// query tx related balance record
				balance := model.Balance{}
				repo.db.Model(&tx).Related(&balance)
				balanceAmount, result := new(big.Int).SetString(balance.Amount, 10)
				if !result {
					repo.logger.Log("Fail to extract balance amount")
					return
				}

				// balance amount
				balanceAmountDecimal := decimal.NewFromBigInt(balanceAmount, 0)
				withdrawLock, result := new(big.Int).SetString(balance.WithdrawLock, 10)
				if !result {
					repo.logger.Log("Fail to extract withdrawLock")
					return
				}

				ethBalance := model.Balance{}
				ethAmountDecimal := decimal.New(0, 0)
				withdrawLockDecimal := decimal.NewFromBigInt(withdrawLock, 0) // withdrawLock amount
				if tx.TxType == model.TxTypeDeposit {
					// deposit tx, add trans amount to balance amount
					balanceAmountDecimal = balanceAmountDecimal.Add(transferAmountDecimal)
				} else if tx.TxType == model.TxTypeWithdraw {
					// query raw tx from blockchain
					rawTx, err := repo.QueryEthereumTx(ctx, tx.TxID)
					if err != nil {
						repo.logger.Log("Fail to query ethereum tx", err.Error())
						return
					}

					// calculate tx fee
					fee := decimal.NewFromBigInt(rawTx.GasPrice(), 0).Mul(decimal.New(int64(receipt.CumulativeGasUsed), 0))
					inputData := "0x" + common.Bytes2Hex(rawTx.Data()) // tx input field data
					if strings.Contains(inputData, ERC20TransferMethodHex) {
						// erc20 tx
						// calculate erc20 balance record's amount
						balanceAmountDecimal = balanceAmountDecimal.Sub(transferAmountDecimal)

						// query eth balance by sender address
						if err := repo.db.Where("address = ? AND symbol = ?", tx.Address, model.ETHSymbol).First(&ethBalance).Error; err != nil {
							repo.logger.Log("Fail to query ethereum balance record", err.Error())
							return
						}
						ethBalanceAmount, err := decimal.NewFromString(ethBalance.Amount)
						if err != nil {
							repo.logger.Log("Fail to extract amount row for eth balance", err.Error())
							return
						}
						ethAmountDecimal = ethBalanceAmount.Sub(fee)
					} else {
						// eth transfer
						// calculate erc20 balance record's amount
						balanceAmountDecimal = balanceAmountDecimal.Sub(transferAmountDecimal).Sub(fee)
					}

					// calculate withdrawLock
					withdrawLockDecimal = withdrawLockDecimal.Sub(decimal.NewFromBigInt(transferAmount, 0))
				}
				balanceAmountStr := balanceAmountDecimal.String()
				withdrawLockStr := withdrawLockDecimal.String()

				// 1. create balance update log record
				// 2. update amount and withdrawLock rows for balance
				// 3. update state and confirmations rows for tx
				ts := repo.db.Begin()
				ts.Create(&model.BalanceLog{TxID: tx.TxID, From: balance.Amount, To: balanceAmountStr, BalanceID: balance.ID})
				ts.Model(&balance).UpdateColumns(model.Balance{Amount: balanceAmountStr, WithdrawLock: withdrawLockStr})
				repo.logger.Log("UpdateEthereumTx", tx.TxID, "StateFrom", tx.State, "To", model.StateSuccess, "ConfirmationsFrom", tx.Confirmations, "To", confirmations)
				ts.Model(&tx).UpdateColumns(model.Tx{State: model.StateSuccess, Confirmations: confirmations})

				// erc20 transfer, update eth balance record's amount  for fee
				if len(ethBalance.Address) > 1 {
					ts.Create(&model.BalanceLog{TxID: tx.TxID, From: ethBalance.Amount, To: ethAmountDecimal.String(), BalanceID: ethBalance.ID})
					ts.Model(&ethBalance).UpdateColumns(model.Balance{Amount: ethAmountDecimal.String()})
				}
				// commit ts
				if err := ts.Commit().Error; err != nil {
					ts.Rollback()
					repo.logger.Log("UpdateEthereumTx fail to commit data:", err.Error())
					return
				}
			} else {
				// Confirming tx, update state and confirmations rows for tx
				repo.logger.Log("UpdateEthereumTx", tx.TxID, "StateFrom", tx.State, "To", model.StateConfirming, "ConfirmationsFrom", tx.Confirmations, "To", confirmations)
				err := repo.db.Model(&tx).UpdateColumns(model.Tx{State: model.StateConfirming, Confirmations: confirmations}).Error
				if err != nil {
					repo.logger.Log("Fail to update tx txid", tx.TxID, "Err", err.Error())
				}
			}
		} else if receipt.Status == 0 {
			// Invalid tx, update state row for tx
			repo.logger.Log("UpdateEthereumTx", tx.TxID, "StateFrom", tx.State, "To", model.StateFail)
			if err := repo.db.Model(&tx).UpdateColumn("state", model.StateFail).Error; err != nil {
				repo.logger.Log("Fail to update tx txid", tx.TxID, "err", err.Error())
			}
		}
	}
}
