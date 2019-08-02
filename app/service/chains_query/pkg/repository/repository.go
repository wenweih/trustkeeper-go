package repository

import(
  log "github.com/go-kit/kit/log"
  "github.com/jinzhu/gorm"
  "github.com/qor/transition"
  "github.com/btcsuite/btcd/rpcclient"
  "github.com/ethereum/go-ethereum/ethclient"
  "trustkeeper-go/library/mq"
  "trustkeeper-go/app/service/chains_query/pkg/model"
)

// Repo repo obj
type repo struct {
  bitcoinClient *rpcclient.Client
  omniClient *rpcclient.Client
  ethClient     *ethclient.Client
  MQ            *mq.MessagingClient
  db            *gorm.DB
  logger        log.Logger
}

// New new repo
func New(
  btcClient *rpcclient.Client,
  omniClient *rpcclient.Client,
  ethClient *ethclient.Client,
  mq *mq.MessagingClient,
  db *gorm.DB,
  logger log.Logger) IBiz {
  db.AutoMigrate(
    model.BtcUtxo{},
    model.BtcBlock{},
    model.Balance{},
    &transition.StateChangeLog{},
  )
  repo := repo{
    bitcoinClient: btcClient,
    omniClient: omniClient,
    ethClient: ethClient,
    MQ: mq,
    db: db,
    logger: logger}
  var biz IBiz = &repo
  return biz
}

func (repo *repo) close() error {
  return nil
}
