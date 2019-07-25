package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "trustkeeper-go/app/service/transaction/pkg/endpoint"
	pb "trustkeeper-go/app/service/transaction/pkg/grpc/pb"
)

// makeAssignAssetsToWalletHandler creates the handler logic
func makeAssignAssetsToWalletHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AssignAssetsToWalletEndpoint, decodeAssignAssetsToWalletRequest, encodeAssignAssetsToWalletResponse, options...)
}

// decodeAssignAssetsToWalletResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain AssignAssetsToWallet request.
// TODO implement the decoder
func decodeAssignAssetsToWalletRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Transaction' Decoder is not impelemented")
}

// encodeAssignAssetsToWalletResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAssignAssetsToWalletResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Transaction' Encoder is not impelemented")
}
func (g *grpcServer) AssignAssetsToWallet(ctx context1.Context, req *pb.AssignAssetsToWalletRequest) (*pb.AssignAssetsToWalletReply, error) {
	_, rep, err := g.assignAssetsToWallet.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AssignAssetsToWalletReply), nil
}
