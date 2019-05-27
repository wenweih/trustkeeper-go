package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "trustkeeper-go/app/service/dashboard/pkg/endpoint"
	pb "trustkeeper-go/app/service/dashboard/pkg/grpc/pb"
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
