package grpc

import (
	"context"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "trustkeeper-go/app/service/account/pkg/endpoint"
	pb "trustkeeper-go/app/service/account/pkg/grpc/pb"
)

// makeCreateHandler creates the handler logic
func makeCreateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)
}

// decodeCreateResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Create request.
// TODO implement the decoder
func decodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateRequest)
	return endpoint.CreateRequest{Email: req.Email, Password: req.Password}, nil
}

// encodeCreateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.CreateResponse)
	if resp.E1 != nil {
		return &pb.CreateReply{Result: false}, resp.E1
	}
	return &pb.CreateReply{Result: true}, nil
}
func (g *grpcServer) Create(ctx context1.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	_, rep, err := g.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateReply), nil
}
