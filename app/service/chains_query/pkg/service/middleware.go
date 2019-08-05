package service

import (
	"context"
	"trustkeeper-go/app/service/chains_query/pkg/repository"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(ChainsQueryService) ChainsQueryService

type loggingMiddleware struct {
	logger log.Logger
	next   ChainsQueryService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ChainsQueryService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ChainsQueryService) ChainsQueryService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) BitcoincoreBlock(ctx context.Context, blockHash *chainhash.Hash) (b0 *btcjson.GetBlockVerboseResult, e1 error) {
	defer func() {
		l.logger.Log("method", "BitcoincoreBlock", "blockHash", blockHash, "b0", b0, "e1", e1)
	}()
	return l.next.BitcoincoreBlock(ctx, blockHash)
}

func (l loggingMiddleware) QueryOmniProperty(ctx context.Context, propertyid int64) (r0 *repository.OmniProperty, e1 error) {
	defer func() {
		l.logger.Log("method", "QueryOmniProperty", "propertyid", propertyid, "e1", e1)
	}()
	return l.next.QueryOmniProperty(ctx, propertyid)
}

func (l loggingMiddleware) ERC20TokenInfo(ctx context.Context, tokenHex string) (token *repository.ERC20Token, err error) {
	defer func() {
		l.logger.Log("method", "ERC20TokenInfo", "tokenHex", tokenHex, "token", token, "err", err)
	}()
	return l.next.ERC20TokenInfo(ctx, tokenHex)
}
