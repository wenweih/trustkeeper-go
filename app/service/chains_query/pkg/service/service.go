package service

import (
	"context"
	"errors"
	"strings"
	"trustkeeper-go/app/service/chains_query/pkg/configure"
	"trustkeeper-go/app/service/chains_query/pkg/repository"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	log "github.com/go-kit/kit/log"

	"github.com/ethereum/go-ethereum/ethclient"

	"trustkeeper-go/library/mq"
)

// ChainsQueryService describes the service.
type ChainsQueryService interface {
	BitcoincoreBlock(ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)
}

type basicChainsQueryService struct {
	biz repository.IBiz
}

// NewBasicChainsQueryService returns a naive, stateless implementation of ChainsQueryService.
func NewBasicChainsQueryService(conf configure.Conf, logger log.Logger) (*basicChainsQueryService, error) {
	btcclient, err := rpcclient.New(conf.BTCconnCfg, nil)
	if err != nil {
		return nil, errors.New(strings.Join([]string{"rpcclient error", err.Error()}, ":"))
	}

	ethereumClient, err := ethclient.Dial(conf.ETHRPC)
	if err != nil {
		return nil, errors.New(strings.Join([]string{"Ethereum client error", err.Error()}, ":"))
	}

	messageClient := &mq.MessagingClient{}
	if err := messageClient.ConnectToBroker(conf.MQ); err != nil {
		return nil, errors.New(strings.Join([]string{"mq connection to broker error", err.Error()}, ":"))
	}

	return &basicChainsQueryService{
		biz: repository.New(btcclient, ethereumClient, messageClient),
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
