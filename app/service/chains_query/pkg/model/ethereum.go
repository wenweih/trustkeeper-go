package model

import (
	"bytes"
	"encoding/gob"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

const (
	ETHSymbol string = "ETH"
)

// EthBlock notify block info
type EthBlock struct {
	gorm.Model
	Hash   string `gorm:"not null;unique_index:idx_hash_height"`
	Height int64  `gorm:"type:bigint;unique_index:idx_hash_height;not null"`
	ReOrg  bool   `gorm:"default:false"`
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

// TxPoolInspect ethereum transaction pool datatype
type TxPoolInspect struct {
	Pending map[string]map[uint64]string `json:"pending"`
	Queued  map[string]map[uint64]string `json:"queued"`
}

func EncodeETHBlock(block types.Block) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Reset()
	txs := block.Transactions()
	txes := make([]*ETHSimpleTx, 0)
	for _, tx := range txs {
		ms, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
		if err != nil {
			return nil, err
		}
		var txFee = new(big.Int)
		txFee = txFee.Mul(tx.GasPrice(), big.NewInt(int64(tx.Gas())))
		inputeData := hexutil.Encode(ms.Data())
		to := ""
		if tx.To() != nil {
			to = tx.To().Hex()
		}
		txes = append(txes, &ETHSimpleTx{
			THash:     tx.Hash().String(),
			To:        to,
			From:      ms.From().String(),
			HeightHex: hexutil.EncodeBig(block.Number()),
			ValueHex:  hexutil.EncodeBig(tx.Value()),
			FeeHex:    hexutil.EncodeBig(txFee),
			Txid:      tx.Hash().String(),
			Data:      inputeData,
		})
	}

	ethereumBlock := EthereumBlock{
		Hash:   block.Hash(),
		Header: block.Header(),
		Tx:     txes,
	}
	e := gob.NewEncoder(buf)
	if err := e.Encode(ethereumBlock); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ToEther(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}

func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)
	wei := new(big.Int)
	wei.SetString(result.String(), 10)
	return wei
}

// EncodeETHTx encode eth tx
func EncodeETHTx(tx *types.Transaction) (string, error) {
	txb, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "", err
	}
	txHex := hexutil.Encode(txb)
	return txHex, nil
}

// DecodeETHTx ethereum transaction hex
func DecodeETHTx(txHex string) (*types.Transaction, error) {
	txc, err := hexutil.Decode(txHex)
	if err != nil {
		return nil, err
	}

	var txde types.Transaction
	t, err := &txde, rlp.Decode(bytes.NewReader(txc), &txde)
	if err != nil {
		return nil, err
	}
	return t, nil
}
