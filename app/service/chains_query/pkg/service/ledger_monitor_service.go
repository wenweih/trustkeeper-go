package service

import (
  "context"
  "math/big"
  log "github.com/go-kit/kit/log"
  "github.com/streadway/amqp"
  "github.com/btcsuite/btcd/btcjson"
  "github.com/btcsuite/btcd/chaincfg/chainhash"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum"
  "trustkeeper-go/app/service/chains_query/pkg/configure"
)

// LedgerMonitorService describes the service.
type LedgerMonitorService interface {
	BitcoincoreBlock(ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)
	EthereumSubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
  EthereumBlock(ctx context.Context, blockNumber *big.Int) (*types.Block, error)
  MQPublish(msg []byte, exchangeName, exchangeType, bindingKey, queueName string) error
  MQSubscribe(
    exchangeName, exchangeType, queueName,
    bindingKey, consumerName string, handleFunc func(amqp.Delivery)) error
}

// NewLedgerMonitorService returns a ChainsQueryService with all of the expected middleware wired in.
func NewLedgerMonitorService(conf configure.Conf, logger log.Logger) (LedgerMonitorService, error) {
	bs, err := NewBasicChainsQueryService(conf, logger)
	if err != nil {
		return nil, err
	}

	var svc LedgerMonitorService = bs
	return svc, nil
}

func (b *basicChainsQueryService) EthereumSubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return b.biz.EthereumSubscribeNewHead(ctx, ch)
}

func (b *basicChainsQueryService) EthereumBlock(ctx context.Context, number *big.Int) (*types.Block, error) {
  return b.biz.EthereumBlock(ctx, number)
}


func (b *basicChainsQueryService) MQPublish(msg []byte, exchangeName, exchangeType, bindingKey, queueName string) error {
  return b.biz.MQPublish(msg, exchangeName, exchangeType, bindingKey, queueName)
}

func (b *basicChainsQueryService) MQSubscribe(exchangeName, exchangeType, queueName,
  bindingKey, consumerName string, handleFunc func(amqp.Delivery)) error {
    return b.biz.MQSubscribe(exchangeName, exchangeType, queueName, bindingKey, consumerName, handleFunc)
}
