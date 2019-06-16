package service

import (
	"fmt"
	"context"
)

// WalletManagementService describes the service.
type WalletManagementService interface {
	// Add your methods here
	CreateChain(ctx context.Context, symbol, bit44ID string, status bool) (err error)
}

type basicWalletManagementService struct{}


func (b *basicWalletManagementService) CreateChain(ctx context.Context, symbol string, bit44ID string, status bool) (err error) {
	fmt.Println("sysbol: ", symbol, " bit44id: ", bit44ID, " status: ", status)
	return nil
}

// NewBasicWalletManagementService returns a naive, stateless implementation of WalletManagementService.
func NewBasicWalletManagementService() WalletManagementService {
	return &basicWalletManagementService{}
}

// New returns a WalletManagementService with all of the expected middleware wired in.
func New(middleware []Middleware) WalletManagementService {
	var svc WalletManagementService = NewBasicWalletManagementService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
