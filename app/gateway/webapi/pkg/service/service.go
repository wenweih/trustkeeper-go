package service

import (
	"errors"
	"strings"
	"context"
	"log"
	"regexp"
	accountService "trustkeeper-go/app/service/account/pkg/service"
	dashboardService "trustkeeper-go/app/service/dashboard/pkg/service"

	stdjwt "github.com/go-kit/kit/auth/jwt"
	sdetcd "github.com/go-kit/kit/sd/etcdv3"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	accountGrpcClient "trustkeeper-go/app/service/account/client/grpc"
	dashboardGrpcClient "trustkeeper-go/app/service/dashboard/client/grpc"
	"github.com/caarlos0/env"
	"trustkeeper-go/library/common"
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
	accountServices		accountService.AccountService
	dashboardServices	dashboardService.DashboardService
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
	return b.accountServices.Auth(ctx)
}

func (b *basicWebapiService) Signup(ctx context.Context, user Credentials) (bool, error) {
	if _, err := b.accountServices.Create(ctx, user.Email, user.Password, user.OrgName); err != nil {
		e := makeError(ctx, err)
		return false, e
	}
	return true, nil
}

func (b *basicWebapiService) Signin(ctx context.Context, user Credentials) (token string, err error) {
	token, err = b.accountServices.Signin(ctx, user.Email, user.Password)
	return
}
func (b *basicWebapiService) Signout(ctx context.Context) (result bool, err error) {
	if err := b.accountServices.Signout(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (b *basicWebapiService) GetRoles(ctx context.Context) (s0 []string, e1 error) {
	roles, err := b.accountServices.Roles(ctx)
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
func NewBasicWebapiService() WebapiService {
	type config struct {
		Etcdsrv	string	`env:"etcdsrv"`
	}
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("fail to parse env: ", err.Error())
	}
	client, err := sdetcd.NewClient(context.Background(), []string{cfg.Etcdsrv}, sdetcd.ClientOptions{})
	if err != nil {
		log.Printf("unable to connect to etcd: %s", err.Error())
		log.Fatalln(err.Error())
		return new(basicWebapiService)
	}
	AccountEntries, err := client.GetEntries(common.AccountSrv)
	if err != nil {
		log.Printf("unable to get prefix entries: %s", err.Error())
		return new(basicWebapiService)
	}
	if len(AccountEntries) == 0 {
		log.Printf("entries not eixst")
		return new(basicWebapiService)
	}
	accountSrvConn, err := grpc.Dial(AccountEntries[0], grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to connect to : %s", err.Error())
	}
	// 把带有 jwt token 的上下文设置到 grpc 请求的上下文中
	// https://github.com/go-kit/kit/blob/master/auth/jwt/README.md ContextToGRPC
	accountServiceclient, err := accountGrpcClient.New(accountSrvConn, []grpctransport.ClientOption{(grpctransport.ClientBefore(stdjwt.ContextToGRPC()))})
	if err != nil {
		log.Println(err.Error())
	}


	DashboardEntries, err := client.GetEntries(common.DashboardSrv)
	if err != nil {
		log.Printf("unable to get prefix entries: %s", err.Error())
		return new(basicWebapiService)
	}
	if len(DashboardEntries) == 0 {
		log.Printf("entries not eixst")
		return new(basicWebapiService)
	}
	dashboardSrvconn, err := grpc.Dial(DashboardEntries[0], grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to connect to : %s", err.Error())
	}
	dashboardServiceClient, err := dashboardGrpcClient.New(dashboardSrvconn, []grpctransport.ClientOption{})
	if err != nil {
		log.Println(err.Error())
	}

	return &basicWebapiService{
		accountServices: accountServiceclient,
		dashboardServices: dashboardServiceClient,
	}
}

// New returns a WebapiService with all of the expected middleware wired in.
func New(middleware []Middleware) WebapiService {
	var svc WebapiService = NewBasicWebapiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
