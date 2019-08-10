package model

import(
  "github.com/jinzhu/gorm"
  "bytes"
  "math/big"
  "encoding/gob"
  "github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
)

// EthBlock notify block info
type EthBlock struct {
  gorm.Model
  Hash    string    `gorm:"not null;unique_index:idx_hash_height"`
  Height  int64     `gorm:"type:bigint;unique_index:idx_hash_height;not null"`
  ReOrg   bool      `gorm:"default:false"`
}

// ETHSimpleTx ethereum transaction
type ETHSimpleTx struct {
	THash     string `json:"thash"`
	From      string `json:"from"`
	To        string `json:"to"`
	Txid      string `json:"txid"`
	HeightHex string `json:"height"`
	ValueHex  string `json:"value"`
	FeeHex    string `json:"fee"`
  Data      string `json:"data"`
}

// EthereumBlock custom ethereum block struct
type EthereumBlock struct {
  Hash   common.Hash
  Header *types.Header
  Tx     []*ETHSimpleTx
}

func EncodeETHBlock(block types.Block) ([]byte, error) {
  buf := new(bytes.Buffer)
  buf.Reset()

  txs := block.Transactions()
	txes := make([]*ETHSimpleTx, 0)
	for _, tx := range txs {
		ms, _ := tx.AsMessage(types.NewEIP155Signer(big.NewInt(1)))
		var to string
		pto := tx.To()
		if pto != nil {
			to = (*pto).Hex()
		}
		var txFee = new(big.Int)
		txFee = txFee.Mul(tx.GasPrice(), big.NewInt(int64(tx.Gas())))
		txes = append(txes, &ETHSimpleTx{
			THash:     tx.Hash().String(),
			To:        to,
			From:      ms.From().String(),
			HeightHex: hexutil.EncodeBig(block.Number()),
			ValueHex:  hexutil.EncodeBig(tx.Value()),
			FeeHex:    hexutil.EncodeBig(txFee),
			Txid:      tx.Hash().String(),
      Data:      hexutil.Encode(tx.Data()),
		})
	}

  ethereumBlock := EthereumBlock{
    Hash: block.Hash(),
    Header: block.Header(),
    Tx: txes,
  }
  e := gob.NewEncoder(buf)
  if err := e.Encode(ethereumBlock); err != nil {
    return nil, err
  }
  return buf.Bytes(), nil
}
