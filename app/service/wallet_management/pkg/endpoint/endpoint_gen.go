// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "trustkeeper-go/app/service/wallet_management/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetChainsEndpoint           endpoint.Endpoint
	CreateChainEndpoint         endpoint.Endpoint
	AssignedXpubToGroupEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.WalletManagementService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AssignedXpubToGroupEndpoint: MakeAssignedXpubToGroupEndpoint(s),
		CreateChainEndpoint:         MakeCreateChainEndpoint(s),
		GetChainsEndpoint:           MakeGetChainsEndpoint(s),
	}
	for _, m := range mdw["GetChains"] {
		eps.GetChainsEndpoint = m(eps.GetChainsEndpoint)
	}
	for _, m := range mdw["CreateChain"] {
		eps.CreateChainEndpoint = m(eps.CreateChainEndpoint)
	}
	for _, m := range mdw["AssignedXpubToGroup"] {
		eps.AssignedXpubToGroupEndpoint = m(eps.AssignedXpubToGroupEndpoint)
	}
	return eps
}
