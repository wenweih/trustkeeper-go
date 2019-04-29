package service

import (
	"context"
)

// WebapiService describes the service.
type WebapiService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Signup(ctx context.Context, user Credentials) (result bool, err error)
	Signin(ctx context.Context, user Credentials) (token string, err error)
	Signout(ctx context.Context, token string) (result bool, err error)
}

// Credentials Signup Signin params
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type basicWebapiService struct{}

func (b *basicWebapiService) Signup(ctx context.Context, user Credentials) (result bool, err error) {
	// TODO implement the business logic of Signup
	return result, err
}
func (b *basicWebapiService) Signin(ctx context.Context, user Credentials) (token string, err error) {
	// TODO implement the business logic of Signin
	return token, err
}
func (b *basicWebapiService) Signout(ctx context.Context, token string) (result bool, err error) {
	// TODO implement the business logic of Signout
	return result, err
}

// NewBasicWebapiService returns a naive, stateless implementation of WebapiService.
func NewBasicWebapiService() WebapiService {
	return &basicWebapiService{}
}

// New returns a WebapiService with all of the expected middleware wired in.
func New(middleware []Middleware) WebapiService {
	var svc WebapiService = NewBasicWebapiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
