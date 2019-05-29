package endpoint

import (
	"context"
	service "trustkeeper-go/app/service/dashboard/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetGroupsRequest collects the request parameters for the GetGroups method.
type GetGroupsRequest struct {
	Uuid string `json:"uuid"`
}

// GetGroupsResponse collects the response parameters for the GetGroups method.
type GetGroupsResponse struct {
	Groups []*service.Group `json:"groups"`
	Err    error            `json:"err"`
}

// MakeGetGroupsEndpoint returns an endpoint that invokes GetGroups on the service.
func MakeGetGroupsEndpoint(s service.DashboardService) endpoint.Endpoint {
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

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetGroups implements Service. Primarily useful in a client.
func (e Endpoints) GetGroups(ctx context.Context, uuid string) (groups []*service.Group, err error) {
	request := GetGroupsRequest{Uuid: uuid}
	response, err := e.GetGroupsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetGroupsResponse).Groups, response.(GetGroupsResponse).Err
}

// CreateGroupRequest collects the request parameters for the CreateGroup method.
type CreateGroupRequest struct {
	Uuid string `json:"uuid"`
}

// CreateGroupResponse collects the response parameters for the CreateGroup method.
type CreateGroupResponse struct {
	Result bool  `json:"result"`
	Err    error `json:"err"`
}

// MakeCreateGroupEndpoint returns an endpoint that invokes CreateGroup on the service.
func MakeCreateGroupEndpoint(s service.DashboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateGroupRequest)
		result, err := s.CreateGroup(ctx, req.Uuid)
		return CreateGroupResponse{
			Err:    err,
			Result: result,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateGroupResponse) Failed() error {
	return r.Err
}

// CreateGroup implements Service. Primarily useful in a client.
func (e Endpoints) CreateGroup(ctx context.Context, uuid string) (result bool, err error) {
	request := CreateGroupRequest{Uuid: uuid}
	response, err := e.CreateGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateGroupResponse).Result, response.(CreateGroupResponse).Err
}
