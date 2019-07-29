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
}

func (repo *repo) QueryBitcoincoreBlock (ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error){
  return repo.bitcoinCLient.GetBlockVerbose(blockHash)
}

func (repo *repo) EthereumSubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
  return repo.ethClient.SubscribeNewHead(ctx, ch)
}
