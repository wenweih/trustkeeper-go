package service

import (
	// "fmt"
	"log"
	"context"
	"trustkeeper-go/app/service/dashboard/pkg/configure"
	"trustkeeper-go/app/service/dashboard/pkg/repository"
	"trustkeeper-go/app/service/dashboard/pkg/model"
	walletKeyGrpcClient "trustkeeper-go/app/service/wallet_key/client/grpc"
	walletKeyService "trustkeeper-go/app/service/wallet_key/pkg/service"
	sdetcd "github.com/go-kit/kit/sd/etcdv3"
	"google.golang.org/grpc"
	"trustkeeper-go/library/common"
	grpctransport "github.com/go-kit/kit/transport/grpc"
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
	walletSrv	walletKeyService.WalletKeyService
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
func NewBasicDashboardService(conf configure.Conf) *basicDashboardService {
	client, err := sdetcd.NewClient(context.Background(), []string{conf.EtcdServer}, sdetcd.ClientOptions{})
	if err != nil {
		log.Printf("unable to connect to etcd: %s", err.Error())
		log.Fatalln(err.Error())
		return new(basicDashboardService)
	}
	walletKeyEntries, err := client.GetEntries(common.WalletKeySrv)
	if err != nil {
		log.Printf("unable to get prefix entries: %s", err.Error())
		return new(basicDashboardService)
	}
	if len(walletKeyEntries) == 0 {
		log.Printf("entries not eixst")
		return new(basicDashboardService)
	}
	walletKeySrvconn, err := grpc.Dial(walletKeyEntries[0], grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to connect to : %s", err.Error())
	}
	walletKeyServiceClient, err := walletKeyGrpcClient.New(walletKeySrvconn, []grpctransport.ClientOption{})
	if err != nil {
		log.Println(err.Error())
	}

	db := repository.DB(conf.DBInfo)

	bas := basicDashboardService{
		biz: repository.New(db),
		walletSrv: walletKeyServiceClient,
	}
	return &bas
}

// New returns a DashboardService with all of the expected middleware wired in.
func New(conf configure.Conf, middleware []Middleware) DashboardService {
	var svc DashboardService = NewBasicDashboardService(conf)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func NewJobsService(conf configure.Conf) JobService {
	var svc JobService = NewBasicDashboardService(conf)
	return svc
}
