package service

import (
	"context"
	"trustkeeper-go/app/service/transaction/pkg/repository"
)

// TransactionService describes the service.
type TransactionService interface {
	AssignAssetsToWallet(ctx context.Context, address string, assets []*repository.SimpleAsset) (err error)
}

type basicTransactionService struct{}

func (b *basicTransactionService) AssignAssetsToWallet(ctx context.Context, address string, assets []*repository.SimpleAsset) (err error) {
	// TODO implement the business logic of AssignAssetsToWallet
	return err
}

// NewBasicTransactionService returns a naive, stateless implementation of TransactionService.
func NewBasicTransactionService() TransactionService {
	return &basicTransactionService{}
}

// New returns a TransactionService with all of the expected middleware wired in.
func New(middleware []Middleware) TransactionService {
	var svc TransactionService = NewBasicTransactionService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
