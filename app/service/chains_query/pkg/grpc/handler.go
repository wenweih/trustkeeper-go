package grpc

import (
	"fmt"
	"context"
	"errors"
	endpoint "trustkeeper-go/app/service/chains_query/pkg/endpoint"
	pb "trustkeeper-go/app/service/chains_query/pkg/grpc/pb"

	"github.com/jinzhu/copier"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeBitcoincoreBlockHandler creates the handler logic
func makeBitcoincoreBlockHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.BitcoincoreBlockEndpoint,
		decodeBitcoincoreBlockRequest, encodeBitcoincoreBlockResponse, options...)
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
func (g *grpcServer) BitcoincoreBlock(
	ctx context1.Context, req *pb.BitcoincoreBlockRequest) (*pb.BitcoincoreBlockReply, error) {
	_, rep, err := g.bitcoincoreBlock.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BitcoincoreBlockReply), nil
}

// makeQueryOmniPropertyHandler creates the handler logic
func makeQueryOmniPropertyHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.QueryOmniPropertyEndpoint,
		decodeQueryOmniPropertyRequest, encodeQueryOmniPropertyResponse, options...)
}

// decodeQueryOmniPropertyResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain QueryOmniProperty request.
func decodeQueryOmniPropertyRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.QueryOmniPropertyRequest)
	if !ok {
		return nil, fmt.Errorf("pb QueryOmniPropertyRequest type assersion error")
	}
	return endpoint.QueryOmniPropertyRequest{Propertyid: req.Propertyid}, nil
}

// encodeQueryOmniPropertyResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeQueryOmniPropertyResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.QueryOmniPropertyResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint QueryOmniPropertyResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	property := pb.OmniProperty{}
	if err := copier.Copy(&property, resp.Property); err != nil {
		return nil, err
	}
	return &pb.QueryOmniPropertyReply{OmniProperty: &property}, nil
}
func (g *grpcServer) QueryOmniProperty(
	ctx context1.Context, req *pb.QueryOmniPropertyRequest) (*pb.QueryOmniPropertyReply, error) {
	_, rep, err := g.queryOmniProperty.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.QueryOmniPropertyReply), nil
}
