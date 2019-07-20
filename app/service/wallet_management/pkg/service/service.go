package service

import (
	"context"
	"fmt"
	walletKeyGrpcClient "trustkeeper-go/app/service/wallet_key/client"
	walletKeyService "trustkeeper-go/app/service/wallet_key/pkg/service"
	"trustkeeper-go/app/service/wallet_management/pkg/configure"
	"trustkeeper-go/app/service/wallet_management/pkg/repository"
	"trustkeeper-go/library/database/orm"
	"trustkeeper-go/library/database/redis"

	log "github.com/go-kit/kit/log"
)

// WalletManagementService describes the service.
type WalletManagementService interface {
	GetChains(ctx context.Context) (chains []*repository.SimpleChain, err error)
	CreateChain(ctx context.Context, symbol, bit44ID string, status bool) (err error)
	AssignedXpubToGroup(ctx context.Context, groupid string) (err error)
	CreateWallet(ctx context.Context, groupid, chainname string, bip44change int) (wallet *repository.Wallet, err error)
	// GetWallets(ctx context.Context, groupid string)
	Close() error
}

type basicWalletManagementService struct {
	biz    repository.IBiz
	KeySrv walletKeyService.WalletKeyService
}

func (b *basicWalletManagementService) Close() error {
	return b.biz.Close()
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

	redisPool := redis.NewPool(conf.Redis)
	wkClient, err := walletKeyGrpcClient.New(conf.ConsulAddress, logger)
	if err != nil {
		return nil, fmt.Errorf("walletKeyGrpcClient: %s", err.Error())
	}

	bas := basicWalletManagementService{
		biz:    repository.New(redisPool, db),
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

func (b *basicWalletManagementService) AssignedXpubToGroup(ctx context.Context, groupid string) (err error) {
	err = b.biz.UpdateXpubState(ctx, repository.Idle, repository.Assigned, groupid)
	return
}

// NewBasicWalletManagementService returns a naive, stateless implementation of WalletManagementService.
func NewBasicWalletManagementService() WalletManagementService {
	return &basicWalletManagementService{}
}

func (b *basicWalletManagementService) GetChains(ctx context.Context) (chains []*repository.SimpleChain, err error) {
	chains, err = b.biz.GetChains(ctx, map[string]interface{}{})
	return
}

func (b *basicWalletManagementService) CreateWallet(ctx context.Context, groupid string, chainname string, bip44change int) (wallet *repository.Wallet, err error) {
	wallet, err = b.biz.CreateWallet(ctx, groupid, chainname, bip44change)
	return
}
