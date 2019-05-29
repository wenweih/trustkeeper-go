package service

import (
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
	CreateGroup(ctx context.Context, uuid, name, desc string) (result bool, err error)
	GetGroups(ctx context.Context, uuid string) (groups []*Group, err error)
}

type basicDashboardService struct {
	repo repository.DashboardRepo
	conf configure.Conf
}

func (b *basicDashboardService) GetGroups(ctx context.Context, uuid string) (groups []*Group, err error) {
	return groups, err
}

func (b *basicDashboardService) CreateGroup(ctx context.Context, uuid, name, desc string) (bool, error) {
	group := &model.Group{CreatorID: uuid, Name: name, Desc: desc}
	err := b.repo.Create(group)
	if err != nil {
		return false, err
	}
	return true, nil
}

// NewBasicDashboardService returns a naive, stateless implementation of DashboardService.
func NewBasicDashboardService(conf configure.Conf) DashboardService {
	db := repository.DB(conf.DBInfo)
	bas := basicDashboardService{
		repo: repository.New(db),
		conf: conf,
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
