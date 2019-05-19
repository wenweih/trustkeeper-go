package service

import (
	"context"
	"log"
	"trustkeeper-go/app/service/account/pkg/grpc/pb"

	sdetcd "github.com/go-kit/kit/sd/etcdv3"
	"google.golang.org/grpc"
)

// WebapiService describes the service.
type WebapiService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Signup(ctx context.Context, user Credentials) (result bool, err error)
	Signin(ctx context.Context, user Credentials) (token string, err error)
	Signout(ctx context.Context, token string) (result bool, err error)
	GetRoles(ctx context.Context, token string) ([]string, error)
}

// Credentials Signup Signin params
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type basicWebapiService struct {
	accountServiceClient pb.AccountClient
}

func (b *basicWebapiService) Signup(ctx context.Context, user Credentials) (result bool, err error) {
	resp, err := b.accountServiceClient.Create(ctx, &pb.CreateRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return false, err
	}
	return resp.Result, nil
}
func (b *basicWebapiService) Signin(ctx context.Context, user Credentials) (token string, err error) {
	resp, err := b.accountServiceClient.Signin(ctx, &pb.SigninRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, err
}
func (b *basicWebapiService) Signout(ctx context.Context, token string) (result bool, err error) {
	resp, err := b.accountServiceClient.Signout(ctx, &pb.SignoutRequest{Token: token})
	if err != nil {
		return false, err
	}
	return resp.Result, nil
}


func (b *basicWebapiService) GetRoles(ctx context.Context, token string) (s0 []string, e1 error) {
	resp, err := b.accountServiceClient.Roles(ctx, &pb.RolesRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return resp.Roles, nil
}

// NewBasicWebapiService returns a naive, stateless implementation of WebapiService.
func NewBasicWebapiService() WebapiService {
	var etcdServer = "http://localhost:2379"
	client, err := sdetcd.NewClient(context.Background(), []string{etcdServer}, sdetcd.ClientOptions{})
	if err != nil {
		log.Printf("unable to connect to etcd: %s", err.Error())
		return new(basicWebapiService)
	}
	entries, err := client.GetEntries("/services/account/")
	if err != nil {
		log.Printf("unable to get prefix entries: %s", err.Error())
		return new(basicWebapiService)
	}

	if len(entries) == 0 {
		log.Printf("entries not eixst")
		return new(basicWebapiService)
	}
	conn, err := grpc.Dial(entries[0], grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to connect to : %s", err.Error())
	}

	return &basicWebapiService{
		accountServiceClient: pb.NewAccountClient(conn),
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
