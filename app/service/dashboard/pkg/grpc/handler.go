package grpc

import (
	"context"
	"fmt"
	endpoint "trustkeeper-go/app/service/dashboard/pkg/endpoint"
	pb "trustkeeper-go/app/service/dashboard/pkg/grpc/pb"
	"trustkeeper-go/app/service/dashboard/pkg/repository"

	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/jinzhu/copier"
	context1 "golang.org/x/net/context"
)

func makeGetGroupsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupsEndpoint, decodeGetGroupsRequest, encodeGetGroupsResponse, options...)
}

func decodeGetGroupsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.GetGroupsRequest)
	if !ok {
		return nil, fmt.Errorf("interface{} to pb GetGroupsRequest type assertion error")
	}
	return endpoint.GetGroupsRequest{NamespaceID: req.NamespaceID}, nil
}

func encodeGetGroupsResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.GetGroupsResponse)
	if !ok {
		return nil, fmt.Errorf("interface{} to endpoint GetGroupsResponse type assertion error")
	}

	pbGroups := make([]*pb.Group, len(resp.Groups))
	for i, g := range resp.Groups {
		pbGroups[i] = &pb.Group{Name: g.Name, Desc: g.Desc, Id: g.ID}
	}
	return &pb.GetGroupsReply{Groups: pbGroups}, nil
}

func (g *grpcServer) GetGroups(ctx context1.Context, req *pb.GetGroupsRequest) (*pb.GetGroupsReply, error) {
	_, rep, err := g.getGroups.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupsReply), nil
}

func makeCreateGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateGroupEndpoint, decodeCreateGroupRequest, encodeCreateGroupResponse, options...)
}

func decodeCreateGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.CreateGroupRequest)
	if !ok {
		return nil, fmt.Errorf("interface{} to pb CreateGroupRequest type assertion error")
	}
	return endpoint.CreateGroupRequest{UUID: req.Uuid, Name: req.Name, Desc: req.Desc, NamespaceID: req.NamespaceID}, nil
}

func encodeCreateGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.CreateGroupResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.CreateGroupReply{Group: &pb.Group{Name: resp.Name, Desc: resp.Desc, Id: resp.ID}}, nil
}

func (g *grpcServer) CreateGroup(ctx context1.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupReply, error) {
	_, rep, err := g.createGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateGroupReply), nil
}

func makeUpdateGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateGroupEndpoint, decodeUpdateGroupRequest, encodeUpdateGroupResponse, options...)
}

func decodeUpdateGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.UpdateGroupRequest)
	if !ok {
		return nil, fmt.Errorf("pb UpdateGroupRequest type assertion error")
	}
	return endpoint.UpdateGroupRequest{GroupID: req.Groupid, Name: req.Name, Desc: req.Desc}, nil
}

func encodeUpdateGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	_, ok := r.(endpoint.UpdateGroupResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint UpdateGroupRequest type assertion error")
	}
	return &pb.UpdateGroupReply{}, nil
}

func (g *grpcServer) UpdateGroup(ctx context1.Context, req *pb.UpdateGroupRequest) (*pb.UpdateGroupReply, error) {
	_, rep, err := g.updateGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateGroupReply), nil
}

func makeGetGroupAssetsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupAssetsEndpoint, decodeGetGroupAssetRequest, encodeGetGroupAssetResponse, options...)
}

func decodeGetGroupAssetRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.GetGroupAssetRequest)
	if !ok {
		return nil, fmt.Errorf("interface{} to pb GetGroupAssetRequest type assertion error")
	}
	return endpoint.GetGroupAssetRequest{GroupID: req.Groupid}, nil
}

func encodeGetGroupAssetResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.GetGroupAssetResponse)
	if !ok {
		return nil, fmt.Errorf("interface{} to endpoint GetGroupAssetResponse type assertion error")
	}

	pbChainAssets := make([]*pb.ChainAsset, 0)
	if err := copier.Copy(&pbChainAssets, &resp.ChainAssets); err != nil {
		return nil, err
	}
	return &pb.GetGroupAssetReply{Chainassets: pbChainAssets}, nil
}
func (g *grpcServer) GetGroupAsset(ctx context1.Context, req *pb.GetGroupAssetRequest) (*pb.GetGroupAssetReply, error) {
	_, rep, err := g.getGroupAssets.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupAssetReply), nil
}

func makeChangeGroupAssetsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ChangeGroupAssetsEndpoint, decodeChangeGroupAssetsRequest, encodeChangeGroupAssetsResponse, options...)
}

func decodeChangeGroupAssetsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.ChangeGroupAssetsRequest)
	if !ok {
		return nil, fmt.Errorf("interface{} to pb ChangeGroupAssetsRequest type assertion error")
	}
	endpointChainAssets := []*repository.ChainAsset{}
	if err := copier.Copy(&endpointChainAssets, req.Chainassets); err != nil {
		return nil, err
	}
	return endpoint.ChangeGroupAssetsRequest{ChainAssets: endpointChainAssets, Groupid: req.Groupid}, nil
}

func encodeChangeGroupAssetsResponse(_ context.Context, r interface{}) (interface{}, error) {
	response, ok := r.(endpoint.ChangeGroupAssetsResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint ChangeGroupAssetsResponse type assertion error")
	}
	pbChainAssets := []*pb.ChainAsset{}
	if err := copier.Copy(&pbChainAssets, &response.ChainAssets); err != nil {
		return nil, err
	}
	return &pb.ChangeGroupAssetsReply{Chainassets: pbChainAssets}, nil
}
func (g *grpcServer) ChangeGroupAssets(ctx context1.Context, req *pb.ChangeGroupAssetsRequest) (*pb.ChangeGroupAssetsReply, error) {
	_, rep, err := g.changeGroupAssets.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ChangeGroupAssetsReply), nil
}

func makeAddAssetHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddAssetEndpoint, decodeAddAssetRequest, encodeAddAssetResponse, options...)
}

func decodeAddAssetRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.AddAssetRequest)
	if !ok {
		return nil, fmt.Errorf("interface{} to pb AddAssetRequest type assertion error")
	}
	return endpoint.AddAssetRequest{
		Groupid: req.Groupid,
		Chainid: req.Chainid,
		Symbol: req.Symbol,
		Identify: req.Identify,
		Decimal: req.Decimal}, nil
}

func encodeAddAssetResponse(_ context.Context, r interface{}) (interface{}, error) {
	response, ok := r.(endpoint.AddAssetResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint AddAssetResponse type assertion error")
	}
	pbSimpleAsset := pb.SimpleAsset{}
	if err := copier.Copy(&pbSimpleAsset, response.Asset); err != nil {
		return nil, err
	}
	return &pb.AddAssetReply{Asset: &pbSimpleAsset}, nil
}

func (g *grpcServer) AddAsset(ctx context1.Context, req *pb.AddAssetRequest) (*pb.AddAssetReply, error) {
	_, rep, err := g.addAsset.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddAssetReply), nil
}
