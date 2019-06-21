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

func (l loggingMiddleware) GenerateMnemonic(ctx context.Context, namespaceid string, bip44ids []int32, bip44accountSize int) (xpubs []*Bip44ThirdXpubsForChain, err error) {
	defer func() {
		l.logger.Log("method", "GenerateMnemonic", "namespaceid", namespaceid, "err", err)
	}()
	return l.next.GenerateMnemonic(ctx, namespaceid, bip44ids, bip44accountSize)
}

func (l loggingMiddleware) Close() error {
	defer func() {
		l.logger.Log("method", "Close", "close resource", "(database, redis etc...)")
	}()
	return l.next.Close()
}
