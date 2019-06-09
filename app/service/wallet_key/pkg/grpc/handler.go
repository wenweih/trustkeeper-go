package grpc

import (
	"context"
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
	req := r.(*pb.GenerateMnemonicRequest)
	return endpoint.GenerateMnemonicRequest{Uuid: req.Uuid}, nil
}

// encodeGenerateMnemonicResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeGenerateMnemonicResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GenerateMnemonicResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.GenerateMnemonicReply{Xpub: resp.Xpub}, nil
}
func (g *grpcServer) GenerateMnemonic(ctx context1.Context, req *pb.GenerateMnemonicRequest) (*pb.GenerateMnemonicReply, error) {
	_, rep, err := g.generateMnemonic.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GenerateMnemonicReply), nil
}
