package service

import (
	"context"
	"strconv"
	"trustkeeper-go/app/service/dashboard/pkg/configure"
	"trustkeeper-go/app/service/dashboard/pkg/model"
	"trustkeeper-go/app/service/dashboard/pkg/repository"
	"trustkeeper-go/library/database/orm"

	log "github.com/go-kit/kit/log"
)

// DashboardService describes the service.
type DashboardService interface {
	CreateGroup(ctx context.Context, uuid, name, desc string, namespaceID string) (group *repository.GetGroupsResp, err error)
	GetGroups(ctx context.Context, namespaceID string) (groups []*repository.GetGroupsResp, err error)
	UpdateGroup(ctx context.Context, groupID, name, desc string) (err error)
	GetGroupAssets(ctx context.Context, groupID string) (chainAssets []*repository.ChainAsset, err error)
	ChangeGroupAssets(ctx context.Context, chainAssets []*repository.ChainAsset, groupid string) (result []*repository.ChainAsset, err error)
	// CreateWallet(ctx context.Context, groupid string, chainname string, bip44change int) (err error)
	Close() error
}

type basicDashboardService struct {
	biz repository.IBiz
}

func (b *basicDashboardService) Close() error {
	return b.biz.Close()
}

func (b *basicDashboardService) GetGroups(ctx context.Context, namespaceID string) (groups []*repository.GetGroupsResp, err error) {
	groups, err = b.biz.GetGroups(ctx, map[string]interface{}{"namespace_id": namespaceID})
	return
}

func (b *basicDashboardService) CreateGroup(ctx context.Context, usrID, name, desc string, namespaceID string) (g *repository.GetGroupsResp, err error) {
	group := &model.Group{CreatorID: usrID, Name: name, Desc: desc, NamespaceID: namespaceID}
	if err := b.biz.CreateGroup(ctx, group); err != nil {
		return nil, err
	}
	g, err = &repository.GetGroupsResp{Name: group.Name, Desc: group.Desc, ID: strconv.FormatUint(uint64(group.ID), 10)}, nil
	return
}

func (b *basicDashboardService) UpdateGroup(ctx context.Context, groupID string, name string, desc string) (err error) {
	err = b.biz.UpdateGroup(ctx, groupID, name, desc)
	return
}

func (b *basicDashboardService) GetGroupAssets(ctx context.Context, groupID string) (chainAssets []*repository.ChainAsset, err error) {
	chainAssets, err = b.biz.QueryChainAsset(ctx, map[string]interface{}{"group_id": groupID})
	return
}

func (b *basicDashboardService) ChangeGroupAssets(ctx context.Context, chainAssets []*repository.ChainAsset, groupid string) (result []*repository.ChainAsset, err error) {
	result, err = b.biz.ChangeGroupAssets(ctx, chainAssets, groupid)
	return
}

// returns a naive, stateless implementation of DashboardService.
func newBasicDashboardService(conf configure.Conf, logger log.Logger) (*basicDashboardService, error) {
	db, err := orm.DB(conf.DBInfo)
	if err != nil {
		return nil, err
	}

	bas := basicDashboardService{
		biz: repository.New(db),
	}
	return &bas, nil
}

// New returns a DashboardService with all of the expected middleware wired in.
func New(conf configure.Conf, logger log.Logger, middleware []Middleware) (DashboardService, error) {
	bs, err := newBasicDashboardService(conf, logger)
	if err != nil {
		return nil, err
	}

	var svc DashboardService = bs
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc, nil
}

// NewJobsService job service
func NewJobsService(conf configure.Conf, logger log.Logger) (JobService, error) {
	bs, err := newBasicDashboardService(conf, logger)
	if err != nil {
		return nil, err
	}
	var svc JobService = bs
	return svc, nil
}
