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

func (repo *repo) ConstructTxBTC(ctx context.Context, from, to, amount string) (string, int64, error) {
  fromAddress, err := btcutil.DecodeAddress(from, &chaincfg.RegressionNetParams)
  if err != nil {
    return "", 0, fmt.Errorf("Invalid From address:" + err.Error())
  }
  fromPkScript, err := txscript.PayToAddrScript(fromAddress)
  if err != nil {
    return "", 0, err
  }

  toAddress, err := btcutil.DecodeAddress(to, &chaincfg.RegressionNetParams)
  if err != nil {
    return "", 0, fmt.Errorf("Invalid To address:" + err.Error())
  }
  toPkScript, err := txscript.PayToAddrScript(toAddress)
  if err != nil {
    return "", 0, err
  }

  balance := model.Balance{}
  if err := repo.db.Where("address = ? AND symbol = ?", from, model.BTCSymbol).First(&balance).Error; err != nil {
    return "", 0, err
  }

  balanceAmount, result := new(big.Int).SetString(balance.Amount, 10)
  if !result {
    return "", 0, fmt.Errorf("fail to extract amount from balance")
  }
  amountFloat, err := strconv.ParseFloat(amount, 64)
  if err != nil {
    return "", 0, fmt.Errorf("fail to parse float %s", err.Error())
  }
  transferAmount, err := btcutil.NewAmount(amountFloat)
  if err != nil {
    return "", 0, fmt.Errorf("Transfer amount %s", err.Error())
  }
  transferAmountBig := big.NewInt(int64(transferAmount))
  if balanceAmount.Cmp(transferAmountBig) <= 0 {
    return "", 0, fmt.Errorf("Insufficient balance")
  }
  var txids []string
  if err := repo.db.Model(&model.Tx{State: model.StateSuccess,
    ChainName: model.ChainBitcoin,
    BalanceID: balance.ID,
    Asset: model.BTCSymbol}).Pluck("tx_id", &txids).Error;
    err != nil {
      return "", 0, fmt.Errorf("Fail to query comfirmations btc tx:" + err.Error())
    }
  var matureUTXOs []model.BtcUtxo
  if err := repo.db.Where("txid IN (?) AND balance_id = ? AND re_org = ? AND state = ?",
    txids, balance.ID, false, model.UTXOStateUnSelected).Find(&matureUTXOs).Error; err != nil{
    return "", 0, fmt.Errorf("fail to query mature utxos:" + err.Error())
  }

  ledgerInfo, err := repo.QueryBTCLedgerInfo(ctx)
  if err != nil {
    return "", 0, err
  }
  feeKB, err := repo.bitcoinClient.EstimateSmartFee(int64(6))
  if err != nil {
    return "", 0, err
  }
  feeRate := mempool.SatoshiPerByte(feeKB.FeeRate)

  if feeKB.FeeRate <= 0 {
    feeRate = mempool.SatoshiPerByte(100)
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
  fee := feeRate.Fee(uint32(msgTx.SerializeSize() + 181 + 34))
  if (vinAmount - transferAmountBig.Int64() - int64(fee)) > 0 {
    txOutReCharge := wire.NewTxOut((vinAmount - transferAmountBig.Int64() - int64(fee)), fromPkScript)
    msgTx.AddTxOut(txOutReCharge)
  }else {
    selectedutxoForFee, _, selectedCoinsForFee, err := CoinSelect(int64(ledgerInfo.Headers), fee, unselectedutxos)
    if err != nil {
      return "", 0, fmt.Errorf("Select UTXO for fee %s", err)
    }
    for _, coin := range selectedCoinsForFee.Coins() {
      vinAmount += int64(coin.Value())
    }
    txOutReCharge := wire.NewTxOut((vinAmount - transferAmountBig.Int64() - int64(fee)), fromPkScript)
    msgTx.AddTxOut(txOutReCharge)
    selectedutxos = append(selectedutxos, selectedutxoForFee...)
  }

  utxoids := make([]uint, len(selectedutxos))
  for i, selectedUTXO := range selectedutxos {
    utxoids[i] = selectedUTXO.ID
  }

  if err := repo.db.Table("btc_utxos").Where("id IN (?)", utxoids).UpdateColumn("state", model.UTXOStateLocked).Error; err != nil {
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
    return "", fmt.Errorf("Bitcoin SendRawTransaction %s", err)
  }
  txid := txHash.String()
  ts := repo.db.Begin()
  for _, vin := range tx.MsgTx().TxIn {
    utxo := model.BtcUtxo{}
    txRecord := model.Tx{}
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
    ts.FirstOrCreate(&txRecord,
      model.Tx{
        TxID: txid,
        TxType: model.TxTypeWithdraw,
        Address: utxo.Balance.Address,
        Asset: utxo.Balance.Symbol,
        Amount: strconv.FormatFloat(utxo.Amount * btcutil.SatoshiPerBitcoin, 'f', -int(0), 64),
        BalanceID: utxo.Balance.ID,
        ChainName: model.ChainBitcoin,
    })
  }
  if err := ts.Commit().Error; err != nil {
    ts.Rollback()
    return "", fmt.Errorf("Fail to operate db when send bitcoincore tx %s", err.Error())
  }
  return txid, nil
}
