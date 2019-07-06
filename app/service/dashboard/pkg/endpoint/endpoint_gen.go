// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "trustkeeper-go/app/service/dashboard/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateGroupEndpoint endpoint.Endpoint
	GetGroupsEndpoint   endpoint.Endpoint
	UpdateGroupEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.DashboardService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateGroupEndpoint: MakeCreateGroupEndpoint(s),
		GetGroupsEndpoint:   MakeGetGroupsEndpoint(s),
		UpdateGroupEndpoint: MakeUpdateGroupEndpoint(s),
	}
	for _, m := range mdw["CreateGroup"] {
		eps.CreateGroupEndpoint = m(eps.CreateGroupEndpoint)
	}
	for _, m := range mdw["GetGroups"] {
		eps.GetGroupsEndpoint = m(eps.GetGroupsEndpoint)
	}
	for _, m := range mdw["UpdateGroup"] {
		eps.UpdateGroupEndpoint = m(eps.UpdateGroupEndpoint)
	}
	return eps
}
