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

func (l loggingMiddleware) QueryToken(ctx context.Context, identify string) (symbol string, err error) {
	defer func() {
		l.logger.Log("method", "QueryToken", "identify", identify, "symbol", symbol, "err", err)
	}()
	return l.next.QueryToken(ctx, identify)
}
