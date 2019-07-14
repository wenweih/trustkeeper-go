package endpoint

import (
	"context"
	service "trustkeeper-go/app/gateway/webapi/pkg/service"

	"trustkeeper-go/app/gateway/webapi/pkg/repository"

	endpoint "github.com/go-kit/kit/endpoint"
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
type SignoutRequest struct{}

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
type GetGroupsRequest struct{}

// GetGroupsResponse collects the response parameters for the Group method.
type GetGroupsResponse struct {
	Groups []*repository.GetGroupsResp `json:"groups"`
	Err    error                       `json:"err"`
}

// MakeGetGroupsEndpoint returns an endpoint that invokes Group on the service.
func MakeGetGroupsEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		groups, err := s.GetGroups(ctx)
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
func (e Endpoints) GetGroups(ctx context.Context) (groups []*repository.GetGroupsResp, err error) {
	request := GetGroupsRequest{}
	response, err := e.GetGroupsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(GetGroupsResponse).Groups, nil
}

// CreateGroupRequest collects the request parameters for the CreateGroup method.
type CreateGroupRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// CreateGroupResponse collects the response parameters for the CreateGroup method.
type CreateGroupResponse struct {
	Group *repository.GetGroupsResp `json:"group"`
	Err   error                     `json:"err"`
}

// MakeCreateGroupEndpoint returns an endpoint that invokes CreateGroup on the service.
func MakeCreateGroupEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateGroupRequest)
		group, err := s.CreateGroup(ctx, req.Name, req.Desc)
		if err != nil {
			return CreateGroupResponse{
				Err: err,
			}, err
		}
		return CreateGroupResponse{
			Err:   nil,
			Group: group,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateGroupResponse) Failed() error {
	return r.Err
}

// CreateGroup implements Service. Primarily useful in a client.
func (e Endpoints) CreateGroup(ctx context.Context, name string, desc string) (group *repository.GetGroupsResp, err error) {
	request := CreateGroupRequest{
		Desc: desc,
		Name: name,
	}
	response, err := e.CreateGroupEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(CreateGroupResponse).Group, nil
}

// UpdateGroupRequest collects the request parameters for the UpdateGroup method.
type UpdateGroupRequest struct {
	Groupid string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
}

// UpdateGroupResponse collects the response parameters for the UpdateGroup method.
type UpdateGroupResponse struct {
	Err error `json:"err"`
}

// MakeUpdateGroupEndpoint returns an endpoint that invokes UpdateGroup on the service.
func MakeUpdateGroupEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateGroupRequest)
		err := s.UpdateGroup(ctx, req.Groupid, req.Name, req.Desc)
		if err != nil {
			return UpdateGroupResponse{}, err
		}
		return UpdateGroupResponse{}, nil
	}
}

// Failed implements Failer.
func (r UpdateGroupResponse) Failed() error {
	return r.Err
}

// UpdateGroup implements Service. Primarily useful in a client.
func (e Endpoints) UpdateGroup(ctx context.Context, groupid string, name string, desc string) (err error) {
	request := UpdateGroupRequest{
		Desc:    desc,
		Groupid: groupid,
		Name:    name,
	}
	_, err = e.UpdateGroupEndpoint(ctx, request)
	return
}

// UserInfoRequest collects the request parameters for the UserInfo method.
type UserInfoRequest struct{}

// UserInfoResponse collects the response parameters for the UserInfo method.
type UserInfoResponse struct {
	Roles   []string `json:"roles"`
	OrgName string   `json:"org_name"`
	Err     error    `json:"err"`
}

// MakeUserInfoEndpoint returns an endpoint that invokes UserInfo on the service.
func MakeUserInfoEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		roles, orgName, err := s.UserInfo(ctx)
		return UserInfoResponse{
			Err:     err,
			OrgName: orgName,
			Roles:   roles,
		}, nil
	}
}

// Failed implements Failer.
func (r UserInfoResponse) Failed() error {
	return r.Err
}

// UserInfo implements Service. Primarily useful in a client.
func (e Endpoints) UserInfo(ctx context.Context) (roles []string, orgName string, err error) {
	request := UserInfoRequest{}
	response, err := e.UserInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UserInfoResponse).Roles, response.(UserInfoResponse).OrgName, response.(UserInfoResponse).Err
}

// GetGroupAssetsRequest collects the request parameters for the GetGroupAssets method.
type GetGroupAssetsRequest struct {
	Groupid string `json:"groupid"`
}

// GetGroupAssetsResponse collects the response parameters for the GetGroupAssets method.
type GetGroupAssetsResponse struct {
	GroupAssets []*repository.GroupAsset `json:"group_assets"`
	Err         error                    `json:"err"`
}

// MakeGetGroupAssetsEndpoint returns an endpoint that invokes GetGroupAssets on the service.
func MakeGetGroupAssetsEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupAssetsRequest)
		groupAssets, err := s.GetGroupAssets(ctx, req.Groupid)
		if err != nil {
			return GetGroupAssetsResponse{
				Err: err,
			}, err
		}
		return GetGroupAssetsResponse{
			GroupAssets: groupAssets,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupAssetsResponse) Failed() error {
	return r.Err
}

// GetGroupAssets implements Service. Primarily useful in a client.
func (e Endpoints) GetGroupAssets(ctx context.Context, groupid string) (groupAssets []*repository.GroupAsset, err error) {
	request := GetGroupAssetsRequest{Groupid: groupid}
	response, err := e.GetGroupAssetsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(GetGroupAssetsResponse).GroupAssets, nil
}

// ChangeGroupAssetsRequest collects the request parameters for the ChangeGroupAssets method.
type ChangeGroupAssetsRequest struct {
	ChainAssets []*repository.GroupAsset `json:"chain_assets"`
	Groupid     string                   `json:"groupid"`
}

// ChangeGroupAssetsResponse collects the response parameters for the ChangeGroupAssets method.
type ChangeGroupAssetsResponse struct {
	GroupAssets []*repository.GroupAsset `json:"groupassets"`
	Err error `json:"err"`
}

// MakeChangeGroupAssetsEndpoint returns an endpoint that invokes ChangeGroupAssets on the service.
func MakeChangeGroupAssetsEndpoint(s service.WebapiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeGroupAssetsRequest)
		result, err := s.ChangeGroupAssets(ctx, req.ChainAssets, req.Groupid)
		if err != nil {
			return ChangeGroupAssetsResponse{Err: err}, err
		}
		return ChangeGroupAssetsResponse{GroupAssets: result}, nil
	}
}

// Failed implements Failer.
func (r ChangeGroupAssetsResponse) Failed() error {
	return r.Err
}

// ChangeGroupAssets implements Service. Primarily useful in a client.
func (e Endpoints) ChangeGroupAssets(ctx context.Context, chainAssets []*repository.GroupAsset, groupid string) (result []*repository.GroupAsset, err error) {
	request := ChangeGroupAssetsRequest{
		ChainAssets: chainAssets,
		Groupid:     groupid,
	}
	resp, err := e.ChangeGroupAssetsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(ChangeGroupAssetsResponse).GroupAssets, nil
}
