package service

import (
	"context"
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

func (l loggingMiddleware) BitcoincoreBlock(ctx context.Context, blockHash string) (b0 *btcjson.GetBlockVerboseResult, e1 error) {
	defer func() {
		l.logger.Log("method", "BitcoincoreBlock", "blockHash", blockHash, "b0", b0, "e1", e1)
	}()
	return l.next.BitcoincoreBlock(ctx, blockHash)
}
