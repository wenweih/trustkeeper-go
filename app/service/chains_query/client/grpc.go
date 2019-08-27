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

	var eRC20TokenInfoEndpoint endpoint.Endpoint
	{
		eRC20TokenInfoEndpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "ERC20TokenInfo",
			encodeERC20TokenInfoRequest, decodeERC20TokenInfoResponse, pb.ERC20TokenInfoReply{}, options...).Endpoint()
	}

	var constructTxBTCEndpoint endpoint.Endpoint
	{
		constructTxBTCEndpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "ConstructTxBTC",
			encodeConstructTxBTCRequest, decodeConstructTxBTCResponse, pb.ConstructTxBTCReply{}, options...).Endpoint()
	}

	var sendBTCTxEndpoint endpoint.Endpoint
	{
		sendBTCTxEndpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "SendBTCTx",
			encodeSendBTCTxRequest, decodeSendBTCTxResponse, pb.SendBTCTxReply{}, options...).Endpoint()
	}

	var queryBalanceEndpoint endpoint.Endpoint
	{
		queryBalanceEndpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "QueryBalance",
			encodeQueryBalanceRequest, decodeQueryBalanceResponse, pb.QueryBalanceReply{}, options...).Endpoint()
	}

	var walletValidateEndpoint endpoint.Endpoint
	{
		walletValidateEndpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "WalletValidate",
			encodeWalletValidateRequest, decodeWalletValidateResponse, pb.WalletValidateReply{}, options...).Endpoint()
	}

	var constructTxETHEndpoint endpoint.Endpoint
	{
		constructTxETHEndpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "ConstructTxETH",
			encodeConstructTxETHRequest, decodeConstructTxETHResponse, pb.ConstructTxETHReply{}, options...).Endpoint()
	}

	var sendETHTxEndpoint endpoint.Endpoint
	{
		sendETHTxEndpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "SendETHTx",
			encodeSendETHTxRequest, decodeSendETHTxResponse, pb.SendETHTxReply{}, options...).Endpoint()
	}

	var constructTxERC20Endpoint endpoint.Endpoint
	{
		constructTxERC20Endpoint = grpc1.NewClient(conn, "pb.ChainsQuery", "ConstructTxERC20",
			encodeConstructTxERC20Request, decodeConstructTxERC20Response, pb.ConstructTxERC20Reply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
		BitcoincoreBlockEndpoint:  bitcoincoreBlockEndpoint,
		ConstructTxBTCEndpoint:    constructTxBTCEndpoint,
		ConstructTxERC20Endpoint:  constructTxERC20Endpoint,
		ConstructTxETHEndpoint:    constructTxETHEndpoint,
		ERC20TokenInfoEndpoint:    eRC20TokenInfoEndpoint,
		QueryBalanceEndpoint:      queryBalanceEndpoint,
		QueryOmniPropertyEndpoint: queryOmniPropertyEndpoint,
		SendBTCTxEndpoint:         sendBTCTxEndpoint,
		SendETHTxEndpoint:         sendETHTxEndpoint,
		WalletValidateEndpoint:    walletValidateEndpoint,
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

// encodeERC20TokenInfoRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain ERC20TokenInfo request to a gRPC request.
func encodeERC20TokenInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.ERC20TokenInfoRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint ERC20TokenInfoRequest type assertion error")
	}
	return &pb.ERC20TokenInfoRequest{TokenHex: r.TokenHex}, nil
}

// decodeERC20TokenInfoResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeERC20TokenInfoResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.ERC20TokenInfoReply)
	if !ok {
		return nil, fmt.Errorf("pb ERC20TokenInfoReply type assertion error")
	}
	token := repository.ERC20Token{}
	if err := copier.Copy(&token, resp.ERC20Token); err != nil {
		return nil, err
	}
	return endpoint1.ERC20TokenInfoResponse{Token: &token}, nil
}

// encodeConstructTxBTCRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain ConstructTxBTC request to a gRPC request.
func encodeConstructTxBTCRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.ConstructTxBTCRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint ConstructTxBTCRequest type assertion error")
	}
	return &pb.ConstructTxBTCRequest{Amount: r.Amount, From: r.From, To: r.To}, nil
}

// decodeConstructTxBTCResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeConstructTxBTCResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.ConstructTxBTCReply)
	if !ok {
		e := fmt.Errorf("pb ConstructTxBTCReply type assertion error")
		return endpoint1.ConstructTxBTCResponse{Err: e}, e
	}
	return endpoint1.ConstructTxBTCResponse{UnsignedTxHex: resp.UnsignedTxHex, VinAmount: resp.VinAmount}, nil
}

// encodeSendBTCTxRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain SendBTCTx request to a gRPC request.
func encodeSendBTCTxRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.SendBTCTxRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint SendBTCTxRequest type assertion error")
	}
	return &pb.SendBTCTxRequest{SignedTxHex: r.SignedTxHex}, nil
}

// decodeSendBTCTxResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeSendBTCTxResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.SendBTCTxReply)
	if !ok {
		e := fmt.Errorf("pb SendBTCTxReply type assertion error")
		return endpoint1.SendBTCTxResponse{Err: e}, e
	}
	return endpoint1.SendBTCTxResponse{TxID: resp.TxID}, nil
}

// encodeSendETHTxRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain SendETHTx request to a gRPC request.
func encodeSendETHTxRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.SendETHTxRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint SendETHTxRequest type assertion error")
	}
	return &pb.SendETHTxRequest{SignedTxHex: r.SignedTxHex}, nil
}

// decodeSendETHTxResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeSendETHTxResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.SendETHTxReply)
	if !ok {
		e := fmt.Errorf("pb SendETHTxReply type assertion error")
		return endpoint1.SendETHTxResponse{Err: e}, e
	}
	return endpoint1.SendETHTxResponse{TxID: resp.TxID}, nil
}

// encodeQueryBalanceRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain QueryBalance request to a gRPC request.
func encodeQueryBalanceRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.QueryBalanceRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint QueryBalanceRequest type assertion error")
	}
	return &pb.QueryBalanceRequest{Symbol: r.Symbol, Address: r.Address}, nil
}

// decodeQueryBalanceResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeQueryBalanceResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.QueryBalanceReply)
	if !ok {
		e := fmt.Errorf("pb QueryBalanceReply type assertion error")
		return endpoint1.QueryBalanceResponse{Err: e}, e
	}
	return endpoint1.QueryBalanceResponse{Balance: resp.Balance}, nil
}

// encodeWalletValidateRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain WalletValidate request to a gRPC request.
func encodeWalletValidateRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.WalletValidateRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint WalletValidateRequest type assertion error")
	}
	return &pb.WalletValidateRequest{Address: r.Address, ChainName: r.ChainName}, nil
}

// decodeWalletValidateResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeWalletValidateResponse(_ context.Context, reply interface{}) (interface{}, error) {
	_, ok := reply.(*pb.WalletValidateReply)
	if !ok {
		e := fmt.Errorf("pb WalletValidateReply type assertion error")
		return endpoint1.WalletValidateResponse{Err: e}, e
	}
	return endpoint1.WalletValidateResponse{}, nil
}

// encodeConstructTxETHRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain ConstructTxETH request to a gRPC request.
func encodeConstructTxETHRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.ConstructTxETHRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint ConstructTxETHRequest type assertion error")
	}
	return &pb.ConstructTxETHRequest{Amount: r.Amount, From: r.From, To: r.To}, nil
}

// decodeConstructTxETHResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeConstructTxETHResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.ConstructTxETHReply)
	if !ok {
		e := fmt.Errorf("pb ConstructTxETHReply type assertion error")
		return endpoint1.ConstructTxETHResponse{Err: e}, e
	}
	return endpoint1.ConstructTxETHResponse{UnsignedTxHex: resp.UnsignedTxHex, ChainID: resp.ChainID}, nil
}

// encodeConstructTxERC20Request is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain ConstructTxERC20 request to a gRPC request.
func encodeConstructTxERC20Request(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.ConstructTxERC20Request)
	if !ok {
		return nil, fmt.Errorf("endpoint ConstructTxERC20Request type assertion error")
	}
	return &pb.ConstructTxERC20Request{Amount: r.Amount, From: r.From, To: r.To, Contract: r.Contract}, nil
}

// decodeConstructTxERC20Response is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeConstructTxERC20Response(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.ConstructTxERC20Reply)
	if !ok {
		e := fmt.Errorf("pb ConstructTxERC20Reply type assertion error")
		return endpoint1.ConstructTxERC20Response{Err: e}, e
	}
	return endpoint1.ConstructTxERC20Response{UnsignedTxHex: resp.UnsignedTxHex, ChainID: resp.ChainID}, nil
}
