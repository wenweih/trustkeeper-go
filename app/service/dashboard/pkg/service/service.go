package service

import (
	// "fmt"
	"context"
	"trustkeeper-go/app/service/dashboard/pkg/configure"
	"trustkeeper-go/app/service/dashboard/pkg/repository"
	"trustkeeper-go/app/service/dashboard/pkg/model"
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
	biz repository.IBiz
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
func NewBasicDashboardService(conf configure.Conf) DashboardService {
	db := repository.DB(conf.DBInfo)
	bas := basicDashboardService{
		biz: repository.New(db),
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
	db := repository.DB(conf.DBInfo)
	bas := basicDashboardService{
		biz: repository.New(db),
	}
	var svc JobService = &bas
	return svc
}
