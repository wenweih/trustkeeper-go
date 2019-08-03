// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "trustkeeper-go/app/service/wallet_management/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_management/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	getChains                       grpc.Handler
	createChain                     grpc.Handler
	assignedXpubToGroup             grpc.Handler
	createWallet                    grpc.Handler
	getWallets                      grpc.Handler
	queryWalletsForGroupByChainName grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.WalletManagementServer {
	return &grpcServer{
		assignedXpubToGroup:             makeAssignedXpubToGroupHandler(endpoints, options["AssignedXpubToGroup"]),
		createChain:                     makeCreateChainHandler(endpoints, options["CreateChain"]),
		createWallet:                    makeCreateWalletHandler(endpoints, options["CreateWallet"]),
		getChains:                       makeGetChainsHandler(endpoints, options["GetChains"]),
		getWallets:                      makeGetWalletsHandler(endpoints, options["GetWallets"]),
		queryWalletsForGroupByChainName: makeQueryWalletsForGroupByChainNameHandler(endpoints, options["QueryWalletsForGroupByChainName"]),
	}
}
