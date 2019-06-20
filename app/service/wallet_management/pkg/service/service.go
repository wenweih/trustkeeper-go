package service

import (
	"fmt"
	"context"
	"trustkeeper-go/library/database/orm"
	log "github.com/go-kit/kit/log"
	"trustkeeper-go/app/service/wallet_management/pkg/configure"
	"trustkeeper-go/app/service/wallet_management/pkg/repository"
	walletKeyService "trustkeeper-go/app/service/wallet_key/pkg/service"
	walletKeyGrpcClient "trustkeeper-go/app/service/wallet_key/client"
)

// WalletManagementService describes the service.
type WalletManagementService interface {
	// Add your methods here
	CreateChain(ctx context.Context, symbol, bit44ID string, status bool) (err error)
}

type basicWalletManagementService struct{
	biz  repository.IBiz
	KeySrv	walletKeyService.WalletKeyService
}

func (b *basicWalletManagementService) CreateChain(ctx context.Context, symbol string, bit44ID string, status bool) (err error) {
	fmt.Println("sysbol: ", symbol, " bit44id: ", bit44ID, " status: ", status)
	return nil
}

// NewBasicWalletManagementService returns a naive, stateless implementation of WalletManagementService.
func newBasicWalletManagementService(conf configure.Conf, logger log.Logger) (*basicWalletManagementService, error) {
	db, err := orm.DB(conf.DBInfo)
	if err != nil {
		return nil, err
	}

	wkClient, err := walletKeyGrpcClient.New(conf.ConsulAddress, logger)
	if err != nil {
		return nil, fmt.Errorf("walletKeyGrpcClient: %s", err.Error())
	}

	bas := basicWalletManagementService{
		biz: repository.New(db),
		KeySrv: wkClient,
	}
	return &bas, nil
}

// New returns a WalletManagementService with all of the expected middleware wired in.
func New(conf configure.Conf, logger log.Logger, middleware []Middleware) (WalletManagementService, error) {
	bs, err := newBasicWalletManagementService(conf, logger)
	if err != nil {
		return nil, err
	}
	var svc WalletManagementService = bs
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc, nil
}

func NewJobsService(conf configure.Conf, logger log.Logger) (JobService, error) {
	bs, err := newBasicWalletManagementService(conf, logger)
	if err != nil {
		return nil, err
	}
	var svc JobService = bs
	return svc, nil
}
