package grpc

import (
	"context"
	"fmt"
	endpoint "trustkeeper-go/app/service/dashboard/pkg/endpoint"
	pb "trustkeeper-go/app/service/dashboard/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
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
