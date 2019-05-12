package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(AccountService) AccountService

type loggingMiddleware struct {
	logger log.Logger
	next   AccountService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AccountService) AccountService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, email string, password string) (e1 error) {
	defer func() {
		l.logger.Log("method", "Create", "email", email, "password", password, "e1", e1)
	}()
	return l.next.Create(ctx, email, password)
}

func (l loggingMiddleware) Sign(ctx context.Context, email string, password string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "Sign", "email", email, "password", password, "s0", s0, "e1", e1)
	}()
	return l.next.Sign(ctx, email, password)
}
