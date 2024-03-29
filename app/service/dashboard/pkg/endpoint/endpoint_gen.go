// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "trustkeeper-go/app/service/dashboard/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateGroupEndpoint       endpoint.Endpoint
	GetGroupsEndpoint         endpoint.Endpoint
	UpdateGroupEndpoint       endpoint.Endpoint
	GetGroupAssetsEndpoint    endpoint.Endpoint
	ChangeGroupAssetsEndpoint endpoint.Endpoint
	AddAssetEndpoint          endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.DashboardService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddAssetEndpoint:          MakeAddAssetEndpoint(s),
		ChangeGroupAssetsEndpoint: MakeChangeGroupAssetsEndpoint(s),
		CreateGroupEndpoint:       MakeCreateGroupEndpoint(s),
		GetGroupAssetsEndpoint:    MakeGetGroupAssetsEndpoint(s),
		GetGroupsEndpoint:         MakeGetGroupsEndpoint(s),
		UpdateGroupEndpoint:       MakeUpdateGroupEndpoint(s),
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
	for _, m := range mdw["GetGroupAssets"] {
		eps.GetGroupAssetsEndpoint = m(eps.GetGroupAssetsEndpoint)
	}
	for _, m := range mdw["ChangeGroupAssets"] {
		eps.ChangeGroupAssetsEndpoint = m(eps.ChangeGroupAssetsEndpoint)
	}
	for _, m := range mdw["AddAsset"] {
		eps.AddAssetEndpoint = m(eps.AddAssetEndpoint)
	}
	return eps
}
