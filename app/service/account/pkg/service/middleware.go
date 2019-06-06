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

func (l loggingMiddleware) Create(ctx context.Context, email, password, orgName string) (uuid string, e1 error) {
	defer func() {
		l.logger.Log("method", "Create", "email", email, "OrgName", orgName, "e1", e1)
	}()
	return l.next.Create(ctx, email, password, orgName)
}

func (l loggingMiddleware) Signin(ctx context.Context, email string, password string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "Signin", "email", email, "e1", e1)
	}()
	return l.next.Signin(ctx, email, password)
}

func (l loggingMiddleware) Signout(ctx context.Context) (e0 error) {
	defer func() {
		l.logger.Log("method", "Signout", "e0", e0)
	}()
	return l.next.Signout(ctx)
}

func (l loggingMiddleware) Roles(ctx context.Context) (s0 []string, e1 error) {
	defer func() {
		l.logger.Log("method", "Roles", "s0", s0, "e1", e1)
	}()
	return l.next.Roles(ctx)
}

func (l loggingMiddleware) Auth(ctx context.Context) (uuid string, err error) {
	defer func() {
		l.logger.Log("method", "Auth", "uuid", uuid, "err", err)
	}()
	return l.next.Auth(ctx)
}
