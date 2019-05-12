package grpc

import (
	"context"
	"errors"
	endpoint "trustkeeper-go/app/service/account/pkg/endpoint"
	pb "trustkeeper-go/app/service/account/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

func makeCreateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)
}

func decodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateRequest)
	return endpoint.CreateRequest{Email: req.Email, Password: req.Password}, nil
}

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

func makeSignHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SignEndpoint, decodeSignRequest, encodeSignResponse, options...)
}

func decodeSignRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Account' Decoder is not impelemented")
}

func encodeSignResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Account' Encoder is not impelemented")
}
func (g *grpcServer) Sign(ctx context1.Context, req *pb.SignRequest) (*pb.SignReply, error) {
	_, rep, err := g.sign.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SignReply), nil
}
