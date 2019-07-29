package repository

import(
  "github.com/btcsuite/btcd/rpcclient"
)

// Repo repo obj
type repo struct {
  bitcoinCLient *rpcclient.Client
}

// New new repo
func New(btcClient *rpcclient.Client) IBiz {
  repo := repo{
    bitcoinCLient: btcClient,
  }
  var biz IBiz = &repo
  return biz
}

func (repo *repo) close() error {
  return nil
}
