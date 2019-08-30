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

func (l loggingMiddleware) ConstructTxBTC(ctx context.Context, from string, to string, amount string) (unsignedTxHex string, vinAmount int64, err error) {
	defer func() {
		l.logger.Log("method", "ConstructTxBTC", "from", from, "to", to, "amount", amount, "unsignedTxHex", unsignedTxHex, "err", err)
	}()
	return l.next.ConstructTxBTC(ctx, from, to, amount)
}

func (l loggingMiddleware) SendBTCTx(ctx context.Context, signedTxHex string) (txID string, err error) {
	defer func() {
		l.logger.Log("method", "SendBTCTx", "signedTxHex", signedTxHex, "txID", txID, "err", err)
	}()
	return l.next.SendBTCTx(ctx, signedTxHex)
}

func (l loggingMiddleware) QueryBalance(ctx context.Context, symbol string, address string) (balance string, err error) {
	defer func() {
		l.logger.Log("method", "QueryBalance", "symbol", symbol, "address", address, "balance", balance, "err", err)
	}()
	return l.next.QueryBalance(ctx, symbol, address)
}

func (l loggingMiddleware) WalletValidate(ctx context.Context, chainName string, address string) (err error) {
	defer func() {
		l.logger.Log("method", "WalletValidate", "chainName", chainName, "address", address, "err", err)
	}()
	return l.next.WalletValidate(ctx, chainName, address)
}

func (l loggingMiddleware) ConstructTxETH(ctx context.Context, from string, to string, amount string) (unsignedTxHex string, chainID string, err error) {
	defer func() {
		l.logger.Log("method", "ConstructTxETH", "from", from, "to", to, "amount", amount, "unsignedTxHex", unsignedTxHex, "err", err)
	}()
	return l.next.ConstructTxETH(ctx, from, to, amount)
}

func (l loggingMiddleware) SendETHTx(ctx context.Context, signedTxHex string) (txID string, err error) {
	defer func() {
		l.logger.Log("method", "SendETHTx", "signedTxHex", signedTxHex, "txID", txID, "err", err)
	}()
	return l.next.SendETHTx(ctx, signedTxHex)
}

func (l loggingMiddleware) ConstructTxERC20(ctx context.Context, from string, to string, amount string, contract string) (unsignedTxHex string, chainID string, err error) {
	defer func() {
		l.logger.Log("method", "ConstructTxERC20", "from", from, "to", to, "amount", amount, "contract", contract, "unsignedTxHex", unsignedTxHex, "chainID", chainID, "err", err)
	}()
	return l.next.ConstructTxERC20(ctx, from, to, amount, contract)
}

func (l loggingMiddleware) ConstructTxOmni(ctx context.Context, from string, to string, amount string, symbol string) (unsignedTxHex string, vinAmount int64, err error) {
	defer func() {
		l.logger.Log("method", "ConstructTxOmni", "from", from, "to", to, "amount", amount, "symbol", symbol, "unsignedTxHex", unsignedTxHex, "vinAmount", vinAmount, "err", err)
	}()
	return l.next.ConstructTxOmni(ctx, from, to, amount, symbol)
}
