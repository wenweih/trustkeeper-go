package repository

import (
  "context"
  "github.com/btcsuite/btcd/chaincfg/chainhash"
  "github.com/btcsuite/btcd/btcjson"
)

type IBiz interface {
  QueryBitcoincoreBlock(ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)
}

func (repo *repo) QueryBitcoincoreBlock (ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error){
  return repo.bitcoinCLient.GetBlockVerbose(blockHash)
}
