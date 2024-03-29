package service

import (
	"context"
	"trustkeeper-go/app/service/dashboard/pkg/repository"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(DashboardService) DashboardService

type loggingMiddleware struct {
	logger log.Logger
	next   DashboardService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a DashboardService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next DashboardService) DashboardService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) GetGroups(ctx context.Context, namespaceID string) (groups []*repository.GetGroupsResp, err error) {
	defer func() {
		l.logger.Log("method", "GetGroups", "namespaceID", namespaceID, "err", err)
	}()
	return l.next.GetGroups(ctx, namespaceID)
}

func (l loggingMiddleware) CreateGroup(ctx context.Context, uuid, name, desc string, namespaceID string) (group *repository.GetGroupsResp, err error) {
	defer func() {
		l.logger.Log("method", "CreateGroup", "uuid", uuid, "NamespaceID", namespaceID, "err", err)
	}()
	return l.next.CreateGroup(ctx, uuid, name, desc, namespaceID)
}

func (l loggingMiddleware) Close() error {
	defer func() {
		l.logger.Log("method", "Close", "close resource", "(database, redis etc...)")
	}()
	return l.next.Close()
}

func (l loggingMiddleware) UpdateGroup(ctx context.Context, groupID string, name string, desc string) (err error) {
	defer func() {
		l.logger.Log("method", "UpdateGroup", "groupID", groupID, "name", name, "desc", desc, "err", err)
	}()
	return l.next.UpdateGroup(ctx, groupID, name, desc)
}

func (l loggingMiddleware) GetGroupAssets(ctx context.Context, groupID string) (chainAssets []*repository.ChainAsset, err error) {
	defer func() {
		l.logger.Log("method", "GetGroupAsset", "groupID", groupID, "chainAssets", chainAssets, "err", err)
	}()
	return l.next.GetGroupAssets(ctx, groupID)
}

func (l loggingMiddleware) ChangeGroupAssets(ctx context.Context, chainAssets []*repository.ChainAsset, groupid string) (result []*repository.ChainAsset, err error) {
	defer func() {
		l.logger.Log("method", "ChangeGroupAssets", "groupid", groupid, "chainAssets", chainAssets, "err", err)
	}()
	return l.next.ChangeGroupAssets(ctx, chainAssets, groupid)
}

func (l loggingMiddleware) AddAsset(ctx context.Context, groupid string, chainid string, symbol string, identify string, decimal string) (asset *repository.SimpleAsset, err error) {
	defer func() {
		l.logger.Log("method", "AddAsset", "groupid", groupid, "chainid", chainid, "symbol", symbol, "identify", identify, "decimal", decimal, "asset", asset, "err", err)
	}()
	return l.next.AddAsset(ctx, groupid, chainid, symbol, identify, decimal)
}
