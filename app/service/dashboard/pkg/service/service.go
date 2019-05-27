package service

import (
	"context"
)

type Group struct {
  Name string
}

// DashboardService describes the service.
type DashboardService interface {
	GetGroups(ctx context.Context, uuid string) (groups []*Group, err error)
}

type basicDashboardService struct{}

func (b *basicDashboardService) GetGroups(ctx context.Context, uuid string) (groups []*Group, err error) {
	// TODO implement the business logic of GetGroups
	return groups, err
}

// NewBasicDashboardService returns a naive, stateless implementation of DashboardService.
func NewBasicDashboardService() DashboardService {
	return &basicDashboardService{}
}

// New returns a DashboardService with all of the expected middleware wired in.
func New(middleware []Middleware) DashboardService {
	var svc DashboardService = NewBasicDashboardService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
