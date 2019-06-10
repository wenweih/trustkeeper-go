package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(WalletManagementService) WalletManagementService

type loggingMiddleware struct {
	logger log.Logger
	next   WalletManagementService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a WalletManagementService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next WalletManagementService) WalletManagementService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateChain(ctx context.Context, symbol string, bit44ID string, status bool) (err error) {
	defer func() {
		l.logger.Log("method", "CreateChain", "symbol", symbol, "bit44ID", bit44ID, "status", status, "err", err)
	}()
	return l.next.CreateChain(ctx, symbol, bit44ID, status)
}
