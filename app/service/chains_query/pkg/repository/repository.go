package repository

import(
  "github.com/btcsuite/btcd/rpcclient"
  "github.com/ethereum/go-ethereum/ethclient"
  "trustkeeper-go/library/mq"
)

// Repo repo obj
type repo struct {
  bitcoinCLient *rpcclient.Client
  ethClient     *ethclient.Client
  MQ            *mq.MessagingClient
}

// New new repo
func New(btcClient *rpcclient.Client, ethClient *ethclient.Client, mq *mq.MessagingClient) IBiz {
  repo := repo{
    bitcoinCLient: btcClient,
    ethClient: ethClient,
    MQ: mq,
  }
  var biz IBiz = &repo
  return biz
}

func (repo *repo) close() error {
  return nil
}
