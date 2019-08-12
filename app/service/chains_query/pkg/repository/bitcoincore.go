package repository

import (
  "context"
  "github.com/btcsuite/btcd/btcjson"
  "github.com/btcsuite/btcd/chaincfg/chainhash"
)

const (
  // DepositBitcoincoreComfirmation more than DepositBitcoincoreComfirmation mean deposit successfully
  DepositBitcoincoreComfirmation uint64 = 6
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

func (repo *repo) QueryBTCTx(ctx context.Context, txID string) (*btcjson.TxRawResult, error) {
  hash, err := chainhash.NewHashFromStr(txID)
  if err != nil {
    return nil, err
  }
  tx, err := repo.bitcoinClient.GetRawTransactionVerbose(hash)
  if err != nil {
    return nil, err
  }
  return tx, nil
}
