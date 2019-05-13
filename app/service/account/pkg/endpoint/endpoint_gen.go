// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "trustkeeper-go/app/service/account/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	SigninEndpoint   endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.AccountService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateEndpoint: MakeCreateEndpoint(s),
		SigninEndpoint:   MakeSigninEndpoint(s),
	}
	for _, m := range mdw["Create"] {
		eps.CreateEndpoint = m(eps.CreateEndpoint)
	}
	for _, m := range mdw["Signin"] {
		eps.SigninEndpoint = m(eps.SigninEndpoint)
	}
	return eps
}
