package repository

import (
  "fmt"
  "bytes"
  "context"
  "strings"
  "math/big"
  "encoding/gob"
  "encoding/hex"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/common/hexutil"
  "trustkeeper-go/app/service/chains_query/pkg/model"
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
    Hash: simpleBlock.Hash.String(),
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
        balance model.Balance
      )
      if err := ts.Where("address = ? AND identify = ?",
        receiver.Hex(), strings.ToLower(tx.To)).First(&balance).Error;
        err != nil && err.Error() != "record not found" {
        ts.Rollback()
        return nil, fmt.Errorf("Fail to query erc20 balance record, address:" + receiver.Hex() + "identify:" + tx.To + err.Error())
      } else if (err != nil && err.Error() == "record not found"){
        continue
      }
      ts.FirstOrCreate(&txRecord, model.Tx{
        TxID: tx.Txid,
        TxType: "deposit",
        Address: receiver.Hex(),
        Asset: balance.Symbol,
        Amount: amount.String(),
        BalanceID: balance.ID,
        State: model.StateConfirming,
        ChainName: model.ChainEthereum})
    } else {
      var (
        txRecord model.Tx
        balance model.Balance
      )
      if err := ts.Where("address = ? AND Symbol = ?", tx.To, "ETH").First(&balance).Error;
      err != nil && err.Error() != "record not found" {
        ts.Rollback()
        return nil, fmt.Errorf("Fail to query eth balance record, address:" + tx.To + " " + err.Error())
      } else if (err != nil && err.Error() == "record not found") {
        continue
      }
      amountBig, err := hexutil.DecodeBig(tx.ValueHex)
      if err != nil {
        return nil, err
      }

      ts.FirstOrCreate(&txRecord, model.Tx{
        TxID: tx.Txid,
        TxType: "deposit",
        Address: tx.To,
        Asset: balance.Symbol,
        Amount: amountBig.String(),
        BalanceID: balance.ID,
        State: model.StateConfirming,
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
    repo.logger.Log("UpdateEthereumTx:", err.Error())
  }
  bestBlockHead, err := repo.ethClient.HeaderByNumber(ctx, nil)
  if err != nil {
    repo.logger.Log("UpdateEthereumTx:", err.Error())
  }
  ts := repo.db.Begin()
  for _, tx := range txes {
    receipt, err := repo.QueryEthereumTxReceipt(ctx, tx.TxID)
    if err != nil {
      repo.logger.Log("UpdateEthereumTx:", err.Error())
    }
    if receipt.BlockNumber == nil {
      // blockNumber field will be null until the transaction is included into a mined block
      repo.logger.Log("UpdateEthereumTx", tx.TxID, "StateFrom", tx.State, "To", model.StatePending)
      ts.Model(&tx).UpdateColumn("state", model.StatePending)
    } else if receipt.Status == 1 {
      // Since block 4370000 (Byzantium), a status indicator has been added to receipts. 1 mean success, 0 mean fail
      confirmations := bestBlockHead.Number.Int64() - receipt.BlockNumber.Int64() + 1
      var state string
      if confirmations >= DepositEthereumComfirmation {
        state = model.StateSuccess
      }else {
        state = model.StateConfirming
      }
      repo.logger.Log("UpdateEthereumTx", tx.TxID, "StateFrom", tx.State, "From", state, "ConfirmationsFrom", tx.Confirmations, "To", confirmations)
      ts.Model(&tx).UpdateColumns(model.Tx{State: state, Confirmations: confirmations})
    } else if (receipt.Status == 0) {
      repo.logger.Log("UpdateEthereumTx", tx.TxID, "StateFrom", tx.State, "To", model.StateFail)
      ts.Model(&tx).UpdateColumn("state", model.StateFail)
    }
  }
  if err := ts.Commit().Error; err != nil {
    repo.logger.Log("UpdateEthereumTx:", err.Error())
  }
}
