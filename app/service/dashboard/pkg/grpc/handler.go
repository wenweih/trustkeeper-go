package grpc

import (
	"fmt"
	"context"
	"errors"
	endpoint "trustkeeper-go/app/service/dashboard/pkg/endpoint"
	pb "trustkeeper-go/app/service/dashboard/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeGetGroupsHandler creates the handler logic
func makeGetGroupsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupsEndpoint, decodeGetGroupsRequest, encodeGetGroupsResponse, options...)
}

// decodeGetGroupsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetGroups request.
// TODO implement the decoder
func decodeGetGroupsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Dashboard' Decoder is not impelemented")
}

// encodeGetGroupsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetGroupsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Dashboard' Encoder is not impelemented")
}
func (g *grpcServer) GetGroups(ctx context1.Context, req *pb.GetGroupsRequest) (*pb.GetGroupsReply, error) {
	_, rep, err := g.getGroups.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupsReply), nil
}

// makeCreateGroupHandler creates the handler logic
func makeCreateGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateGroupEndpoint, decodeCreateGroupRequest, encodeCreateGroupResponse, options...)
}

// decodeCreateGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateGroup request.
// TODO implement the decoder
func decodeCreateGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.CreateGroupRequest)
	if !ok {
		return nil, fmt.Errorf("interface{} to pb CreateGroupRequest type assertion error")
	}
	return endpoint.CreateGroupRequest{UUID: req.Uuid, Name: req.Name, Desc: req.Desc}, nil
}

// encodeCreateGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.CreateGroupResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.CreateGroupReply{Result: resp.Result}, nil
}

func (g *grpcServer) CreateGroup(ctx context1.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupReply, error) {
	_, rep, err := g.createGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateGroupReply), nil
}
