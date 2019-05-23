package grpc

import (
	"context"
	"errors"
	endpoint "trustkeeper-go/app/service/wallet_key/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_key/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeGenerateMnemonicHandler creates the handler logic
func makeGenerateMnemonicHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GenerateMnemonicEndpoint, decodeGenerateMnemonicRequest, encodeGenerateMnemonicResponse, options...)
}

// decodeGenerateMnemonicResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GenerateMnemonic request.
// TODO implement the decoder
func decodeGenerateMnemonicRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'WalletKey' Decoder is not impelemented")
}

// encodeGenerateMnemonicResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGenerateMnemonicResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'WalletKey' Encoder is not impelemented")
}
func (g *grpcServer) GenerateMnemonic(ctx context1.Context, req *pb.GenerateMnemonicRequest) (*pb.GenerateMnemonicReply, error) {
	_, rep, err := g.generateMnemonic.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GenerateMnemonicReply), nil
}
