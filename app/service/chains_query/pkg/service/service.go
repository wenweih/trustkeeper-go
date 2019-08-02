package service

import (
	"context"
	"errors"
	"strings"
	"trustkeeper-go/app/service/chains_query/pkg/configure"
	"trustkeeper-go/app/service/chains_query/pkg/repository"
	"trustkeeper-go/library/database/orm"
	"trustkeeper-go/library/mq"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/go-kit/kit/log"
)

// ChainsQueryService describes the service.
type ChainsQueryService interface {
	BitcoincoreBlock(ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)
	QueryOmniProperty(ctx context.Context, propertyid int64) (*repository.OmniProperty, error)
}

type basicChainsQueryService struct {
	biz repository.IBiz
}

// NewBasicChainsQueryService returns a naive, stateless implementation of ChainsQueryService.
func NewBasicChainsQueryService(conf configure.Conf, logger log.Logger) (*basicChainsQueryService, error) {
	btcclient, err := rpcclient.New(conf.BTCconnCfg, nil)
	if err != nil {
		return nil, errors.New(strings.Join([]string{"btc rpcclient error", err.Error()}, ":"))
	}

	omniClient, err := rpcclient.New(conf.OmniconnCfg, nil)
	if err != nil {
		return nil, errors.New(strings.Join([]string{"omni rpcclient error", err.Error()}, ":"))
	}

	ethereumClient, err := ethclient.Dial(conf.ETHRPC)
	if err != nil {
		return nil, errors.New(strings.Join([]string{"Ethereum client error", err.Error()}, ":"))
	}

	messageClient := &mq.MessagingClient{}
	if err := messageClient.ConnectToBroker(conf.MQ); err != nil {
		return nil, errors.New(strings.Join([]string{"mq connection to broker error", err.Error()}, ":"))
	}

	db, err := orm.DB(conf.DBInfo)
	if err != nil {
		return nil, err
	}
	return &basicChainsQueryService{
		biz: repository.New(btcclient, omniClient, ethereumClient, messageClient, db, logger),
	}, nil
}

// New returns a ChainsQueryService with all of the expected middleware wired in.
func New(conf configure.Conf, logger log.Logger, middleware []Middleware) (ChainsQueryService, error) {
	bs, err := NewBasicChainsQueryService(conf, logger)
	if err != nil {
		return nil, err
	}

	var svc ChainsQueryService = bs
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc, nil
}

func (b *basicChainsQueryService) BitcoincoreBlock(ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error) {
	return b.biz.QueryBitcoincoreBlock(ctx, blockHash)
}

func (b *basicChainsQueryService) QueryOmniProperty(ctx context.Context, propertyid int64) (*repository.OmniProperty, error) {
	return b.biz.QueryOmniProperty(propertyid)
}
