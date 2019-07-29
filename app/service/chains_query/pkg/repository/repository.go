package repository

import(
  "github.com/btcsuite/btcd/rpcclient"
  "github.com/ethereum/go-ethereum/ethclient"
)

// Repo repo obj
type repo struct {
  bitcoinCLient *rpcclient.Client
  ethClient      *ethclient.Client
}

// New new repo
func New(btcClient *rpcclient.Client, ethClient *ethclient.Client) IBiz {
  repo := repo{
    bitcoinCLient: btcClient,
    ethClient: ethClient,
  }
  var biz IBiz = &repo
  return biz
}

func (repo *repo) close() error {
  return nil
}
