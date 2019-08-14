package grpc

import (
	"context"
	"errors"
	"fmt"
	endpoint "trustkeeper-go/app/service/chains_query/pkg/endpoint"
	pb "trustkeeper-go/app/service/chains_query/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/jinzhu/copier"
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

// makeERC20TokenInfoHandler creates the handler logic
func makeERC20TokenInfoHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ERC20TokenInfoEndpoint, decodeERC20TokenInfoRequest, encodeERC20TokenInfoResponse, options...)
}

// decodeERC20TokenInfoResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ERC20TokenInfo request.
func decodeERC20TokenInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.ERC20TokenInfoRequest)
	if !ok {
		return nil, fmt.Errorf("pb ERC20TokenInfoRequest type assersion error")
	}
	return endpoint.ERC20TokenInfoRequest{TokenHex: req.TokenHex}, nil
}

// encodeERC20TokenInfoResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeERC20TokenInfoResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.ERC20TokenInfoResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint ERC20TokenInfoResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	token := pb.ERC20Token{}
	if err := copier.Copy(&token, resp.Token); err != nil {
		return nil, err
	}
	return &pb.ERC20TokenInfoReply{ERC20Token: &token}, nil
}

func (g *grpcServer) ERC20TokenInfo(ctx context1.Context, req *pb.ERC20TokenInfoRequest) (*pb.ERC20TokenInfoReply, error) {
	_, rep, err := g.eRC20TokenInfo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ERC20TokenInfoReply), nil
}

// makeConstructTxBTCHandler creates the handler logic
func makeConstructTxBTCHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ConstructTxBTCEndpoint, decodeConstructTxBTCRequest, encodeConstructTxBTCResponse, options...)
}

// decodeConstructTxBTCResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ConstructTxBTC request.
func decodeConstructTxBTCRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.ConstructTxBTCRequest)
	if !ok {
		return nil, fmt.Errorf("pb ConstructTxBTCRequest type assersion error")
	}
	return endpoint.ConstructTxBTCRequest{From: req.From, To: req.To, Amount: req.Amount}, nil
}

// encodeConstructTxBTCResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeConstructTxBTCResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.ConstructTxBTCResponse)
	if !ok {
		return nil, fmt.Errorf("interface{} to endpoint ConstructTxBTCResponse type assertion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.ConstructTxBTCReply{UnsignedTxHex: resp.UnsignedTxHex}, nil
}

func (g *grpcServer) ConstructTxBTC(ctx context1.Context, req *pb.ConstructTxBTCRequest) (*pb.ConstructTxBTCReply, error) {
	_, rep, err := g.constructTxBTC.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ConstructTxBTCReply), nil
}
