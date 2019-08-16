package service

import (
	"context"
	repository "trustkeeper-go/app/service/wallet_key/pkg/repository"

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

func (l loggingMiddleware) GenerateMnemonic(ctx context.Context, namespaceid string, bip44ids []int32, bip44accountSize int) (xpubs []*Bip44ThirdXpubsForChain, version string, err error) {
	defer func() {
		l.logger.Log("method", "GenerateMnemonic", "version", version, "err", err)
	}()
	return l.next.GenerateMnemonic(ctx, namespaceid, bip44ids, bip44accountSize)
}

func (l loggingMiddleware) Close() error {
	defer func() {
		l.logger.Log("method", "Close", "close resource", "(database, redis etc...)")
	}()
	return l.next.Close()
}

func (l loggingMiddleware) SignedBitcoincoreTx(ctx context.Context, walletHD repository.WalletHD, txHex string, vinAmount int64) (signedTxHex string, err error) {
	defer func() {
		l.logger.Log("method", "SignedBitcoincoreTx", "walletHD", walletHD, "txHex", txHex, "signedTxHex", signedTxHex, "err", err)
	}()
	return l.next.SignedBitcoincoreTx(ctx, walletHD, txHex, vinAmount)
}
