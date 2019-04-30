package service

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

// AccountService describes the service.
type AccountService interface {
	Create(ctx context.Context, email, password string) (uuid.UUID, error)
}

type basicAccountService struct{}

func (b *basicAccountService) Create(ctx context.Context, email string, password string) (u0 uuid.UUID, e1 error) {
	// TODO implement the business logic of Create
	id := uuid.NewV4()
	return id, e1
}

// NewBasicAccountService returns a naive, stateless implementation of AccountService.
func NewBasicAccountService() AccountService {
	return &basicAccountService{}
}

// New returns a AccountService with all of the expected middleware wired in.
func New(middleware []Middleware) AccountService {
	var svc AccountService = NewBasicAccountService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
