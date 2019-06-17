package service

import (
	"fmt"
	"errors"
	"strings"
	"context"
	"regexp"
	accountService "trustkeeper-go/app/service/account/pkg/service"
	dashboardService "trustkeeper-go/app/service/dashboard/pkg/service"

	log "github.com/go-kit/kit/log"

	accountGrpcClient "trustkeeper-go/app/service/account/client"
	dashboardGrpcClient "trustkeeper-go/app/service/dashboard/client"
	"github.com/caarlos0/env"
)

// WebapiService describes the service.
type WebapiService interface {
	Signup(ctx context.Context, user Credentials) (result bool, err error)
	Signin(ctx context.Context, user Credentials) (token string, err error)
	Signout(ctx context.Context) (result bool, err error)
	GetRoles(ctx context.Context) ([]string, error)
	GetGroups(ctx context.Context, uuid string) (groups []*dashboardService.Group, err error)
}

// Credentials Signup Signin params
type Credentials struct {
	Email    	string	`json:"email"`
	Password 	string	`json:"password"`
	OrgName		string	`json:"orgname"`
}

type basicWebapiService struct {
	accountSrv accountService.AccountService
	dashboardSrv dashboardService.DashboardService
}

func makeError(ctx context.Context, err error) error {
	errStr := err.Error()
	switch  {
	case strings.Contains(errStr, "violates unique constraint"):
		re := regexp.MustCompile(`[(](.*?)[)]`)
		result := re.FindAll([]byte(err.Error()), -1)
		return errors.New("Fields exist: " + string(result[0]))
	default:
		return err
	}
}

func (b *basicWebapiService) auth(ctx context.Context) (uuid string, err error) {
	return b.accountSrv.Auth(ctx)
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

func (b *basicWebapiService) GetGroups(ctx context.Context, uuid string) (groups []*dashboardService.Group, err error) {
	// TODO implement the business logic of Group
	return groups, err
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
		accountSrv: accountClient,
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
