package service

import (
	"context"
	"strings"
	"trustkeeper-go/app/gateway/webapi/pkg/repository"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(WebapiService) WebapiService

type loggingMiddleware struct {
	logger log.Logger
	next   WebapiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a WebapiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next WebapiService) WebapiService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) Signup(ctx context.Context, user Credentials) (result bool, err error) {
	defer func() {
		l.logger.Log("method", "Signup", "email", user.Email, "result", result, "err", err)
	}()
	return l.next.Signup(ctx, user)
}
func (l loggingMiddleware) Signin(ctx context.Context, user Credentials) (token string, err error) {
	defer func() {
		l.logger.Log("method", "Signin", "user", user.Email, "err", err)
	}()
	return l.next.Signin(ctx, user)
}
func (l loggingMiddleware) Signout(ctx context.Context) (result bool, err error) {
	defer func() {
		l.logger.Log("method", "Signout", "result", result, "err", err)
	}()
	return l.next.Signout(ctx)
}

func (l loggingMiddleware) GetRoles(ctx context.Context) (roles []string, err error) {
	defer func() {
		l.logger.Log("method", "GetRoles", "err", err)
	}()
	return l.next.GetRoles(ctx)
}

func (l loggingMiddleware) GetGroups(ctx context.Context) (groups []*repository.GetGroupsResp, err error) {
	defer func() {
		l.logger.Log("method", "GetGroups", "err", err)
	}()
	return l.next.GetGroups(ctx)
}

func (l loggingMiddleware) CreateGroup(ctx context.Context, name string, desc string) (group *repository.GetGroupsResp, err error) {
	defer func() {
		l.logger.Log("method", "CreateGroup", "name", name, "desc", desc, "group", group, "err", err)
	}()
	return l.next.CreateGroup(ctx, name, desc)
}

func (l loggingMiddleware) UpdateGroup(ctx context.Context, groupid string, name string, desc string) (err error) {
	defer func() {
		l.logger.Log("method", "UpdateGroup", "groupid", groupid, "name", name, "desc", desc, "err", err)
	}()
	return l.next.UpdateGroup(ctx, groupid, name, desc)
}

func (l loggingMiddleware) UserInfo(ctx context.Context) (roles []string, orgName string, err error) {
	defer func() {
		l.logger.Log("method", "UserInfo", "roles", strings.Join(roles, " "), "orgName", orgName, "err", err)
	}()
	return l.next.UserInfo(ctx)
}

func (l loggingMiddleware) GetGroupAssets(ctx context.Context, groupid string) (groupAssets []*repository.GroupAsset, err error) {
	defer func() {
		l.logger.Log("method", "GetGroupAssets", "groupid", groupid, "err", err)
	}()
	return l.next.GetGroupAssets(ctx, groupid)
}

func (l loggingMiddleware) ChangeGroupAssets(ctx context.Context, chainAssets []*repository.GroupAsset, groupid string) (groupAssets []*repository.GroupAsset, err error) {
	defer func() {
		l.logger.Log("method", "ChangeGroupAssets", "chainAssets", chainAssets, "groupid", groupid, "err", err)
	}()
	return l.next.ChangeGroupAssets(ctx, chainAssets, groupid)
}

func (l loggingMiddleware) CreateWallet(ctx context.Context, groupid string, chainname string, bip44change int) (id string, address string, respchainname string, status bool, err error) {
	defer func() {
		l.logger.Log("method", "CreateWallet", "groupid", groupid, "chainname", chainname, "bip44change", bip44change, "id", id, "address", address, "status", status, "err", err)
	}()
	return l.next.CreateWallet(ctx, groupid, chainname, bip44change)
}

func (l loggingMiddleware) GetWallets(ctx context.Context, groupid string, page int, limit, bip44change int) (wallets []*repository.ChainWithWallets, err error) {
	defer func() {
		l.logger.Log("method", "GetWallets", "groupid", groupid, "page", page, "limit", limit, "bip44change", bip44change, "wallets", wallets, "err", err)
	}()
	return l.next.GetWallets(ctx, groupid, page, limit, bip44change)
}

func (l loggingMiddleware) QueryOmniProperty(ctx context.Context, identify string) (asset *repository.SimpleAsset, err error) {
	defer func() {
		l.logger.Log("method", "QueryOmniProperty", "identify", identify, "asset", asset, "err", err)
	}()
	return l.next.QueryOmniProperty(ctx, identify)
}

func (l loggingMiddleware) CreateToken(ctx context.Context, groupid string, chainid string, symbol string, identify string, decimal string, chainName string) (asset *repository.SimpleAsset, err error) {
	defer func() {
		l.logger.Log("method", "CreateToken", "groupid", groupid, "chainid", chainid, "symbol", symbol, "identify", identify, "decimal", decimal, "chainName", chainName, "asset", asset, "err", err)
	}()
	return l.next.CreateToken(ctx, groupid, chainid, symbol, identify, decimal, chainName)
}

func (l loggingMiddleware) EthToken(ctx context.Context, tokenHex string) (token *repository.ERC20Token, err error) {
	defer func() {
		l.logger.Log("method", "EthToken", "tokenHex", tokenHex, "token", token, "err", err)
	}()
	return l.next.EthToken(ctx, tokenHex)
}

func (l loggingMiddleware) SendBTCTx(ctx context.Context, from string, to string, amount string) (txid string, err error) {
	defer func() {
		l.logger.Log("method", "SendBTCTx", "from", from, "to", to, "amount", amount, "txid", txid, "err", err)
	}()
	return l.next.SendBTCTx(ctx, from, to, amount)
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

func (l loggingMiddleware) SendETHTx(ctx context.Context, from string, to string, amount string) (txid string, err error) {
	defer func() {
		l.logger.Log("method", "SendETHTx", "from", from, "to", to, "amount", amount, "txid", txid, "err", err)
	}()
	return l.next.SendETHTx(ctx, from, to, amount)
}
