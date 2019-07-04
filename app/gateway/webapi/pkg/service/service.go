package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	accountService "trustkeeper-go/app/service/account/pkg/service"
	dashboardService "trustkeeper-go/app/service/dashboard/pkg/service"

	"trustkeeper-go/app/gateway/webapi/pkg/repository"
	accountGrpcClient "trustkeeper-go/app/service/account/client"
	dashboardGrpcClient "trustkeeper-go/app/service/dashboard/client"

	"github.com/caarlos0/env"
	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/copier"
)

// WebapiService describes the service.
type WebapiService interface {
	Signup(ctx context.Context, user Credentials) (result bool, err error)
	Signin(ctx context.Context, user Credentials) (token string, err error)
	Signout(ctx context.Context) (result bool, err error)
	GetRoles(ctx context.Context) ([]string, error)
	GetGroups(ctx context.Context) (groups []*repository.GetGroupsResp, err error)
	CreateGroup(ctx context.Context, name, desc string) (group *repository.GetGroupsResp, err error)
}

// Credentials Signup Signin params
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgName  string `json:"orgname"`
}

type basicWebapiService struct {
	accountSrv   accountService.AccountService
	dashboardSrv dashboardService.DashboardService
}

func makeError(ctx context.Context, err error) error {
	errStr := err.Error()
	switch {
	case strings.Contains(errStr, "violates unique constraint"):
		re := regexp.MustCompile(`[(](.*?)[)]`)
		result := re.FindAll([]byte(err.Error()), -1)
		return errors.New("Fields exist: " + string(result[0]))
	default:
		return err
	}
}

func (b *basicWebapiService) auth(ctx context.Context) (accountUID string, namespaceID string, roles []string, err error) {
	uid, nid, roles, err := b.accountSrv.Auth(ctx)
	if err != nil {
		return "", "", nil, err
	}
	namespaceID = nid
	accountUID = uid
	return
}

func (b *basicWebapiService) Signup(ctx context.Context, user Credentials) (bool, error) {
	if _, err := b.accountSrv.Create(ctx,
		user.Email,
		user.Password,
		user.OrgName); err != nil {
		e := makeError(ctx, err)
		return false, e
	}
	return true, nil
}

func (b *basicWebapiService) Signin(ctx context.Context, user Credentials) (token string, err error) {
	token, err = b.accountSrv.Signin(ctx, user.Email, user.Password)
	return
}
func (b *basicWebapiService) Signout(ctx context.Context) (result bool, err error) {
	if err := b.accountSrv.Signout(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (b *basicWebapiService) GetRoles(ctx context.Context) (s0 []string, e1 error) {
	roles, err := b.accountSrv.Roles(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (b *basicWebapiService) GetGroups(ctx context.Context) ([]*repository.GetGroupsResp, error) {
	accountUID, namespaceID, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := context.WithValue(ctx, "auth",
		struct{Roles []string;UID string;NID string}{roles, accountUID, namespaceID})
	groups, err := b.dashboardSrv.GetGroups(ctxWithAuthInfo, namespaceID)
	if err != nil {
		return nil, err
	}
	lgroups := []*repository.GetGroupsResp{}
	if err := copier.Copy(&lgroups, &groups); err != nil {
		return nil, err
	}
	return lgroups, nil
}

// NewBasicWebapiService returns a naive, stateless implementation of WebapiService.
func NewBasicWebapiService(logger log.Logger) (WebapiService, error) {
	type config struct {
		ConsulAddr string `end:"consuladdr"`
	}
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("fail to parse env: " + err.Error())
	}

	accountClient, err := accountGrpcClient.New(cfg.ConsulAddr, logger)
	if err != nil {
		return nil, fmt.Errorf("accountGrpcClient: %s", err.Error())
	}

	dashboardClient, err := dashboardGrpcClient.New(cfg.ConsulAddr, logger)
	if err != nil {
		return nil, fmt.Errorf("dashboardGrpcClient: %s", err.Error())
	}

	return &basicWebapiService{
		accountSrv:   accountClient,
		dashboardSrv: dashboardClient,
	}, nil
}

// New returns a WebapiService with all of the expected middleware wired in.
func New(logger log.Logger, middleware []Middleware) (WebapiService, error) {
	service, err := NewBasicWebapiService(logger)
	if err != nil {
		return nil, err
	}
	var svc WebapiService = service
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc, nil
}

func (b *basicWebapiService) CreateGroup(ctx context.Context, name string, desc string) (*repository.GetGroupsResp, error) {
	accountUID, namespaceid, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := context.WithValue(ctx, "auth",
		struct{Roles []string;UID string;NID string}{roles, accountUID, namespaceid})

	resp, err := b.dashboardSrv.CreateGroup(ctxWithAuthInfo, accountUID, name, desc, namespaceid)
	if err != nil {
		return nil, err
	}

	return &repository.GetGroupsResp{Name: resp.Name, Desc: resp.Desc, ID: resp.ID}, nil
}
