package service

import (
	"fmt"
	"context"
	"trustkeeper-go/app/service/dashboard/pkg/configure"
	"trustkeeper-go/app/service/dashboard/pkg/repository"
	"trustkeeper-go/app/service/dashboard/pkg/model"
	// walletKeyGrpcClient "trustkeeper-go/app/service/wallet_key/client/grpc"
	walletKeyService "trustkeeper-go/app/service/wallet_key/pkg/service"
	walletManagementService "trustkeeper-go/app/service/wallet_management/pkg/service"
	walletManagementGrpcClient "trustkeeper-go/app/service/wallet_management/client"
	log "github.com/go-kit/kit/log"
)

type Group struct {
	Name string
}

// DashboardService describes the service.
type DashboardService interface {
	CreateGroup(ctx context.Context, uuid, name, desc, namespaceID string) (result bool, err error)
	GetGroups(ctx context.Context, uuid string) (groups []*Group, err error)
}

type basicDashboardService struct {
	biz				repository.IBiz
	KeySrv	walletKeyService.WalletKeyService
	WalletSrv walletManagementService.WalletManagementService

}

func (b *basicDashboardService) GetGroups(ctx context.Context, uuid string) (groups []*Group, err error) {
	return groups, err
}

func (b *basicDashboardService) CreateGroup(ctx context.Context, usrID, name, desc, namespaceID string) (bool, error) {
	group := &model.Group{CreatorID: usrID, Name: name, Desc: desc, NamespaceID: namespaceID}
	err := b.biz.Group(group)
	if err != nil {
		return false, err
	}
	return true, nil
}

// NewBasicDashboardService returns a naive, stateless implementation of DashboardService.
func NewBasicDashboardService(conf configure.Conf, logger log.Logger) (*basicDashboardService, error) {
	db, err := repository.DB(conf.DBInfo)
	if err != nil {
		return nil, err
	}
	wmClient, err := walletManagementGrpcClient.New(conf.ConsulAddress, logger)
	if err != nil {
		return nil, fmt.Errorf("walletManagementGrpcClient: %s", err.Error())
	}


	bas := basicDashboardService{
		biz: repository.New(db),
		// walletSrv: walletKeyServiceClient,
		WalletSrv: wmClient,
	}
	return &bas, nil
}

// New returns a DashboardService with all of the expected middleware wired in.
func New(conf configure.Conf, logger log.Logger, middleware []Middleware) (DashboardService, error ){
	bs, err := NewBasicDashboardService(conf, logger)
	if err != nil {
		return nil, err
	}

	var svc DashboardService = bs
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc, nil
}

func NewJobsService(conf configure.Conf, logger log.Logger) (JobService, error) {
	bs, err := NewBasicDashboardService(conf, logger)
	if err != nil {
		return nil, err
	}
	var svc JobService = bs
	return svc, nil
}
