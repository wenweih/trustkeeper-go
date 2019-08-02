package client

import (
	"fmt"
	"context"
	"errors"
	"github.com/jinzhu/copier"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	endpoint1 "trustkeeper-go/app/service/chains_query/pkg/endpoint"
	pb "trustkeeper-go/app/service/chains_query/pkg/grpc/pb"
	service "trustkeeper-go/app/service/chains_query/pkg/service"
	"trustkeeper-go/app/service/chains_query/pkg/repository"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func newGRPCClient(conn *grpc.ClientConn, options []grpc1.ClientOption) (service.ChainsQueryService, error) {
	var bitcoincoreBlockEndpoint endpoint.Endpoint
	{
		bitcoincoreBlockEndpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "BitcoincoreBlock",
			encodeBitcoincoreBlockRequest, decodeBitcoincoreBlockResponse, pb.BitcoincoreBlockReply{}, options...).Endpoint()
	}

	var queryOmniPropertyEndpoint endpoint.Endpoint
	{
		queryOmniPropertyEndpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "QueryOmniProperty",
			encodeQueryOmniPropertyRequest, decodeQueryOmniPropertyResponse, pb.QueryOmniPropertyReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
		BitcoincoreBlockEndpoint:  bitcoincoreBlockEndpoint,
		QueryOmniPropertyEndpoint: queryOmniPropertyEndpoint,
	}, nil
}

// encodeBitcoincoreBlockRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain BitcoincoreBlock request to a gRPC request.
func encodeBitcoincoreBlockRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'ChainsQuery' Encoder is not impelemented")
}

// decodeBitcoincoreBlockResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeBitcoincoreBlockResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'ChainsQuery' Decoder is not impelemented")
}

// encodeQueryOmniPropertyRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain QueryOmniProperty request to a gRPC request.
func encodeQueryOmniPropertyRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.QueryOmniPropertyRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint QueryOmniPropertyRequest type assertion error")
	}
	return &pb.QueryOmniPropertyRequest{Propertyid: r.Propertyid}, nil
}

// decodeQueryOmniPropertyResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeQueryOmniPropertyResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.QueryOmniPropertyReply)
	if !ok {
		return nil, fmt.Errorf("pb QueryOmniPropertyReply type assertion error")
	}
	property := repository.OmniProperty{}
	if err := copier.Copy(&property, resp.OmniProperty); err != nil {
		return nil, err
	}
	return endpoint1.QueryOmniPropertyResponse{Property: &property}, nil
}
