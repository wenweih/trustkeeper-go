package service

import (
  "context"
  log "github.com/go-kit/kit/log"
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
