package service

import (
	"strings"
	"context"
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
