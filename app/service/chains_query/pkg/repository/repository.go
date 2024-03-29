package repository

import (
	"trustkeeper-go/app/service/chains_query/pkg/configure"
	"trustkeeper-go/app/service/chains_query/pkg/model"
	"trustkeeper-go/library/mq"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

// Repo repo obj
type repo struct {
	bitcoinClient *rpcclient.Client
	omniClient    *rpcclient.Client
	ethClient     *ethclient.Client
	MQ            *mq.MessagingClient
	db            *gorm.DB
	logger        log.Logger
	conf          configure.Conf
}

// New new repo
func New(
	btcClient *rpcclient.Client,
	omniClient *rpcclient.Client,
	ethClient *ethclient.Client,
	mq *mq.MessagingClient,
	db *gorm.DB,
	logger log.Logger,
	conf configure.Conf) IBiz {
	db.AutoMigrate(
		model.BtcUtxo{},
		model.BtcBlock{},
		model.Balance{},
		model.Tx{},
		model.BalanceLog{},
		model.EthBlock{},
	)
	repo := repo{
		bitcoinClient: btcClient,
		omniClient:    omniClient,
		ethClient:     ethClient,
		MQ:            mq,
		db:            db,
		logger:        logger,
		conf:          conf,
	}
	repo.DeclareExChange(ExchangeNameForBitcoincoreBestBlock, "fanout")
	repo.DeclareExChange(ExchangeNameForEthereumBestBlock, "fanout")
	repo.DeclareQueue(QueueNameForBitcoincoreTraceTx)
	repo.DeclareQueue(QueueNameForBitcoincoreUpdateTx)
	repo.DeclareQueue(QueueNameForEthereumTraceTx)
	repo.DeclareQueue(QueueNameForEthereumUpdateTx)
	var biz IBiz = &repo
	return biz
}

func (repo *repo) close() error {
	return nil
}
