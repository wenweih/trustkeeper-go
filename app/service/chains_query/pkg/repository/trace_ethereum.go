package repository

import (
  "fmt"
  "bytes"
  "context"
  "strings"
  // "strconv"
  "math/big"
  "encoding/gob"
  "encoding/hex"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/common/hexutil"
  // "github.com/ethereum/go-ethereum/common/math"
  "trustkeeper-go/app/service/chains_query/pkg/model"
)

func (repo *repo) CreateETHBlockWithTx(ctx context.Context, height int64) (error) {
  number := new(big.Int)
  number.SetInt64(height)
  block, err := repo.EthereumBlock(ctx, number)
  if err != nil {
    return err
  }
  blockBytes, err := model.EncodeETHBlock(*block)
  if err != nil {
    return err
  }
  simpleBlock := model.EthereumBlock{}
  buf := bytes.NewBuffer(blockBytes)
  dc := gob.NewDecoder(buf)
  if err := dc.Decode(&simpleBlock); err != nil {
    return err
  }
  dbBlock := model.EthBlock{}
  ts := repo.db.Begin()
  if err := ts.FirstOrCreate(&dbBlock, model.EthBlock{
    Hash: simpleBlock.Hash.String(),
    Height: simpleBlock.Header.Number.Int64(),
  }).Error; err != nil {
    ts.Rollback()
    return fmt.Errorf("create eth block error: %s", err)
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
        return err
      }
      amount := new(big.Int).SetBytes(amountBytes)
      var (
        txRecord model.Tx
        balance model.Balance
      )
      if err := ts.Where("address = ? AND identify = ?", receiver.Hex(), strings.ToLower(tx.To)).First(&balance).Error; err != nil {
        return err
      }
      ts.FirstOrCreate(&txRecord, model.Tx{
        TxID: tx.Txid,
        TxType: "deposit",
        Address: receiver.Hex(),
        Asset: balance.Symbol,
        Amount: amount.String(),
        BalanceID: balance.ID,
      })
    } else {
      var (
        txRecord model.Tx
        balance model.Balance
      )
      if err := ts.Where("address = ? AND Symbol = ?", tx.To, "ETH").First(&balance).Error; err != nil {
        return err
      }
      amountBig, err := hexutil.DecodeBig(tx.ValueHex)
      if err != nil {
        return err
      }

      ts.FirstOrCreate(&txRecord, model.Tx{
        TxID: tx.Txid,
        TxType: "deposit",
        Address: tx.To,
        Asset: balance.Symbol,
        Amount: amountBig.String(),
        BalanceID: balance.ID,
      })
    }
  }
  return ts.Commit().Error
}
