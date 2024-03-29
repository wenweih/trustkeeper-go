package repository

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"trustkeeper-go/app/service/chains_query/pkg/model"
	common "trustkeeper-go/library/util"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcutil"
	"github.com/shopspring/decimal"
)

func (repo *repo) CreateBtcBlockWithUtxoPipeline(ctx context.Context, height int64) <-chan CreateBlockResult {
	queryBlockResult := repo.QueryBTCBlockCH(ctx, height)
	return repo.CreateBTCBlockWithUTXOs(ctx, queryBlockResult)
}

// QueryBTCBlockCH query bitcoin block
func (repo *repo) QueryBTCBlockCH(ctx context.Context, height int64) <-chan UTXOBlockResult {
	blockCh := make(chan UTXOBlockResult)
	go func(height int64) {
		defer close(blockCh)
		blockHash, err := repo.GetBTCBlockHashByHeight(ctx, height)
		if err != nil {
			e := fmt.Errorf("Query bitcoin block hash error: %s", err)
			blockCh <- UTXOBlockResult{Error: e}
			return
		}
		block, err := repo.QueryBitcoincoreBlock(ctx, blockHash)
		if err != nil {
			e := fmt.Errorf("Query bitcoin block error %s", err)
			blockCh <- UTXOBlockResult{Error: e}
			return
		}
		blockCh <- UTXOBlockResult{Block: block}
		return
	}(height)
	return blockCh
}

// CreateBTCBlockWithUTXOs save block and utxo related with subAddress blockResultCh <-chan
func (repo *repo) CreateBTCBlockWithUTXOs(ctx context.Context, queryBlockResultCh <-chan UTXOBlockResult) <-chan CreateBlockResult {
	createBlockCh := make(chan CreateBlockResult)
	go func() {
		defer close(createBlockCh)
		b := <-queryBlockResultCh
		if b.Error != nil {
			createBlockCh <- CreateBlockResult{Error: b.Error}
			return
		}
		rawBlock := b.Block
		var block model.BtcBlock
		ts := repo.db.Begin()
		if err := ts.FirstOrCreate(&block, model.BtcBlock{
			Hash:   rawBlock.Hash,
			Height: rawBlock.Height,
		}).Error; err != nil {
			ts.Rollback()
			createBlockCh <- CreateBlockResult{Error: fmt.Errorf("create block error: %s", err)}
			return
		}

		for _, tx := range rawBlock.Tx {
			for _, vout := range tx.Vout {
				voutScriptPubKeyHex := vout.ScriptPubKey.Hex
				// voutScriptPubKeyHex begin with 6a146f6d6e6900000000 is omni tx
				// hex[20:28] is token identifier
				// hex[28:44] is token amount to transfer
				if strings.Contains(voutScriptPubKeyHex, "6a146f6d6e6900000000") {
					// omni token transfer
					for _, vin := range tx.Vin {
						// query previous vout for vin
						vinTx, err := repo.QueryBTCTx(ctx, vin.Txid)
						if err != nil {
							createBlockCh <- CreateBlockResult{Error: err}
							return
						}
						voutForVin := vinTx.Vout[vin.Vout]
						vinAddresses := voutForVin.ScriptPubKey.Addresses
						for _, voutOmni := range tx.Vout {
							// find omni token receiver address
							// recharge address is equal to omni token sender
							if voutOmni.N != vout.N &&
								!reflect.DeepEqual(vinAddresses, voutOmni.ScriptPubKey.Addresses) {
								omniPropertyID := common.Hex2int(voutScriptPubKeyHex[20:28])
								var (
									balance  model.Balance
									txRecord model.Tx
								)
								hasBalance := true
								for _, balanceAddress := range voutOmni.ScriptPubKey.Addresses {
									if err := ts.Where("address = ? AND identify = ?",
										balanceAddress, strconv.FormatInt(omniPropertyID, 10)).First(&balance).Error; err != nil && err.Error() == "record not found" {
										hasBalance = false
										continue
									} else if err != nil {
										createBlockCh <- CreateBlockResult{Error: fmt.Errorf("Query subscript address from balance_t err: %s", err)}
										return
									}
								}
								if !hasBalance {
									continue
								}
								// create omni deposit tx record
								ts.FirstOrCreate(&txRecord,
									model.Tx{
										TxID:      tx.Txid,
										TxType:    model.TxTypeDeposit,
										Address:   balance.Address,
										Asset:     balance.Symbol,
										Amount:    strconv.FormatInt(common.Hex2int(voutScriptPubKeyHex[28:44]), 10),
										BalanceID: balance.ID,
										ChainName: model.ChainBitcoin})
							}
						}
					}
				} else if vout.Value != 0 && vout.ScriptPubKey.Addresses != nil {
					// btc transfer
					for _, address := range vout.ScriptPubKey.Addresses {
						var (
							balance  model.Balance
							utxo     model.BtcUtxo
							txRecord model.Tx
						)
						// query balance record
						if err := ts.Where("address = ? AND Symbol = ?", address, model.BTCSymbol).
							First(&balance).Error; err != nil && err.Error() == "record not found" {
							continue
						} else if err != nil {
							createBlockCh <- CreateBlockResult{Error: fmt.Errorf("Query sub address err: %s", err)}
							return
						}

						// create utxo for btc balance record
						ts.FirstOrCreate(
							&utxo,
							model.BtcUtxo{
								Txid:       tx.Txid,
								Amount:     vout.Value,
								Height:     rawBlock.Height,
								VoutIndex:  vout.N,
								BalanceID:  balance.ID,
								BtcBlockID: block.ID,
							})

						// create btc deposit tx record
						ts.FirstOrCreate(&txRecord,
							model.Tx{
								TxID:      tx.Txid,
								TxType:    model.TxTypeDeposit,
								Address:   balance.Address,
								Asset:     balance.Symbol,
								Amount:    strconv.FormatFloat(vout.Value*btcutil.SatoshiPerBitcoin, 'f', -int(0), 64),
								BalanceID: balance.ID,
								ChainName: model.ChainBitcoin,
							})
					}
				}
			}
		}
		if err := ts.Commit().Error; err != nil {
			ts.Rollback()
			createBlockCh <- CreateBlockResult{Error: fmt.Errorf("database transaction err: %s", err)}
			return
		}
		createBlockCh <- CreateBlockResult{Block: &block}
	}()
	return createBlockCh
}

// UTXOBlockResult query block result
type UTXOBlockResult struct {
	Error error
	Block *btcjson.GetBlockVerboseResult
}

// CreateBlockResult save block record result
type CreateBlockResult struct {
	Error error
	Block *model.BtcBlock
}

// TrackBlock rollback 6 blocks when save new block records
func (repo *repo) TrackBlock(
	ctx context.Context, bestBlockHeight int64, isTracking bool, queryBlockResultCh <-chan UTXOBlockResult) (bool, int64) {
	b := <-queryBlockResultCh
	if b.Error != nil {
		repo.logger.Log("queryBlockResultChError", b.Error.Error())
	}
	rawBlock := b.Block
	trackHeight := rawBlock.Height

	var (
		dbBlock model.BtcBlock
		utxos   []model.BtcUtxo
	)

	if err := repo.db.First(&dbBlock, "height = ? AND re_org = ?", rawBlock.Height, false).
		Related(&utxos).Error; err != nil && err.Error() == "record not found" {
		dbBlock.Hash = rawBlock.Hash
		dbBlock.Height = rawBlock.Height
		blockCh := make(chan UTXOBlockResult)
		go func(rawBlock *btcjson.GetBlockVerboseResult) {
			defer close(blockCh)
			blockCh <- UTXOBlockResult{Block: rawBlock}
		}(rawBlock)
		createBlockResul := <-repo.CreateBTCBlockWithUTXOs(ctx, blockCh)
		if createBlockResul.Error != nil {
			repo.logger.Log("createBlockResulError", createBlockResul.Error.Error())
		}

		if createBlockResul.Error == nil {
			bestBlock := createBlockResul.Block
			repo.logger.Log("CreateBlock", bestBlock.Height, "Hash", bestBlock.Hash)
		}
	} else if err != nil {
		repo.logger.Log("QueryTrackBlockError", err.Error())
	} else {
		if dbBlock.Hash != rawBlock.Hash {
			ts := repo.db.Begin()
			// update utxos related with the dbBlock
			ts.Model(&dbBlock).Update("re_org", true)
			for _, utxo := range utxos {
				ts.Model(&utxo).Update("re_org", true)
			}
			ts.Commit()

			blockCh := make(chan UTXOBlockResult)
			blockCh <- UTXOBlockResult{Block: rawBlock}
			createBlockResul := <-repo.CreateBTCBlockWithUTXOs(ctx, blockCh)
			close(blockCh)
			if createBlockResul.Error == nil {
				bestBlock := createBlockResul.Block
				repo.logger.Log("CreateBlock", bestBlock.Height, " hash: ", bestBlock.Hash)
			}
			repo.logger.Log("Reorg", dbBlock.Height, "Hash", dbBlock.Hash)
		} else {
			repo.logger.Log("trackingTheSameBlock", dbBlock.Height)
		}
	}

	if trackHeight < bestBlockHeight-5 {
		isTracking = false
	} else {
		isTracking = true
		trackHeight--
	}
	return isTracking, trackHeight
}

func (repo *repo) UpdateBitcoincoreTx(ctx context.Context) {
	txes := make([]model.Tx, 0)
	err := repo.db.Where("chain_name = ?", model.ChainBitcoin).
		Not("state", []string{model.StateSuccess, model.StateFail}).Find(&txes).Error
	if err != nil {
		repo.logger.Log("UpdateEthereumTx:", err.Error())
	}
	for _, tx := range txes {
		ts := repo.db.Begin()
		rawtx, err := repo.QueryBTCTx(ctx, tx.TxID)
		if err != nil {
			repo.logger.Log("UpdateBitcoincoreTx:", err.Error())
			continue
		}
		if rawtx.Confirmations >= DepositBitcoincoreComfirmation {
			transferAmount, result := new(big.Int).SetString(tx.Amount, 10)
			if !result {
				repo.logger.Log("Fail to extract transfer amount")
			}

			balance := model.Balance{}
			ts.Model(&tx).Related(&balance)
			balanceAmount, result := new(big.Int).SetString(balance.Amount, 10)
			if !result {
				repo.logger.Log("Fail to extract balance amount")
			}
			withdrawLock, result := new(big.Int).SetString(balance.WithdrawLock, 10)
			if !result {
				repo.logger.Log("Fail to extract withdrawLock")
			}
			withdrawLockDecimal := decimal.NewFromBigInt(withdrawLock, 0)
			if tx.TxType == model.TxTypeDeposit {
				// calculate amount row for balance record
				balanceAmount = balanceAmount.Add(balanceAmount, transferAmount)
			} else if tx.TxType == model.TxTypeWithdraw {
				if balance.Symbol == model.BTCSymbol {
					// calculate btc transfer withdrawLock
					// query raw tx
					rawTx, err := repo.QueryBTCTx(ctx, tx.TxID)
					if err != nil {
						repo.logger.Log("FailToQueryTxWhenWithdrawSuccess:", err.Error())
						continue
					}

					// calculate withdrawLock and amount rows for btc balance record
					for _, vin := range rawTx.Vin {
						utxo := model.BtcUtxo{}
						ts.Preload("Balance").
							Where("txid = ? AND vout_index = ?", vin.Txid, vin.Vout).First(&utxo)
						utxoAmountDecimal := decimal.NewFromFloat(utxo.Amount).Mul(decimal.New(int64(utxo.Balance.Decimal), 0))
						withdrawLockDecimal = withdrawLockDecimal.Sub(utxoAmountDecimal)
					}
				} else {
					// calculate omni token transfer withdrawLock
					withdrawLockDecimal = withdrawLockDecimal.Sub(decimal.NewFromBigInt(transferAmount, 0))
				}
				balanceAmount = balanceAmount.Sub(balanceAmount, transferAmount)
			}

			balanceAmountStr := balanceAmount.String()
			withdrawLockStr := withdrawLockDecimal.String()

			ts.Create(&model.BalanceLog{TxID: tx.TxID, From: balance.Amount, To: balanceAmountStr, BalanceID: balance.ID})
			ts.Model(&balance).UpdateColumns(model.Balance{Amount: balanceAmountStr, WithdrawLock: withdrawLockStr})

			repo.logger.Log("UpdateBitcoincoreTx", tx.TxID,
				"StateFrom", tx.State, "To", model.StateSuccess, "ConfirmationsFrom", tx.Confirmations, "To", rawtx.Confirmations)
			ts.Model(&tx).UpdateColumns(model.Tx{State: model.StateSuccess, Confirmations: int64(rawtx.Confirmations)})
		} else {
			repo.logger.Log("UpdateBitcoincoreTx", tx.TxID, "Confirmation", tx.Confirmations, "To", rawtx.Confirmations)
			ts.Model(&tx).UpdateColumn("confirmations", rawtx.Confirmations)
		}
		if err := ts.Commit().Error; err != nil {
			ts.Rollback()
			repo.logger.Log("UpdateBitcoincoreTx commit ts:", err.Error())
		}
	}
}
