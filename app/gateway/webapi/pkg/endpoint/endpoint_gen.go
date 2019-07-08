// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "trustkeeper-go/app/gateway/webapi/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	SignupEndpoint      endpoint.Endpoint
	SigninEndpoint      endpoint.Endpoint
	SignoutEndpoint     endpoint.Endpoint
	GetRolesEndpoint    endpoint.Endpoint
	UserInfoEndpoint    endpoint.Endpoint
	GetGroupsEndpoint   endpoint.Endpoint
	CreateGroupEndpoint endpoint.Endpoint
	UpdateGroupEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.WebapiService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateGroupEndpoint: MakeCreateGroupEndpoint(s),
		GetGroupsEndpoint:   MakeGetGroupsEndpoint(s),
		GetRolesEndpoint:    MakeGetRolesEndpoint(s),
		SigninEndpoint:      MakeSigninEndpoint(s),
		SignoutEndpoint:     MakeSignoutEndpoint(s),
		SignupEndpoint:      MakeSignupEndpoint(s),
		UpdateGroupEndpoint: MakeUpdateGroupEndpoint(s),
		UserInfoEndpoint:    MakeUserInfoEndpoint(s),
	}
	for _, m := range mdw["Signup"] {
		eps.SignupEndpoint = m(eps.SignupEndpoint)
	}
	for _, m := range mdw["Signin"] {
		eps.SigninEndpoint = m(eps.SigninEndpoint)
	}
	for _, m := range mdw["Signout"] {
		eps.SignoutEndpoint = m(eps.SignoutEndpoint)
	}
	for _, m := range mdw["GetRoles"] {
		eps.GetRolesEndpoint = m(eps.GetRolesEndpoint)
	}
	for _, m := range mdw["UserInfo"] {
		eps.UserInfoEndpoint = m(eps.UserInfoEndpoint)
	}
	for _, m := range mdw["GetGroups"] {
		eps.GetGroupsEndpoint = m(eps.GetGroupsEndpoint)
	}
	for _, m := range mdw["CreateGroup"] {
		eps.CreateGroupEndpoint = m(eps.CreateGroupEndpoint)
	}
	for _, m := range mdw["UpdateGroup"] {
		eps.UpdateGroupEndpoint = m(eps.UpdateGroupEndpoint)
	}
	return eps
}
