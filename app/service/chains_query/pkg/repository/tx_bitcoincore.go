package repository

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"trustkeeper-go/app/service/chains_query/pkg/model"
	common "trustkeeper-go/library/util"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/coinset"
	"github.com/shopspring/decimal"
)

func (repo *repo) ConstructTxBTC(ctx context.Context, from, to, amount string) (string, int64, error) {
	fromPkScript, toPkScript, ledgerInfo, estimateFeeResult, transferAmountBig, balance, err := repo.constructBitcoincoreTxParamsValidate(ctx, from, to, amount, model.BTCSymbol)
	if err != nil {
		return "", 0, err
	}
	var txids []string
	if err := repo.db.Model(&model.Tx{State: model.StateSuccess,
		ChainName: model.ChainBitcoin,
		BalanceID: balance.ID,
		Asset:     model.BTCSymbol}).Pluck("tx_id", &txids).Error; err != nil {
		return "", 0, fmt.Errorf("Fail to query comfirmations btc tx:" + err.Error())
	}
	var matureUTXOs []model.BtcUtxo
	if err := repo.db.Where("txid IN (?) AND balance_id = ? AND re_org = ? AND state = ?",
		txids, balance.ID, false, model.UTXOStateUnSelected).Find(&matureUTXOs).Error; err != nil {
		return "", 0, fmt.Errorf("fail to query mature utxos:" + err.Error())
	}

	selectedutxos, unselectedutxos, selectedCoins, err := CoinSelect(int64(ledgerInfo.Headers), btcutil.Amount(transferAmountBig.Int64()), matureUTXOs)
	if err != nil {
		return "", 0, fmt.Errorf("fail to select UTXO %s", err.Error())
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
	txSize := msgTx.SerializeSize() + 181 + 34
	txSizeDecimal := decimal.New(int64(txSize), 0)
	feeRateDecimal := decimal.NewFromFloat(estimateFeeResult.FeeRate).Div(decimal.New(1000, 0))
	fee := feeRateDecimal.Mul(txSizeDecimal).Mul(decimal.New(100000000, 0)).IntPart()
	if (vinAmount - transferAmountBig.Int64() - fee) > 0 {
		txOutReCharge := wire.NewTxOut((vinAmount - transferAmountBig.Int64() - int64(fee)), fromPkScript)
		msgTx.AddTxOut(txOutReCharge)
	} else {
		selectedutxoForFee, _, selectedCoinsForFee, err := CoinSelect(int64(ledgerInfo.Headers), btcutil.Amount(fee), unselectedutxos)
		if err != nil {
			return "", 0, fmt.Errorf("Select UTXO for fee %s", err)
		}
		for _, coin := range selectedCoinsForFee.Coins() {
			vinAmount += int64(coin.Value())
		}
		for _, feeUTXO := range selectedutxoForFee {
			txHash, err := chainhash.NewHashFromStr(feeUTXO.Txid)
			if err != nil {
				return "", 0, fmt.Errorf("Fail to construct fee vin %s", err.Error())
			}
			prevOut := wire.NewOutPoint(txHash, feeUTXO.VoutIndex)
			msgTx.AddTxIn(wire.NewTxIn(prevOut, []byte{txscript.OP_0, txscript.OP_0}, nil))
		}
		txOutReCharge := wire.NewTxOut((vinAmount - transferAmountBig.Int64() - int64(fee)), fromPkScript)
		msgTx.AddTxOut(txOutReCharge)
		selectedutxos = append(selectedutxos, selectedutxoForFee...)
	}
	utxoids := make([]uint, len(selectedutxos))
	utxoAmountDecimal := decimal.NewFromFloat(0)
	for i, selectedUTXO := range selectedutxos {
		selectedDecimal := decimal.NewFromFloat(selectedUTXO.Amount)
		utxoAmountDecimal = utxoAmountDecimal.Add(selectedDecimal)
		utxoids[i] = selectedUTXO.ID
	}
	balanceDecimalBig, result := new(big.Int).SetString(strconv.FormatUint(balance.Decimal, 10), 10)
	if !result {
		return "", 0, fmt.Errorf("Fail to convert decimal")
	}
	utxoAmountDecimal = utxoAmountDecimal.Mul(decimal.NewFromBigInt(balanceDecimalBig, 0))
	ts := repo.db.Begin()
	balanceWithdrawLockDecimal, err := decimal.NewFromString(balance.WithdrawLock)
	if err != nil {
		return "", 0, fmt.Errorf("Fail to extract balance withdrawLock to decimal")
	}
	ts.Model(&balance).UpdateColumn("withdraw_lock", balanceWithdrawLockDecimal.Add(utxoAmountDecimal).String())
	ts.Table("btc_utxos").Where("id IN (?)", utxoids).UpdateColumn("state", model.UTXOStateLocked)
	if err := ts.Commit().Error; err != nil {
		return "", 0, fmt.Errorf("Fail to update selected utxo state %s", err.Error())
	}

	buf := bytes.NewBuffer(make([]byte, 0, msgTx.SerializeSize()))
	msgTx.Serialize(buf)
	rawTxHex := hex.EncodeToString(buf.Bytes())
	return rawTxHex, vinAmount, nil
}

func (repo *repo) ConstructTxOmni(ctx context.Context, from, to, amount, symbol string) (unsignedTxHex string, vinAmount int64, err error) {
	fromPkScript, toPkScript, ledgerInfo, estimateFeeResult, transferAmountBig, balance, err := repo.constructBitcoincoreTxParamsValidate(ctx, from, to, amount, symbol)
	if err != nil {
		return "", 0, err
	}

	// OmniToken transfer tx
	msgTx := wire.NewMsgTx(wire.TxVersion)

	// op_return data
	omniVersion := common.Int2byte(uint64(0), 2) // omnicore version
	txType := common.Int2byte(uint64(0), 2)      // omnicore tx type: simple send
	tokenPropertyid, err := strconv.Atoi(balance.Identify)
	if err != nil {
		return "", 0, fmt.Errorf("tokenPropertyid to int %s", err)
	}
	tokenIdentifier := common.Int2byte(uint64(tokenPropertyid), 4) // omni token identifier
	tokenAmount := common.Int2byte(transferAmountBig.Uint64(), 8)  // omni token transfer amount

	// tx script builder
	b := txscript.NewScriptBuilder()
	b.AddOp(txscript.OP_RETURN) // 2byte = 4 hex
	b.AddData([]byte("omni"))   // transaction maker 4 byte = 8 hex
	b.AddData(omniVersion)      // 2 byte = 4 hex
	b.AddData(txType)           // 2 byte = 4 hex
	b.AddData(tokenIdentifier)  // 4 byte = 8 hex
	b.AddData(tokenAmount)      // 8 byte = 16 hex

	// pk script
	pkScript, err := b.Script()
	if err != nil {
		return "", 0, fmt.Errorf("Bitcoin Token pkScript %s", err)
	}

	msgTx.AddTxOut(wire.NewTxOut(0, pkScript)) // omni amount vout

	// calculate fee 1vin + 3vouts
	txSize := msgTx.SerializeSize() + 181 + 80
	txSizeDecimal := decimal.New(int64(txSize), 0)
	feeRateDecimal := decimal.NewFromFloat(estimateFeeResult.FeeRate).Div(decimal.New(1000, 0))
	fee := feeRateDecimal.Mul(txSizeDecimal).Mul(decimal.New(100000000, 0)).IntPart()

	// query sender btc balance record in db
	btcBalance := &model.Balance{}
	if err := repo.db.Where("address = ? AND symbol = ?", from, model.BTCSymbol).First(&btcBalance).Error; err != nil {
		return "", 0, fmt.Errorf("Fail to query btc balance %s", err.Error())
	}

	// query mature btc tx in db
	var txids []string
	if err := repo.db.Model(&model.Tx{State: model.StateSuccess,
		ChainName: model.ChainBitcoin,
		BalanceID: btcBalance.ID,
		Asset:     model.BTCSymbol}).Pluck("tx_id", &txids).Error; err != nil {
		return "", 0, fmt.Errorf("Fail to query comfirmations btc tx:" + err.Error())
	}

	// coin selection
	var matureUTXOs []model.BtcUtxo
	if err := repo.db.Where("txid IN (?) AND balance_id = ? AND re_org = ? AND state = ?",
		txids, btcBalance.ID, false, model.UTXOStateUnSelected).Find(&matureUTXOs).Error; err != nil {
		return "", 0, fmt.Errorf("fail to query mature utxos:" + err.Error())
	}
	selectedutxos, _, selectedCoins, err := CoinSelect(int64(ledgerInfo.Headers), btcutil.Amount(fee), matureUTXOs)
	if err != nil {
		return "", 0, fmt.Errorf("fail to select UTXO %s", err.Error())
	}

	// calculate the sum of vins
	vinAmount = 0
	for _, coin := range selectedCoins.Coins() {
		vinAmount += int64(coin.Value())
	}

	// vins for tx
	for _, vin := range selectedutxos {
		txHash, err := chainhash.NewHashFromStr(vin.Txid)
		if err != nil {
			return "", 0, fmt.Errorf("Fail to construct fee vin %s", err.Error())
		}
		prevOut := wire.NewOutPoint(txHash, vin.VoutIndex)
		msgTx.AddTxIn(wire.NewTxIn(prevOut, []byte{txscript.OP_0, txscript.OP_0}, nil))
	}
	msgTx.AddTxOut(wire.NewTxOut(vinAmount-fee, fromPkScript)) //sendingaddress vout
	msgTx.AddTxOut(wire.NewTxOut(0, toPkScript))               //reference(omni token receiver) vout

	// calculate btc balance's withdrawLock
	utxoids := make([]uint, len(selectedutxos))
	utxoAmountDecimal := decimal.NewFromFloat(0)
	for i, selectedUTXO := range selectedutxos {
		selectedDecimal := decimal.NewFromFloat(selectedUTXO.Amount)
		utxoAmountDecimal = utxoAmountDecimal.Add(selectedDecimal)
		utxoids[i] = selectedUTXO.ID
	}
	btcBalanceDecimalBig, result := new(big.Int).SetString(strconv.FormatUint(btcBalance.Decimal, 10), 10)
	if !result {
		return "", 0, fmt.Errorf("Fail to convert decimal")
	}
	utxoAmountDecimal = utxoAmountDecimal.Mul(decimal.NewFromBigInt(btcBalanceDecimalBig, 0))
	balanceWithdrawLockDecimal, err := decimal.NewFromString(btcBalance.WithdrawLock)
	if err != nil {
		return "", 0, fmt.Errorf("Fail to extract btc balance withdrawLock to decimal")
	}

	// omni balance's withdrawLock
	omniBalanceWithdrawLockDecimal, err := decimal.NewFromString(balance.WithdrawLock)
	if err != nil {
		return "", 0, fmt.Errorf("Fail to extract balance withdrawLock to decimal")
	}

	ts := repo.db.Begin()
	// update omni balance's withdrawLock
	ts.Model(&balance).UpdateColumn("withdraw_lock", omniBalanceWithdrawLockDecimal.Add(decimal.NewFromBigInt(transferAmountBig, 0)).String())
	// update btc balance's withdrawLock
	ts.Model(&btcBalance).UpdateColumn("withdraw_lock", balanceWithdrawLockDecimal.Add(utxoAmountDecimal).String())
	// update selected utxo
	ts.Table("btc_utxos").Where("id IN (?)", utxoids).UpdateColumn("state", model.UTXOStateLocked)
	// commit ts
	if err := ts.Commit().Error; err != nil {
		return "", 0, fmt.Errorf("Fail to update selected utxo state %s", err.Error())
	}

	buf := bytes.NewBuffer(make([]byte, 0, msgTx.SerializeSize()))
	msgTx.Serialize(buf)
	rawTxHex := hex.EncodeToString(buf.Bytes())
	return rawTxHex, vinAmount, nil
}

func (repo *repo) SendBTCTx(ctx context.Context, signedTxHex string) (string, error) {
	txByte, err := hex.DecodeString(signedTxHex)
	if err != nil {
		return "", fmt.Errorf("Fail to decode tx hex %s", err.Error())
	}
	var msgTx wire.MsgTx
	if err := msgTx.Deserialize(bytes.NewReader(txByte)); err != nil {
		return "", fmt.Errorf("fail to deserialize tx %s", err.Error())
	}
	tx := btcutil.NewTx(&msgTx)
	txHash, err := repo.bitcoinClient.SendRawTransaction(tx.MsgTx(), false)
	if err != nil {
		// rollback utxo and balance withdrawLock row
		ts := repo.db.Begin()
		var balanceWithdrawLock = struct {
			sync.RWMutex
			m map[uint]decimal.Decimal
		}{m: make(map[uint]decimal.Decimal)}
		for _, vin := range tx.MsgTx().TxIn {
			utxo := model.BtcUtxo{}
			// rollback selected utxo using by vins
			ts.Preload("Balance").
				Where("txid = ? AND vout_index = ?", vin.PreviousOutPoint.Hash.String(), vin.PreviousOutPoint.Index).
				First(&utxo).UpdateColumn("state", model.UTXOStateUnSelected)
			utxoAmountDecimal := decimal.NewFromFloat(utxo.Amount)
			balanceDecimalBig, result := new(big.Int).SetString(strconv.FormatUint(utxo.Balance.Decimal, 10), 10)
			if !result {
				return "", fmt.Errorf("Fail to convert decimal")
			}

			// construct balanceWithdrawLock by balance
			utxoAmountDecimal = utxoAmountDecimal.Mul(decimal.NewFromBigInt(balanceDecimalBig, 0))
			balanceWithdrawLock.Lock()
			balanceWithdrawLock.m[utxo.Balance.ID] = balanceWithdrawLock.m[utxo.Balance.ID].Add(utxoAmountDecimal)
			balanceWithdrawLock.Unlock()
		}
		for k, v := range balanceWithdrawLock.m {
			balance := model.Balance{}
			ts.First(&balance, k)
			vinWithdrawLock, err := decimal.NewFromString(balance.WithdrawLock)
			if err != nil {
				ts.Rollback()
				return "", fmt.Errorf("Extract originWithdrawLock error %s", err.Error())
			}
			ts.Model(&balance).UpdateColumn("withdraw_lock", vinWithdrawLock.Sub(v).String())

			for _, vout := range msgTx.TxOut {
				pkScriptHex := hex.EncodeToString(vout.PkScript)
				if strings.Contains(pkScriptHex, "6f6d6e69") {
					omniPropertyID := common.Hex2int(pkScriptHex[26:34])
					transferAmount := common.Hex2int(pkScriptHex[36:])
					omnibalance := model.Balance{}
					ts.Where("address = ? AND identify = ?", balance.Address, strconv.FormatInt(omniPropertyID, 10)).First(&omnibalance)
					omniWithdrawLock, err := decimal.NewFromString(omnibalance.WithdrawLock)
					if err != nil {
						ts.Rollback()
						return "", fmt.Errorf("Extract omni balance WithdrawLock error %s", err.Error())
					}
					ts.Model(&omnibalance).UpdateColumn("withdraw_lock", omniWithdrawLock.Sub(decimal.New(transferAmount, 0)).String())
				}
			}
		}
		if err := ts.Commit().Error; err != nil {
			return "", fmt.Errorf("Fail to rollback utxo and balance withdrawLock row %s", err.Error())
		}
		return "", fmt.Errorf("Bitcoin SendRawTransaction %s", err)
	}

	txid := txHash.String()
	var balanceWithdrawTx = struct {
		sync.RWMutex
		m map[uint]decimal.Decimal
	}{m: make(map[uint]decimal.Decimal)}

	ts := repo.db.Begin()
	for _, vin := range tx.MsgTx().TxIn {
		utxo := model.BtcUtxo{}
		if err := ts.Preload("Balance").
			Where("txid = ? AND vout_index = ?", vin.PreviousOutPoint.Hash.String(), vin.PreviousOutPoint.Index).First(&utxo).
			UpdateColumns(model.BtcUtxo{UsedBy: txid, State: model.UTXOStateSelected}).Error; err != nil {
			ts.Rollback()
			return "", fmt.Errorf("Fail to update utxo state when send tx %s", err.Error())
		}
		if len(utxo.Balance.Address) < 1 {
			ts.Rollback()
			return "", fmt.Errorf("Fail to extract balance record from utxo relationship")
		}

		utxoAmountDecimal := decimal.NewFromFloat(utxo.Amount)
		balanceDecimalBig, result := new(big.Int).SetString(strconv.FormatUint(utxo.Balance.Decimal, 10), 10)
		if !result {
			return "", fmt.Errorf("Fail to convert decimal")
		}

		// construct balanceWithdrawTx by balance and utxo
		utxoAmountDecimal = utxoAmountDecimal.Mul(decimal.NewFromBigInt(balanceDecimalBig, 0))
		balanceWithdrawTx.Lock()
		balanceWithdrawTx.m[utxo.Balance.ID] = balanceWithdrawTx.m[utxo.Balance.ID].Add(utxoAmountDecimal)
		balanceWithdrawTx.Unlock()
	}

	for k, v := range balanceWithdrawTx.m {
		balance := model.Balance{}
		txRecord := model.Tx{}
		ts.First(&balance, k)
		ts.FirstOrCreate(&txRecord,
			model.Tx{
				TxID:      txid,
				TxType:    model.TxTypeWithdraw,
				Address:   balance.Address,
				Asset:     balance.Symbol,
				Amount:    v.String(),
				BalanceID: balance.ID,
				ChainName: model.ChainBitcoin,
			})
		for _, vout := range msgTx.TxOut {
			pkScriptHex := hex.EncodeToString(vout.PkScript)
			if strings.Contains(pkScriptHex, "6f6d6e69") {
				omniPropertyID := common.Hex2int(pkScriptHex[26:34])
				omniTransferAmount := common.Hex2int(pkScriptHex[36:])
				omnibalance := model.Balance{}
				omnitxRecord := model.Tx{}
				ts.Where("address = ? AND identify = ?", balance.Address, strconv.FormatInt(omniPropertyID, 10)).First(&omnibalance)
				ts.FirstOrCreate(&omnitxRecord,
					model.Tx{
						TxID:      txid,
						TxType:    model.TxTypeWithdraw,
						Address:   omnibalance.Address,
						Asset:     omnibalance.Symbol,
						Amount:    strconv.FormatInt(omniTransferAmount, 10),
						BalanceID: omnibalance.ID,
						ChainName: model.ChainBitcoin,
					})
			}
		}
	}
	if err := ts.Commit().Error; err != nil {
		ts.Rollback()
		return "", fmt.Errorf("Fail to add withdraw tx for balance record %s", err.Error())
	}
	return txid, nil
}

func (repo *repo) constructBitcoincoreTxParamsValidate(ctx context.Context, from, to, amount, symbol string) (
	fromP2AS, toP2AS []byte,
	ledgerInfo *btcjson.GetBlockChainInfoResult, estimateFeeResult *btcjson.EstimateSmartFeeResult,
	transferAmountBig *big.Int, balance *model.Balance, err error) {
	// validate and convert sender to pay2addressscript
	fromAddress, err := btcutil.DecodeAddress(from, &chaincfg.RegressionNetParams)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate Invalid From address:" + err.Error())
	}
	fromP2AS, err = txscript.PayToAddrScript(fromAddress)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate Fail to convert sender Address to p2as %s", err.Error())
	}

	// validate and convert receiver to pay2addressscript
	toAddress, err := btcutil.DecodeAddress(to, &chaincfg.RegressionNetParams)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate Invalid To address:" + err.Error())
	}
	toP2AS, err = txscript.PayToAddrScript(toAddress)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate Fail to convert receiver Address to p2as %s", err.Error())
	}

	// query sender balance record in db
	balance = &model.Balance{}
	if err := repo.db.Where("address = ? AND symbol = ?", from, symbol).First(&balance).Error; err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	balanceAmount, result := new(big.Int).SetString(balance.Amount, 10)
	if !result {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate: Fail to extract amount from balance")
	}

	// convert transfer amount
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate: Fail to parse float %s", err.Error())
	}
	transferAmount, err := btcutil.NewAmount(amountFloat)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate: Transfer amount %s", err.Error())
	}
	transferAmountBig = big.NewInt(int64(transferAmount))

	// calculate sender usable balance's amount
	withdrawLockAmount, result := new(big.Int).SetString(balance.WithdrawLock, 10)
	if !result {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate: Fail to extract withdrawLock from balance")
	}
	balanceAmount = balanceAmount.Sub(balanceAmount, withdrawLockAmount)

	// validate balance'amount is whether sufficient transferAmount
	if symbol == model.BTCSymbol && balanceAmount.Cmp(transferAmountBig) <= 0 {
		// btc transfer
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate: Insufficient balance")
	} else if symbol != model.BTCSymbol && balanceAmount.Cmp(transferAmountBig) < 0 {
		// omni transfer
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("constructBitcoincoreTxParamsValidate: Insufficient balance")
	}

	// chain info
	ledgerInfo, err = repo.QueryBTCLedgerInfo(ctx)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	estimateFeeResult, err = repo.bitcoinClient.EstimateSmartFee(6)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("EstimateSmartFee %s", err.Error())
	}
	return fromP2AS, toP2AS, ledgerInfo, estimateFeeResult, transferAmountBig, balance, nil
}
