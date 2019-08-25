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
func encodeConstructTxBTCResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.ConstructTxBTCResponse)
	if !ok {
		return nil, fmt.Errorf("interface{} to endpoint ConstructTxBTCResponse type assertion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.ConstructTxBTCReply{UnsignedTxHex: resp.UnsignedTxHex, VinAmount: resp.VinAmount}, nil
}

func (g *grpcServer) ConstructTxBTC(ctx context1.Context, req *pb.ConstructTxBTCRequest) (*pb.ConstructTxBTCReply, error) {
	_, rep, err := g.constructTxBTC.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ConstructTxBTCReply), nil
}

// makeSendBTCTxHandler creates the handler logic
func makeSendBTCTxHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendBTCTxEndpoint, decodeSendBTCTxRequest, encodeSendBTCTxResponse, options...)
}

// decodeSendBTCTxResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendBTCTx request.
// TODO implement the decoder
func decodeSendBTCTxRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.SendBTCTxRequest)
	if !ok {
		return nil, fmt.Errorf("pb SendBTCTxRequest type assersion error")
	}
	return endpoint.SendBTCTxRequest{SignedTxHex: req.SignedTxHex}, nil
}

// encodeSendBTCTxResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendBTCTxResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.SendBTCTxResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint SendBTCTxResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.SendBTCTxReply{TxID: resp.TxID}, nil
}
func (g *grpcServer) SendBTCTx(ctx context1.Context, req *pb.SendBTCTxRequest) (*pb.SendBTCTxReply, error) {
	_, rep, err := g.sendBTCTx.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendBTCTxReply), nil
}

// makeQueryBalanceHandler creates the handler logic
func makeQueryBalanceHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.QueryBalanceEndpoint, decodeQueryBalanceRequest, encodeQueryBalanceResponse, options...)
}

// decodeQueryBalanceResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain QueryBalance request.
// TODO implement the decoder
func decodeQueryBalanceRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.QueryBalanceRequest)
	if !ok {
		return nil, fmt.Errorf("pb QueryBalanceRequest type assersion error")
	}
	return endpoint.QueryBalanceRequest{Symbol: req.Symbol, Address: req.Address}, nil
}

// encodeQueryBalanceResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeQueryBalanceResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.QueryBalanceResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint QueryBalanceResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.QueryBalanceReply{Balance: resp.Balance}, nil
}
func (g *grpcServer) QueryBalance(ctx context1.Context, req *pb.QueryBalanceRequest) (*pb.QueryBalanceReply, error) {
	_, rep, err := g.queryBalance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.QueryBalanceReply), nil
}

// makeWalletValidateHandler creates the handler logic
func makeWalletValidateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.WalletValidateEndpoint, decodeWalletValidateRequest, encodeWalletValidateResponse, options...)
}

// decodeWalletValidateResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain WalletValidate request.
// TODO implement the decoder
func decodeWalletValidateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.WalletValidateRequest)
	if !ok {
		return nil, fmt.Errorf("pb WalletValidateRequest type assersion error")
	}
	return endpoint.WalletValidateRequest{Address: req.Address, ChainName: req.ChainName}, nil
}

// encodeWalletValidateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeWalletValidateResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.WalletValidateResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint WalletValidateResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.WalletValidateReply{}, nil
}
func (g *grpcServer) WalletValidate(ctx context1.Context, req *pb.WalletValidateRequest) (*pb.WalletValidateReply, error) {
	_, rep, err := g.walletValidate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WalletValidateReply), nil
}

// makeConstructTxETHHandler creates the handler logic
func makeConstructTxETHHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ConstructTxETHEndpoint, decodeConstructTxETHRequest, encodeConstructTxETHResponse, options...)
}

// decodeConstructTxETHResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ConstructTxETH request.
func decodeConstructTxETHRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.ConstructTxETHRequest)
	if !ok {
		return nil, fmt.Errorf("pb ConstructTxETHRequest type assersion error")
	}
	return endpoint.ConstructTxETHRequest{From: req.From, To: req.To, Amount: req.Amount}, nil
}

// encodeConstructTxETHResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeConstructTxETHResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.ConstructTxETHResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint ConstructTxETHResponse type assertion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.ConstructTxETHReply{UnsignedTxHex: resp.UnsignedTxHex, ChainID: resp.ChainID}, nil
}

func (g *grpcServer) ConstructTxETH(ctx context1.Context, req *pb.ConstructTxETHRequest) (*pb.ConstructTxETHReply, error) {
	_, rep, err := g.constructTxETH.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ConstructTxETHReply), nil
}
