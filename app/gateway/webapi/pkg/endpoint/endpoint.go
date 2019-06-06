package endpoint

import (
	"context"
	service "trustkeeper-go/app/gateway/webapi/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
	// groupModel "trustkeeper-go/app/service/dashboard/pkg/model"
	groupService "trustkeeper-go/app/service/dashboard/pkg/service"
)

// SignupRequest collects the request parameters for the Signup method.
type SignupRequest struct {
	User service.Credentials `json:"user"`
}

// SignupResponse collects the response parameters for the Signup method.
type SignupResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"err"`
}

// MakeSignupEndpoint returns an endpoint that invokes Signup on the service.
func MakeSignupEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignupRequest)
		result, err := s.Signup(ctx, req.User)
		return SignupResponse{
			Err:    err,
			Result: result,
		}, err
	}
}

// Failed implements Failer.
func (r SignupResponse) Failed() error {
	return r.Err
}

// SigninRequest collects the request parameters for the Signin method.
type SigninRequest struct {
	User service.Credentials `json:"user"`
}

// SigninResponse collects the response parameters for the Signin method.
type SigninResponse struct {
	Token string `json:"token"`
	Err   error  `json:"err"`
}

// MakeSigninEndpoint returns an endpoint that invokes Signin on the service.
func MakeSigninEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SigninRequest)
		token, err := s.Signin(ctx, req.User)
		return SigninResponse{
			Err:   err,
			Token: token,
		}, nil
	}
}

// Failed implements Failer.
func (r SigninResponse) Failed() error {
	return r.Err
}

// SignoutRequest collects the request parameters for the Signout method.
type SignoutRequest struct {
	// Token string `json:"token"`
}

// SignoutResponse collects the response parameters for the Signout method.
type SignoutResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"err"`
}

// MakeSignoutEndpoint returns an endpoint that invokes Signout on the service.
func MakeSignoutEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.Signout(ctx)
		return SignoutResponse{
			Err:    err,
			Result: result,
		}, err
	}
}

// Failed implements Failer.
func (r SignoutResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Signup implements Service. Primarily useful in a client.
func (e Endpoints) Signup(ctx context.Context, user service.Credentials) (result bool, err error) {
	request := SignupRequest{User: user}
	response, err := e.SignupEndpoint(ctx, request)
	if err != nil {
		return false, err
	}
	return response.(SignupResponse).Result, response.(SignupResponse).Err
}

// Signin implements Service. Primarily useful in a client.
func (e Endpoints) Signin(ctx context.Context, user service.Credentials) (token string, err error) {
	request := SigninRequest{User: user}
	response, err := e.SigninEndpoint(ctx, request)
	if err != nil {
		return "", err
	}
	return response.(SigninResponse).Token, response.(SigninResponse).Err
}

// Signout implements Service. Primarily useful in a client.
func (e Endpoints) Signout(ctx context.Context) (result bool, err error) {
	request := SignoutRequest{}
	response, err := e.SignoutEndpoint(ctx, request)
	if err != nil {
		return false, err
	}
	return response.(SignoutResponse).Result, response.(SignoutResponse).Err
}

// GetRolesRequest collects the request parameters for the GetRoles method.
type GetRolesRequest struct{}

// GetRolesResponse collects the response parameters for the GetRoles method.
type GetRolesResponse struct {
	S0 []string `json:"s0"`
	E1 error    `json:"e1"`
}

// MakeGetRolesEndpoint returns an endpoint that invokes GetRoles on the service.
func MakeGetRolesEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		s0, e1 := s.GetRoles(ctx)
		return GetRolesResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetRolesResponse) Failed() error {
	return r.E1
}

// GetRoles implements Service. Primarily useful in a client.
func (e Endpoints) GetRoles(ctx context.Context, token string) (s0 []string, e1 error) {
	request := GetRolesRequest{}
	response, err := e.GetRolesEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(GetRolesResponse).S0, response.(GetRolesResponse).E1
}

// GroupRequest collects the request parameters for the Group method.
type GetGroupsRequest struct {
	Uuid string `json:"uuid"`
}

// GetGroupsResponse collects the response parameters for the Group method.
type GetGroupsResponse struct {
	Groups []*groupService.Group `json:"groups"`
	Err    error               `json:"err"`
}

// MakeGetGroupsEndpoint returns an endpoint that invokes Group on the service.
func MakeGetGroupsEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupsRequest)
		groups, err := s.GetGroups(ctx, req.Uuid)
		return GetGroupsResponse{
			Err:    err,
			Groups: groups,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupsResponse) Failed() error {
	return r.Err
}

// GetGroups implements Service. Primarily useful in a client.
func (e Endpoints) GetGroups(ctx context.Context, uuid string) (groups []*groupService.Group, err error) {
	request := GetGroupsRequest{Uuid: uuid}
	response, err := e.GetGroupsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetGroupsResponse).Groups, response.(GetGroupsResponse).Err
}
