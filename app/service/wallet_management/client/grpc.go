package client

import (
	"fmt"
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	endpoint1 "trustkeeper-go/app/service/wallet_management/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_management/pkg/grpc/pb"
	service "trustkeeper-go/app/service/wallet_management/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func newGRPCClient(conn *grpc.ClientConn, options []grpc1.ClientOption) (service.WalletManagementService, error) {
	var createChainEndpoint endpoint.Endpoint
	{
		createChainEndpoint = grpc1.NewClient(conn, "pb.WalletManagement", "CreateChain", encodeCreateChainRequest, decodeCreateChainResponse, pb.CreateChainReply{}, options...).Endpoint()
	}

	var assignedXpubToGroupEndpoint endpoint.Endpoint
	{
		assignedXpubToGroupEndpoint = grpc1.NewClient(conn, "pb.WalletManagement", "AssignedXpubToGroup", encodeAssignedXpubToGroupRequest, decodeAssignedXpubToGroupResponse, pb.AssignedXpubToGroupReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
		CreateChainEndpoint: createChainEndpoint,
		AssignedXpubToGroupEndpoint: assignedXpubToGroupEndpoint,
	}, nil
}

// encodeCreateChainRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain CreateChain request to a gRPC request.
func encodeCreateChainRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.CreateChainRequest)
	if !ok {
		return nil, fmt.Errorf("request interface to endpoint.CreateChainRequest type assertion error")
	}
	return &pb.CreateChainRequest{
		Symbol: r.Symbol,
		Bitid: 	r.Bit44ID,
    Status: r.Status}, nil
}

// decodeCreateChainResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCreateChainResponse(_ context.Context, reply interface{}) (interface{}, error) {
	_, ok := reply.(*pb.CreateChainReply)
  if !ok{
		e := fmt.Errorf("pb CreateChainReply type assertion error")
    return &endpoint1.CreateChainResponse{Err: e}, e
  }
  return &endpoint1.CreateChainResponse{}, nil
}

// encodeUpdateXpubStateRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain UpdateXpubState request to a gRPC request.
func encodeAssignedXpubToGroupRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.AssignedXpubToGroupRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint AssignedXpubToGroupRequest type assertion error")
	}
	return &pb.AssignedXpubToGroupRequest{Groupid: r.GroupID}, nil
}

// decodeUpdateXpubStateResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeAssignedXpubToGroupResponse(_ context.Context, reply interface{}) (interface{}, error) {
	_, ok := reply.(*pb.AssignedXpubToGroupReply)
	if !ok {
		return nil, fmt.Errorf("pb UpdateXpubStateReply type assertion error")
	}
	return &endpoint1.AssignedXpubToGroupResponse{}, nil
}
