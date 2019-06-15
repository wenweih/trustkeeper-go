package grpc

import (
	"context"
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
func decodeCreateChainRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateChainRequest)
	return endpoint.CreateChainRequest{Symbol: req.Symbol, Bit44ID: req.Bitid, Status: req.Status}, nil
}

// encodeCreateChainResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeCreateChainResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.CreateChainResponse)
	if resp.Err != nil {
		return &pb.CreateChainReply{Result: false}, resp.Err
	}
	return &pb.CreateChainReply{Result: true}, nil
}

func (g *grpcServer) CreateChain(ctx context1.Context, req *pb.CreateChainRequest) (*pb.CreateChainReply, error) {
	_, rep, err := g.createChain.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateChainReply), nil
}
