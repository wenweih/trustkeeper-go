package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "trustkeeper-go/app/service/dashboard/pkg/service"
)

// GetGroupsRequest collects the request parameters for the GetGroups method.
type GetGroupsRequest struct {
	Uuid string `json:"uuid"`
}

// GetGroupsResponse collects the response parameters for the GetGroups method.
type GetGroupsResponse struct {
	Groups []*service.Group `json:"groups"`
	Err    error          `json:"err"`
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
