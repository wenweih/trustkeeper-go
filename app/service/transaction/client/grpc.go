package client

import (
	"fmt"
	"context"
	"github.com/jinzhu/copier"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	endpoint1 "trustkeeper-go/app/service/transaction/pkg/endpoint"
	pb "trustkeeper-go/app/service/transaction/pkg/grpc/pb"
	service "trustkeeper-go/app/service/transaction/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func newGRPCClient(conn *grpc.ClientConn, options []grpc1.ClientOption) (service.TransactionService, error) {
	var assignAssetsToWalletEndpoint endpoint.Endpoint
	{
		assignAssetsToWalletEndpoint = grpc1.NewClient(conn, "pb.Transaction", "AssignAssetsToWallet", encodeAssignAssetsToWalletRequest, decodeAssignAssetsToWalletResponse, pb.AssignAssetsToWalletReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
		AssignAssetsToWalletEndpoint: assignAssetsToWalletEndpoint,
	}, nil
}

// encodeAssignAssetsToWalletRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain AssignAssetsToWallet request to a gRPC request.
func encodeAssignAssetsToWalletRequest(_ context.Context, request interface{}) (interface{}, error) {
	// return nil, errors.New("'Transaction' Encoder is not impelemented")
	r, ok := request.(endpoint1.AssignAssetsToWalletRequest)
  if !ok {
    return nil, fmt.Errorf("request interface to endpoint AssignAssetsToWalletRequest type assertion error")
  }
  pbSimpleAssets := []*pb.SimpleAsset{}
  if err := copier.Copy(&pbSimpleAssets, &r.Assets); err != nil {
    return nil, err
  }
  return &pb.AssignAssetsToWalletRequest{Address: r.Address, SimpleAssets: pbSimpleAssets}, nil
}

// decodeAssignAssetsToWalletResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeAssignAssetsToWalletResponse(_ context.Context, reply interface{}) (interface{}, error) {
	_, ok := reply.(*pb.AssignAssetsToWalletReply)
  if !ok {
    e := fmt.Errorf("pb AssignAssetsToWalletReply type assertion error")
    return endpoint1.AssignAssetsToWalletResponse{Err: e}, e
  }
  return endpoint1.AssignAssetsToWalletResponse{}, nil
}
