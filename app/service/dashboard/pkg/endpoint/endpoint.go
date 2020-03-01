package endpoint

import (
	"context"
	"errors"
	"trustkeeper-go/app/service/dashboard/pkg/repository"
	service "trustkeeper-go/app/service/dashboard/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetGroupsRequest collects the request parameters for the GetGroups method.
type GetGroupsRequest struct {
	NamespaceID string `json:"namespaceid"`
}

// GetGroupsResponse collects the response parameters for the GetGroups method.
type GetGroupsResponse struct {
	Groups []*repository.GetGroupsResp `json:"groups"`
	Err    error                       `json:"err"`
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
func (e Endpoints) GetGroups(ctx context.Context, namespaceID string) (groups []*repository.GetGroupsResp, err error) {
	request := GetGroupsRequest{NamespaceID: namespaceID}
	response, err := e.GetGroupsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(GetGroupsResponse).Groups, nil
}

// CreateGroupRequest collects the request parameters for the CreateGroup method.
type CreateGroupRequest struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	NamespaceID string `json:"namespaceid"`
}

// CreateGroupResponse collects the response parameters for the CreateGroup method.
type CreateGroupResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Err  error  `json:"err"`
}

// MakeCreateGroupEndpoint returns an endpoint that invokes CreateGroup on the service.
func MakeCreateGroupEndpoint(s service.DashboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateGroupRequest)
		g, err := s.CreateGroup(ctx, req.UUID, req.Name, req.Desc, req.NamespaceID)
		return CreateGroupResponse{
			Name: g.Name,
			Desc: g.Desc,
			ID:   g.ID,
			Err:  err,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateGroupResponse) Failed() error {
	return r.Err
}

// CreateGroup implements Service. Primarily useful in a client.
func (e Endpoints) CreateGroup(ctx context.Context, uuid, name, desc string, namespaceID string) (group *repository.GetGroupsResp, err error) {
	request := CreateGroupRequest{UUID: uuid, Name: name, Desc: desc, NamespaceID: namespaceID}
	response, err := e.CreateGroupEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	g, ok := response.(CreateGroupResponse)
	if !ok {
		return nil, errors.New("Endpoint CreateGroupResponse type assersion error")
	}
	return &repository.GetGroupsResp{Name: g.Name, Desc: g.Desc, ID: g.ID}, nil
}

// Close implements Service. Primarily useful in a client.
func (e Endpoints) Close() error {
	return nil
}

// UpdateGroupRequest collects the request parameters for the UpdateGroup method.
type UpdateGroupRequest struct {
	GroupID string `json:"groupid"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
}

// UpdateGroupResponse collects the response parameters for the UpdateGroup method.
type UpdateGroupResponse struct {
	Err error `json:"err"`
}

// MakeUpdateGroupEndpoint returns an endpoint that invokes UpdateGroup on the service.
func MakeUpdateGroupEndpoint(s service.DashboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(UpdateGroupRequest)
		if !ok {
			err := errors.New("endpoint UpdateGroupRequest type assertion error")
			return UpdateGroupResponse{Err: err}, nil
		}
		err := s.UpdateGroup(ctx, req.GroupID, req.Name, req.Desc)
		return UpdateGroupResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r UpdateGroupResponse) Failed() error {
	return r.Err
}

// UpdateGroup implements Service. Primarily useful in a client.
func (e Endpoints) UpdateGroup(ctx context.Context, groupID string, name string, desc string) (err error) {
	request := UpdateGroupRequest{Desc: desc, GroupID: groupID, Name: name}
	_, err = e.UpdateGroupEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

// GetGroupAssetRequest collects the request parameters for the GetGroupAsset method.
type GetGroupAssetRequest struct {
	GroupID string `json:"groupid"`
}

// GetGroupAssetResponse collects the response parameters for the GetGroupAsset method.
type GetGroupAssetResponse struct {
	ChainAssets []*repository.ChainAsset `json:"chainassets"`
	Err         error                    `json:"err"`
}

// MakeGetGroupAssetsEndpoint returns an endpoint that invokes GetGroupAsset on the service.
func MakeGetGroupAssetsEndpoint(s service.DashboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupAssetRequest)
		chainAssets, err := s.GetGroupAssets(ctx, req.GroupID)
		return GetGroupAssetResponse{
			ChainAssets: chainAssets,
			Err:         err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupAssetResponse) Failed() error {
	return r.Err
}

// GetGroupAssets implements Service. Primarily useful in a client.
func (e Endpoints) GetGroupAssets(ctx context.Context, groupID string) (chainAssets []*repository.ChainAsset, err error) {
	request := GetGroupAssetRequest{GroupID: groupID}
	response, err := e.GetGroupAssetsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(GetGroupAssetResponse).ChainAssets, nil
}

// GetGroupAssetsRequest collects the request parameters for the GetGroupAssets method.
type GetGroupAssetsRequest struct {
	GroupID string `json:"group_id"`
}

// GetGroupAssetsResponse collects the response parameters for the GetGroupAssets method.
type GetGroupAssetsResponse struct {
	ChainAssets []*repository.ChainAsset `json:"chain_assets"`
	Err         error                    `json:"err"`
}

// Failed implements Failer.
func (r GetGroupAssetsResponse) Failed() error {
	return r.Err
}

// ChangeGroupAssetsRequest collects the request parameters for the ChangeGroupAssets method.
type ChangeGroupAssetsRequest struct {
	ChainAssets []*repository.ChainAsset `json:"chainassets"`
	Groupid     string                   `json:"groupid"`
}

// ChangeGroupAssetsResponse collects the response parameters for the ChangeGroupAssets method.
type ChangeGroupAssetsResponse struct {
	Err         error                    `json:"err"`
	ChainAssets []*repository.ChainAsset `json:"chainassets"`
}

// MakeChangeGroupAssetsEndpoint returns an endpoint that invokes ChangeGroupAssets on the service.
func MakeChangeGroupAssetsEndpoint(s service.DashboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeGroupAssetsRequest)
		resp, err := s.ChangeGroupAssets(ctx, req.ChainAssets, req.Groupid)
		return ChangeGroupAssetsResponse{ChainAssets: resp, Err: err}, nil
	}
}

// Failed implements Failer.
func (r ChangeGroupAssetsResponse) Failed() error {
	return r.Err
}

// ChangeGroupAssets implements Service. Primarily useful in a client.
func (e Endpoints) ChangeGroupAssets(ctx context.Context, chainAssets []*repository.ChainAsset, groupid string) (result []*repository.ChainAsset, err error) {
	request := ChangeGroupAssetsRequest{ChainAssets: chainAssets, Groupid: groupid}
	resp, err := e.ChangeGroupAssetsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(ChangeGroupAssetsResponse).ChainAssets, nil
}

// AddAssetRequest collects the request parameters for the AddAsset method.
type AddAssetRequest struct {
	Groupid  string `json:"groupid"`
	Chainid  string `json:"chainid"`
	Symbol   string `json:"symbol"`
	Identify string `json:"identify"`
	Decimal  string `json:"decimal"`
}

// AddAssetResponse collects the response parameters for the AddAsset method.
type AddAssetResponse struct {
	Asset *repository.SimpleAsset `json:"asset"`
	Err   error                   `json:"err"`
}

// MakeAddAssetEndpoint returns an endpoint that invokes AddAsset on the service.
func MakeAddAssetEndpoint(s service.DashboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddAssetRequest)
		asset, err := s.AddAsset(ctx, req.Groupid, req.Chainid, req.Symbol, req.Identify, req.Decimal)
		return AddAssetResponse{
			Asset: asset,
			Err:   err,
		}, nil
	}
}

// Failed implements Failer.
func (r AddAssetResponse) Failed() error {
	return r.Err
}

// AddAsset implements Service. Primarily useful in a client.
func (e Endpoints) AddAsset(ctx context.Context, groupid string, chainid string, symbol string, identify string, decimal string) (asset *repository.SimpleAsset, err error) {
	request := AddAssetRequest{
		Chainid:  chainid,
		Decimal:  decimal,
		Groupid:  groupid,
		Identify: identify,
		Symbol:   symbol,
	}
	response, err := e.AddAssetEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(AddAssetResponse).Asset, nil
}
