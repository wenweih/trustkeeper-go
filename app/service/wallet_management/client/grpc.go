package client

import (
	"context"
	"fmt"
	endpoint1 "trustkeeper-go/app/service/wallet_management/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_management/pkg/grpc/pb"
	"trustkeeper-go/app/service/wallet_management/pkg/repository"
	service "trustkeeper-go/app/service/wallet_management/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	"github.com/jinzhu/copier"
	grpc "google.golang.org/grpc"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func newGRPCClient(conn *grpc.ClientConn, options []grpc1.ClientOption) (service.WalletManagementService, error) {
	var createChainEndpoint endpoint.Endpoint
	{
		createChainEndpoint = grpc1.NewClient(conn, "pb.WalletManagement", "CreateChain", encodeCreateChainRequest, decodeCreateChainResponse, pb.CreateChainReply{}, options...).Endpoint()
	}

	var assignedXpubToGroupEndpoint endpoint.Endpoint
	{
		assignedXpubToGroupEndpoint = grpc1.NewClient(conn, "pb.WalletManagement", "AssignedXpubToGroup",
			encodeAssignedXpubToGroupRequest, decodeAssignedXpubToGroupResponse, pb.AssignedXpubToGroupReply{}, options...).Endpoint()
	}

	var getChainsEndpoint endpoint.Endpoint
	{
		getChainsEndpoint = grpc1.NewClient(conn, "pb.WalletManagement", "GetChains",
			encodeGetChainsRequest, decodeGetChainsResponse, pb.GetChainsReply{}, options...).Endpoint()
	}

	var createWalletEndpoint endpoint.Endpoint
	{
		createWalletEndpoint = grpc1.NewClient(conn, "pb.WalletManagement", "CreateWallet",
			encodeCreateWalletRequest, decodeCreateWalletResponse, pb.CreateWalletReply{}, options...).Endpoint()
	}

	var getWalletsEndpoint endpoint.Endpoint
	{
		getWalletsEndpoint = grpc1.NewClient(conn, "pb.WalletManagement", "GetWallets",
			encodeGetWalletsRequest, decodeGetWalletsResponse, pb.GetWalletsReply{}, options...).Endpoint()
	}

	var queryWalletsForGroupByChainNameEndpoint endpoint.Endpoint
	{
		queryWalletsForGroupByChainNameEndpoint = grpc1.NewClient(conn, "pb.WalletManagement", "QueryWalletsForGroupByChainName",
			encodeQueryWalletsForGroupByChainNameRequest,
			decodeQueryWalletsForGroupByChainNameResponse,
			pb.QueryWalletsForGroupByChainNameReply{}, options...).Endpoint()
	}

	var queryWalletHDEndpoint endpoint.Endpoint
	{
		queryWalletHDEndpoint = grpc1.NewClient(conn, "pb.WalletManagement", "QueryWalletHD",
			encodeQueryWalletHDRequest, decodeQueryWalletHDResponse, pb.QueryWalletHDReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
		CreateChainEndpoint:                     createChainEndpoint,
		AssignedXpubToGroupEndpoint:             assignedXpubToGroupEndpoint,
		CreateWalletEndpoint:                    createWalletEndpoint,
		GetChainsEndpoint:                       getChainsEndpoint,
		GetWalletsEndpoint:                      getWalletsEndpoint,
		QueryWalletHDEndpoint:                   queryWalletHDEndpoint,
		QueryWalletsForGroupByChainNameEndpoint: queryWalletsForGroupByChainNameEndpoint,
	}, nil
}

// encodeCreateChainRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain CreateChain request to a gRPC request.
func encodeCreateChainRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.CreateChainRequest)
	if !ok {
		return nil, fmt.Errorf("request interface to endpoint.CreateChainRequest type assertion error")
	}
	return &pb.CreateChainRequest{
		Symbol: r.Symbol,
		Bitid:  r.Bit44ID,
		Status: r.Status}, nil
}

// decodeCreateChainResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCreateChainResponse(_ context.Context, reply interface{}) (interface{}, error) {
	_, ok := reply.(*pb.CreateChainReply)
	if !ok {
		e := fmt.Errorf("pb CreateChainReply type assertion error")
		return &endpoint1.CreateChainResponse{Err: e}, e
	}
	return &endpoint1.CreateChainResponse{}, nil
}

// encodeUpdateXpubStateRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain UpdateXpubState request to a gRPC request.
func encodeAssignedXpubToGroupRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.AssignedXpubToGroupRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint AssignedXpubToGroupRequest type assertion error")
	}
	return &pb.AssignedXpubToGroupRequest{Groupid: r.GroupID}, nil
}

// decodeUpdateXpubStateResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeAssignedXpubToGroupResponse(_ context.Context, reply interface{}) (interface{}, error) {
	_, ok := reply.(*pb.AssignedXpubToGroupReply)
	if !ok {
		return nil, fmt.Errorf("pb UpdateXpubStateReply type assertion error")
	}
	return &endpoint1.AssignedXpubToGroupResponse{}, nil
}

// encodeGetChainsRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetChains request to a gRPC request.
func encodeGetChainsRequest(_ context.Context, request interface{}) (interface{}, error) {
	_, ok := request.(endpoint1.GetChainsRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint GetChainsRequest type assertion error")
	}
	return &pb.GetChainsRequest{}, nil
}

// decodeGetChainsResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetChainsResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.GetChainsReply)
	if !ok {
		return nil, fmt.Errorf("pb GetChainsReply type assertion error")
	}

	chainsResp := make([]*repository.SimpleChain, len(resp.Chains))
	for i, c := range resp.Chains {
		chainsResp[i] = &repository.SimpleChain{
			ID:      c.Id,
			Name:    c.Name,
			Coin:    c.Coin,
			Bip44id: uint(c.Bip44Id),
			Status:  c.Status,
			Decimal: c.Decimal}
	}

	return endpoint1.GetChainsResponse{Chains: chainsResp}, nil
}

// encodeCreateWalletRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain CreateWallet request to a gRPC request.
func encodeCreateWalletRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.CreateWalletRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint CreateWalletRequest type assertion error")
	}
	return &pb.CreateWalletRequest{Groupid: r.Groupid, Bip44Change: int32(r.Bip44change), Chainname: r.Chainname}, nil
}

// decodeCreateWalletResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCreateWalletResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.CreateWalletReply)
	if !ok {
		return nil, fmt.Errorf("pb CreateWalletReply type assertion error")
	}
	wallet := repository.Wallet{}
	if err := copier.Copy(&wallet, resp.Wallet); err != nil {
		return nil, err
	}
	return endpoint1.CreateWalletResponse{Wallet: &wallet}, nil
}

// encodeGetWalletsRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetWallets request to a gRPC request.
func encodeGetWalletsRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.GetWalletsRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint GetWalletsRequest type assertion error")
	}
	return &pb.GetWalletsRequest{Groupid: r.Groupid, Page: r.Page, Limit: r.Limit, Bip44Change: r.Bip44Change}, nil
}

// decodeGetWalletsResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetWalletsResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.GetWalletsReply)
	if !ok {
		e := fmt.Errorf("pb GetWalletsReply type assertion error")
		return endpoint1.GetWalletsResponse{Err: e}, e
	}
	wallets := []*repository.ChainWithWallets{}
	if err := copier.Copy(&wallets, resp.ChainWithWallets); err != nil {
		return endpoint1.GetWalletsResponse{Err: err}, err
	}
	return endpoint1.GetWalletsResponse{ChainWithWallets: wallets}, nil
}

// encodeQueryWalletsForGroupByChainNameRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain QueryWalletsForGroupByChainName request to a gRPC request.
func encodeQueryWalletsForGroupByChainNameRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.QueryWalletsForGroupByChainNameRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint QueryWalletsForGroupByChainNameRequest type assertion error")
	}
	return &pb.QueryWalletsForGroupByChainNameRequest{Groupid: r.Groupid, ChainName: r.ChainName}, nil
}

// decodeQueryWalletsForGroupByChainNameResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeQueryWalletsForGroupByChainNameResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.QueryWalletsForGroupByChainNameReply)
	if !ok {
		return nil, fmt.Errorf("pb QueryWalletsForGroupByChainNameReply type assertion error")
	}
	wallets := []*repository.Wallet{}
	if err := copier.Copy(&wallets, resp.Wallets); err != nil {
		return endpoint1.QueryWalletsForGroupByChainNameResponse{Err: err}, err
	}
	return endpoint1.QueryWalletsForGroupByChainNameResponse{Wallets: wallets}, nil
}

// encodeQueryWalletHDRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain QueryWalletHD request to a gRPC request.
func encodeQueryWalletHDRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.QueryWalletHDRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint QueryWalletHDRequest type assertion error")
	}
	return &pb.QueryWalletHDRequest{Address: r.Address}, nil
}

// decodeQueryWalletHDResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeQueryWalletHDResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, ok := reply.(*pb.QueryWalletHDReply)
	if !ok {
		e := fmt.Errorf("pb QueryWalletHDReply type assertion error")
		return endpoint1.QueryWalletHDResponse{Err: e}, e
	}
	walletHD := repository.WalletHD{}
	if err := copier.Copy(&walletHD, resp.WalletHD); err != nil {
		return endpoint1.QueryWalletHDResponse{Err: err}, err
	}
	return endpoint1.QueryWalletHDResponse{Hd: &walletHD}, nil
}
