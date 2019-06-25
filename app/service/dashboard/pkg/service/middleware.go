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

func (l loggingMiddleware) GetGroups(ctx context.Context, namespaceID uint) (groups []*repository.GetGroupsResp, err error) {
	defer func() {
		l.logger.Log("method", "GetGroups", "namespaceID", namespaceID, "err", err)
	}()
	return l.next.GetGroups(ctx, namespaceID)
}

func (l loggingMiddleware) CreateGroup(ctx context.Context, uuid, name, desc string, namespaceID uint) (result bool, err error) {
	defer func() {
		l.logger.Log("method", "CreateGroup", "uuid", uuid, "result", result, "err", err)
	}()
	return l.next.CreateGroup(ctx, uuid, name, desc, namespaceID)
}

func (l loggingMiddleware) Close() error {
	defer func() {
		l.logger.Log("method", "Close", "close resource", "(database, redis etc...)")
	}()
	return l.next.Close()
}
