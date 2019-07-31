package main

import(
  "fmt"
  "bytes"
  "strconv"
  "errors"
  "strings"
  "context"
  "math/big"
  "encoding/gob"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/common/hexutil"
)

func subHandle(orderHeight *big.Int, head *types.Header) (*big.Int, error) {
	ctx := context.Background()
	number := head.Number
  originBlock, err := svc.EthereumBlock(ctx, head.Number)
	if err != nil {
    e := errors.New(strings.Join([]string{"FailToGetOriginBlock ", err.Error(), " Height ", number.String()}, ""))
    return big.NewInt(head.Number.Int64() - 1), e
	}

	if orderHeight.Cmp(big.NewInt(0)) == 0 {
		orderHeight = originBlock.Number()
	}
  logger.Log(
    "orderHeight",
    orderHeight.Int64(),
    "SubscriptBlockHeight",
    originBlock.Number().Int64(),
    "OriginBlockHash", originBlock.Hash().String())

	for blockNumber := orderHeight.Int64();
  blockNumber <= originBlock.Number().Int64();
  blockNumber++ {
    block, err := svc.EthereumBlock(ctx,  big.NewInt(blockNumber))
		if err != nil {
      e := errors.New(strings.Join([]string{
        "Get block error, height:",
        strconv.FormatInt(blockNumber, 10)}, ""))
			return big.NewInt(blockNumber), e
		}
    data, err := encodeBlock(*block)
    if err != nil {
      fmt.Println(err.Error())
      e := errors.New(strings.Join([]string{"json Marshal raw ethereum block error", err.Error()}, ""))
    	return big.NewInt(blockNumber), e
    }
		if err := svc.MQPublish(
      data,
      "bestblock",
      "direct",
      "ethereum",
      "ethereum_best_block_queue");
      err != nil {
        e := errors.New(strings.Join([]string{"EtherumPublishError", err.Error()}, ""))
        return big.NewInt(blockNumber), e
      }
    orderHeight.Add(orderHeight, big.NewInt(1))
  }
  return orderHeight, nil
}

func encodeBlock(block types.Block) ([]byte, error) {
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

// EthereumBlock custom ethereum block struct
type EthereumBlock struct {
  Hash   common.Hash
  Header *types.Header
  Tx     []*ETHSimpleTx
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
}
