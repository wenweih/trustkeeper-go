package service

import (
	"context"

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

func (l loggingMiddleware) GetGroups(ctx context.Context, uuid string) (groups []*Group, err error) {
	defer func() {
		l.logger.Log("method", "GetGroups", "uuid", uuid, "err", err)
	}()
	return l.next.GetGroups(ctx, uuid)
}

func (l loggingMiddleware) CreateGroup(ctx context.Context, uuid, name, desc, namespaceID string) (result bool, err error) {
	defer func() {
		l.logger.Log("method", "CreateGroup", "uuid", uuid, "result", result, "err", err)
	}()
	return l.next.CreateGroup(ctx, uuid, name, desc, namespaceID)
}
