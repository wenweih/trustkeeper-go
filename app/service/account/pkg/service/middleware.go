package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	gouuid "github.com/satori/go.uuid"
)

// Middleware describes a service middleware.
type Middleware func(AccountService) AccountService

type loggingMiddleware struct {
	logger log.Logger
	next   AccountService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AccountService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AccountService) AccountService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, email string, password string) (u0 gouuid.UUID, e1 error) {
	defer func() {
		l.logger.Log("method", "Create", "email", email, "password", password, "u0", u0, "e1", e1)
	}()
	return l.next.Create(ctx, email, password)
}
