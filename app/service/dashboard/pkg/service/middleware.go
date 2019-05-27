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
		l.logger.Log("method", "GetGroups", "uuid", uuid, "groups", groups, "err", err)
	}()
	return l.next.GetGroups(ctx, uuid)
}
