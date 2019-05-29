package service

import (
	"context"
	"trustkeeper-go/app/service/dashboard/pkg/configure"
	"trustkeeper-go/app/service/dashboard/pkg/repository"
)

type Group struct {
	Name string
}

// DashboardService describes the service.
type DashboardService interface {
	CreateGroup(ctx context.Context, uuid string) (result bool, err error)
	GetGroups(ctx context.Context, uuid string) (groups []*Group, err error)
}

type basicDashboardService struct {
	repo repository.DashboardRepo
	conf configure.Conf
}

func (b *basicDashboardService) GetGroups(ctx context.Context, uuid string) (groups []*Group, err error) {
	return groups, err
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

func (b *basicDashboardService) CreateGroup(ctx context.Context, uuid string) (result bool, err error) {
	// TODO implement the business logic of CreateGroup
	return result, err
}
