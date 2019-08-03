package service

import (
	"context"
	"trustkeeper-go/app/service/transaction/pkg/repository"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(TransactionService) TransactionService

type loggingMiddleware struct {
	logger log.Logger
	next   TransactionService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TransactionService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TransactionService) TransactionService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Close() error {
	defer func() {
		l.logger.Log("method", "Close", "close resource", "(database, redis etc...)")
	}()
	return l.next.Close()
}

func (l loggingMiddleware) AssignAssetsToWallet(ctx context.Context, address string, assets []*repository.SimpleAsset) (err error) {
	defer func() {
		l.logger.Log("method", "AssignAssetsToWallet", "address", address, "assets", assets, "err", err)
	}()
	return l.next.AssignAssetsToWallet(ctx, address, assets)
}

func (l loggingMiddleware) CreateBalancesForAsset(ctx context.Context, wallets []*repository.Wallet, asset *repository.SimpleAsset) (err error) {
	defer func() {
		l.logger.Log("method", "CreateBalancesForAsset", "wallets", wallets, "asset", asset, "err", err)
	}()
	return l.next.CreateBalancesForAsset(ctx, wallets, asset)
}
