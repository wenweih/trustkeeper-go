package service

import (
	"context"
)

// WalletKeyService describes the service.
type WalletKeyService interface {
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GenerateMnemonic(ctx context.Context, uuid string) (xpub string, err error)
}

type basicWalletKeyService struct{}

func (b *basicWalletKeyService) GenerateMnemonic(ctx context.Context, uuid string) (xpub string, err error) {
	// TODO implement the business logic of GenerateMnemonic
	// https://matthewdowney.github.io/extract-xpub-ethereum-bitcoin-ledger-nano-s.html
	return xpub, err
}

// NewBasicWalletKeyService returns a naive, stateless implementation of WalletKeyService.
func NewBasicWalletKeyService() WalletKeyService {
	return &basicWalletKeyService{}
}

// New returns a WalletKeyService with all of the expected middleware wired in.
func New(middleware []Middleware) WalletKeyService {
	var svc WalletKeyService = NewBasicWalletKeyService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
