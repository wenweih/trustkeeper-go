package service

import (
	"context"
	"trustkeeper-go/app/service/wallet_management/pkg/repository"

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

func (l loggingMiddleware) Close() error {
	defer func() {
		l.logger.Log("method", "Close", "close resource", "(database, redis etc...)")
	}()
	return l.next.Close()
}

func (l loggingMiddleware) AssignedXpubToGroup(ctx context.Context, groupid string) (err error) {
	defer func() {
		l.logger.Log("method", "AssignedXpubToGroup", "groupid", groupid, "err", err)
	}()
	return l.next.AssignedXpubToGroup(ctx, groupid)
}

func (l loggingMiddleware) GetChains(ctx context.Context) (chains []*repository.SimpleChain, err error) {
	defer func() {
		l.logger.Log("method", "GetChains", "chains", chains, "err", err)
	}()
	return l.next.GetChains(ctx)
}

func (l loggingMiddleware) CreateWallet(ctx context.Context, groupid string, chainname string, bip44change int) (wallet *repository.Wallet, err error) {
	defer func() {
		l.logger.Log("method", "CreateWallet", "groupid", groupid, "chainname", chainname, "bip44change", bip44change, "err", err)
	}()
	return l.next.CreateWallet(ctx, groupid, chainname, bip44change)
}

func (l loggingMiddleware) GetWallets(ctx context.Context, groupid string, page, limit, bip44change int32) (wallets []*repository.ChainWithWallets, err error) {
	defer func() {
		l.logger.Log("method", "GetWallets", "groupid", groupid, "wallets", wallets, "err", err)
	}()
	return l.next.GetWallets(ctx, groupid, page, limit, bip44change)
}

func (l loggingMiddleware) QueryWalletsForGroupByChainName(ctx context.Context, groupid string, chainName string) (wallets []*repository.Wallet, err error) {
	defer func() {
		l.logger.Log("method", "QueryWalletsForGroupByChainName", "groupid", groupid, "chainName", chainName, "wallets", wallets, "err", err)
	}()
	return l.next.QueryWalletsForGroupByChainName(ctx, groupid, chainName)
}

func (l loggingMiddleware) QueryWalletHD(ctx context.Context, address string) (hd *repository.WalletHD, err error) {
	defer func() {
		l.logger.Log("method", "QueryWalletHD", "address", address, "hd", hd, "err", err)
	}()
	return l.next.QueryWalletHD(ctx, address)
}
