package repository

import (
  "context"
  "github.com/btcsuite/btcd/chaincfg/chainhash"
  "github.com/btcsuite/btcd/btcjson"

  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum"
)

type IBiz interface {
  QueryBitcoincoreBlock(ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)
  EthereumSubscribeNewHead(ctc context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
  MQPublish(msg []byte, exchangeName, exchangeType, bindingKey, queueName string) error
}

func (repo *repo) QueryBitcoincoreBlock (ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error){
  return repo.bitcoinCLient.GetBlockVerboseTxM(blockHash)
}

func (repo *repo) EthereumSubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
  return repo.ethClient.SubscribeNewHead(ctx, ch)
}

func (repo *repo) MQPublish(msg []byte, exchangeName, exchangeType, bindingKey, queueName string) error {
  return repo.MQ.Publish(msg, exchangeName, exchangeType, bindingKey, queueName)
}
