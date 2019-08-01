package repository

import (
  "context"
  "github.com/btcsuite/btcd/btcjson"
  "github.com/btcsuite/btcd/chaincfg/chainhash"
)

func (repo *repo) QueryBitcoincoreBlock (
  ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error){
  return repo.bitcoinClient.GetBlockVerboseTxM(blockHash)
}

func (repo *repo) GetBTCBlockHashByHeight(ctx context.Context, height int64) (*chainhash.Hash, error) {
  return repo.bitcoinClient.GetBlockHash(height)
}

func (repo *repo) QueryBTCLedgerInfo(ctx context.Context) (*btcjson.GetBlockChainInfoResult, error) {
  return repo.bitcoinClient.GetBlockChainInfo()
}
