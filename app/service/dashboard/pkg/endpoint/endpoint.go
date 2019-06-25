package endpoint

import (
	"context"
	service "trustkeeper-go/app/service/dashboard/pkg/service"
	"trustkeeper-go/app/service/dashboard/pkg/repository"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetGroupsRequest collects the request parameters for the GetGroups method.
type GetGroupsRequest struct {
	NamespaceID uint `json:"namespaceid"`
}

// GetGroupsResponse collects the response parameters for the GetGroups method.
type GetGroupsResponse struct {
	Groups []*repository.GetGroupsResp `json:"groups"`
	Err    error            `json:"err"`
}

// MakeGetGroupsEndpoint returns an endpoint that invokes GetGroups on the service.
func MakeGetGroupsEndpoint(s service.DashboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupsRequest)
		groups, err := s.GetGroups(ctx, req.NamespaceID)
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
func (e Endpoints) GetGroups(ctx context.Context, namespaceID uint) (groups []*repository.GetGroupsResp, err error) {
	request := GetGroupsRequest{}
	response, err := e.GetGroupsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(GetGroupsResponse).Groups, nil
}

// CreateGroupRequest collects the request parameters for the CreateGroup method.
type CreateGroupRequest struct {
	UUID 			string `json:"uuid"`
	Name 			string `json:"name"`
	Desc			string `json:"desc"`
	NamespaceID	uint `json:"namespaceid"`
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
		result, err := s.CreateGroup(ctx, req.UUID, req.Name, req.Desc, req.NamespaceID)
		if err != nil {
			return nil, err
		}
		return CreateGroupResponse{
			Result: result,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateGroupResponse) Failed() error {
	return r.Err
}

// CreateGroup implements Service. Primarily useful in a client.
func (e Endpoints) CreateGroup(ctx context.Context, uuid, name, desc string, namespaceID uint) (result bool, err error) {
	request := CreateGroupRequest{UUID: uuid, Name: name, Desc: desc, NamespaceID: namespaceID}
	response, err := e.CreateGroupEndpoint(ctx, request)
	if err != nil {
		return false, err
	}
	return response.(CreateGroupResponse).Result, nil
}

// Close implements Service. Primarily useful in a client.
func (e Endpoints) Close() error {
	return nil
}
