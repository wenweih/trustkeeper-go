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
	ERC20TokenInfo(ctx context.Context, tokenHex string) (token *repository.ERC20Token, err error)

	ConstructTxBTC(ctx context.Context, from, to, amount string) (unsignedTxHex string, vinAmount int64, err error)
	SendBTCTx(ctx context.Context, signedTxHex string) (txID string, err error)

	ConstructTxETH(ctx context.Context, from, to, amount string) (unsignedTxHex, chainID string, err error)

	QueryBalance(ctx context.Context, symbol, address string) (balance string, err error)
	WalletValidate(ctx context.Context, chainName, address string) (err error)
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

	ethereumClient, err := ethclient.Dial(conf.ETHWS)
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
		biz: repository.New(btcclient, omniClient, ethereumClient, messageClient, db, logger, conf),
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

func (b *basicChainsQueryService) ERC20TokenInfo(ctx context.Context, tokenHex string) (*repository.ERC20Token, error) {
	return b.biz.ERC20TokenInfo(ctx, tokenHex)
}

func (b *basicChainsQueryService) ConstructTxBTC(ctx context.Context,
	from string, to string, amount string) (unsignedTxHex string, vinAmount int64, err error) {
	unsignedTxHex, vinAmount, err = b.biz.ConstructTxBTC(ctx, from, to, amount)
	return
}

func (b *basicChainsQueryService) SendBTCTx(ctx context.Context, signedTxHex string) (txID string, err error) {
	txID, err = b.biz.SendBTCTx(ctx, signedTxHex)
	return
}

func (b *basicChainsQueryService) QueryBalance(ctx context.Context, symbol string, address string) (balance string, err error) {
	balance, err = b.biz.QueryBalance(ctx, symbol, address)
	return
}

func (b *basicChainsQueryService) WalletValidate(ctx context.Context, chainName string, address string) (err error) {
	err = b.biz.WalletValidate(ctx, chainName, address)
	return
}

func (b *basicChainsQueryService) ConstructTxETH(ctx context.Context, from string, to string, amount string) (string, string, error) {
	return b.biz.ConstructTxETH(ctx, from, to, amount)
}
