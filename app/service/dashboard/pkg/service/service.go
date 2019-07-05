package service

import (
	"strconv"
	"context"
	"trustkeeper-go/app/service/dashboard/pkg/configure"
	"trustkeeper-go/app/service/dashboard/pkg/repository"
	"trustkeeper-go/app/service/dashboard/pkg/model"
	log "github.com/go-kit/kit/log"
	"trustkeeper-go/library/database/orm"
)

// DashboardService describes the service.
type DashboardService interface {
	CreateGroup(ctx context.Context, uuid, name, desc string, namespaceID string) (group *repository.GetGroupsResp, err error)
	GetGroups(ctx context.Context, namespaceID string) (groups []*repository.GetGroupsResp, err error)
	Close() error
}

type basicDashboardService struct {
	biz				repository.IBiz
}

func (b *basicDashboardService) Close() error{
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

// NewBasicDashboardService returns a naive, stateless implementation of DashboardService.
func NewBasicDashboardService(conf configure.Conf, logger log.Logger) (*basicDashboardService, error) {
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
