package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(WalletKeyService) WalletKeyService

type loggingMiddleware struct {
	logger log.Logger
	next   WalletKeyService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next WalletKeyService) WalletKeyService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GenerateMnemonic(ctx context.Context, uuid string) (xpub string, err error) {
	defer func() {
		l.logger.Log("method", "GenerateMnemonic", "uuid", uuid, "xpub", xpub, "err", err)
	}()
	return l.next.GenerateMnemonic(ctx, uuid)
}
