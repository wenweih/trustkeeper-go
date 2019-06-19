package service

import (
	"fmt"
	"context"
	"trustkeeper-go/library/database/orm"
	"trustkeeper-go/app/service/wallet_management/pkg/configure"
	"trustkeeper-go/app/service/wallet_management/pkg/repository"
)

// WalletManagementService describes the service.
type WalletManagementService interface {
	// Add your methods here
	CreateChain(ctx context.Context, symbol, bit44ID string, status bool) (err error)
}

type basicWalletManagementService struct{
	biz  repository.IBiz
}


func (b *basicWalletManagementService) CreateChain(ctx context.Context, symbol string, bit44ID string, status bool) (err error) {
	fmt.Println("sysbol: ", symbol, " bit44id: ", bit44ID, " status: ", status)
	return nil
}

// NewBasicWalletManagementService returns a naive, stateless implementation of WalletManagementService.
func newBasicWalletManagementService(conf configure.Conf) (*basicWalletManagementService, error) {
	db, err := orm.DB(conf.DBInfo)
	if err != nil {
		return nil, err
	}

	bas := basicWalletManagementService{
		biz: repository.New(db),
	}
	return &bas, nil
}

// New returns a WalletManagementService with all of the expected middleware wired in.
func New(conf configure.Conf, middleware []Middleware) (WalletManagementService, error) {
	bs, err := newBasicWalletManagementService(conf)
	if err != nil {
		return nil, err
	}
	var svc WalletManagementService = bs
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc, nil
}
