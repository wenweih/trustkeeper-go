package repository

import (
  "fmt"
  "time"
  "strings"
  "context"
)

// IBiz repository bussiness logic
type IBiz interface {
  SaveMnemonic(namespaceID string, mnemonic []byte) (version string, err error)
  Close() error
  SignedBitcoincoreTx(ctx context.Context, walletHD WalletHD, txHex string, vinAmount int64) (signedTxHex string, err error)
  SignedEthereumTx(ctx context.Context, walletHD WalletHD, txHex string, chainID string) (string, error)
}

// WalletHD wallet hd info
type WalletHD struct {
  CoinType        int32  `json:"CoinType"`
  Account         int32  `json:"Account"`
  Change          int32  `json:"Change"`
  AddressIndex    uint32  `json:"AddressIndex"`
  MnemonicVersion string `json:"MnemonicVersion"`
}

func (repo *repo) Close() error{
  return repo.close()
}

func (repo * repo)SaveMnemonic(namespaceID string, mnemonic []byte) (string ,error) {
  // https://stackoverflow.com/questions/33119748/convert-time-time-to-string
  t := time.Now().Format("2006-01-02 15:04:05.000000")
  version := strings.Join([]string{namespaceID, t}, "/")
  err := repo.ldb.Put([]byte(version), mnemonic, nil)
  if err != nil {
    return "", fmt.Errorf("Save privite key to leveldb %s", err)
  }
  return version, nil
}
