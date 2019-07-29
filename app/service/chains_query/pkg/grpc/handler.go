package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "trustkeeper-go/app/service/chains_query/pkg/endpoint"
	pb "trustkeeper-go/app/service/chains_query/pkg/grpc/pb"
)

// makeBitcoincoreBlockHandler creates the handler logic
func makeBitcoincoreBlockHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.BitcoincoreBlockEndpoint, decodeBitcoincoreBlockRequest, encodeBitcoincoreBlockResponse, options...)
}

// decodeBitcoincoreBlockResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain BitcoincoreBlock request.
// TODO implement the decoder
func decodeBitcoincoreBlockRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'ChainsQuery' Decoder is not impelemented")
}

// encodeBitcoincoreBlockResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeBitcoincoreBlockResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'ChainsQuery' Encoder is not impelemented")
}
func (g *grpcServer) BitcoincoreBlock(ctx context1.Context, req *pb.BitcoincoreBlockRequest) (*pb.BitcoincoreBlockReply, error) {
	_, rep, err := g.bitcoincoreBlock.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BitcoincoreBlockReply), nil
}
