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
  "trustkeeper-go/app/service/chains_query/pkg/repository"
  "trustkeeper-go/app/service/chains_query/pkg/model"
)

// LedgerMonitorService describes the service.
type LedgerMonitorService interface {
  QueryBTCLedgerInfo(ctx context.Context) (*btcjson.GetBlockChainInfoResult, error)
	BitcoincoreBlock(
    ctx context.Context, blockHash *chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)
  CreateBTCBlockWithUTXOs(
    ctx context.Context, rawBlock *btcjson.GetBlockVerboseResult) (<-chan repository.CreateBlockResult)
  CreateBtcBlockWithUtxoPipeline(
    ctx context.Context, height int64) (<-chan repository.CreateBlockResult)
  TrackBtcBlockPipeline(
    ctx context.Context, trackHeight, bestBlockHeight int64, isTracking bool) (bool, int64)
	EthereumSubscribeNewHead(
    ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
  EthereumBlock(
    ctx context.Context, blockNumber *big.Int) (*types.Block, error)
  MQPublish(msg []byte, exchangeName, exchangeType, bindingKey string) error
  MQSubscribe(
    exchangeName, exchangeType, queueName,
    bindingKey, consumerName string, handleFunc func(amqp.Delivery)) error

  CreateETHBlockWithTx(ctx context.Context, height int64) (*model.EthBlock, error)
  EthereumBestBlock(ctx context.Context) (*types.Block, error)
  EthereumDBBestBlock(ctx context.Context) (*model.EthBlock, error)

  UpdateEthereumTx(ctx context.Context)
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

func (b *basicChainsQueryService) EthereumSubscribeNewHead(
  ctx context.Context,
  ch chan<- *types.Header) (ethereum.Subscription, error) {
	return b.biz.EthereumSubscribeNewHead(ctx, ch)
}

func (b *basicChainsQueryService) EthereumBlock(
  ctx context.Context, number *big.Int) (*types.Block, error) {
  return b.biz.EthereumBlock(ctx, number)
}


func (b *basicChainsQueryService) MQPublish(
  msg []byte, exchangeName, exchangeType, bindingKey string) error {
  return b.biz.MQPublish(msg, exchangeName, exchangeType, bindingKey)
}

func (b *basicChainsQueryService) MQSubscribe(exchangeName, exchangeType, queueName,
  bindingKey, consumerName string, handleFunc func(amqp.Delivery)) error {
    return b.biz.MQSubscribe(
      exchangeName, exchangeType, queueName, bindingKey, consumerName, handleFunc)
}

func (b *basicChainsQueryService) QueryBTCLedgerInfo(
  ctx context.Context) (*btcjson.GetBlockChainInfoResult, error) {
  return b.biz.QueryBTCLedgerInfo(ctx)
}

func (b *basicChainsQueryService) CreateBtcBlockWithUtxoPipeline(
  ctx context.Context,
  height int64) (<-chan repository.CreateBlockResult) {
  return b.biz.CreateBtcBlockWithUtxoPipeline(ctx, height)
}

func (b *basicChainsQueryService) TrackBtcBlockPipeline(
  ctx context.Context,
  trackHeight, bestBlockHeight int64, isTracking bool) (bool, int64) {
  ch := b.biz.QueryBTCBlockCH(ctx, trackHeight)
  return b.biz.TrackBlock(ctx, bestBlockHeight, isTracking, ch)
}

func (b *basicChainsQueryService) CreateBTCBlockWithUTXOs(
  ctx context.Context, rawBlock *btcjson.GetBlockVerboseResult) (<-chan repository.CreateBlockResult) {
  blockCh := make(chan repository.UTXOBlockResult)
  go func (rawBlock *btcjson.GetBlockVerboseResult)  {
    defer close(blockCh)
    blockCh  <- repository.UTXOBlockResult{Block: rawBlock}
  }(rawBlock)
  return b.biz.CreateBTCBlockWithUTXOs(ctx, blockCh)
}

func (b *basicChainsQueryService) CreateETHBlockWithTx(ctx context.Context, height int64) (*model.EthBlock, error) {
  return b.biz.CreateETHBlockWithTx(ctx, height)
}

func (b *basicChainsQueryService) EthereumBestBlock(ctx context.Context) (*types.Block, error) {
  return b.biz.EthereumBestBlock(ctx)
}

func (b *basicChainsQueryService) EthereumDBBestBlock(ctx context.Context) (*model.EthBlock, error) {
  return b.biz.EthererumDBBestBlock(ctx)
}

func (b *basicChainsQueryService) UpdateEthereumTx(ctx context.Context) {
  b.biz.UpdateEthereumTx(ctx)
}
