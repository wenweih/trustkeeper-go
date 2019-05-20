package service

import (
	"context"

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
		l.logger.Log("method", "Signup", "user", user, "result", result, "err", err)
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

func (l loggingMiddleware) GetRoles(ctx context.Context, token string) (s0 []string, e1 error) {
	defer func() {
		l.logger.Log("method", "GetRoles", "s0", s0, "e1", e1)
	}()
	return l.next.GetRoles(ctx, token)
}
