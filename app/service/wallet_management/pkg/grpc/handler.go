package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "trustkeeper-go/app/service/wallet_management/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_management/pkg/grpc/pb"
)

// makeCreateChainHandler creates the handler logic
func makeCreateChainHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateChainEndpoint, decodeCreateChainRequest, encodeCreateChainResponse, options...)
}

// decodeCreateChainResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateChain request.
// TODO implement the decoder
func decodeCreateChainRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'WalletManagement' Decoder is not impelemented")
}

// encodeCreateChainResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateChainResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'WalletManagement' Encoder is not impelemented")
}
func (g *grpcServer) CreateChain(ctx context1.Context, req *pb.CreateChainRequest) (*pb.CreateChainReply, error) {
	_, rep, err := g.createChain.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateChainReply), nil
}
