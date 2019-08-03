// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "trustkeeper-go/app/service/transaction/pkg/endpoint"
	pb "trustkeeper-go/app/service/transaction/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	assignAssetsToWallet   grpc.Handler
	createBalancesForAsset grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.TransactionServer {
	return &grpcServer{
		assignAssetsToWallet:   makeAssignAssetsToWalletHandler(endpoints, options["AssignAssetsToWallet"]),
		createBalancesForAsset: makeCreateBalancesForAssetHandler(endpoints, options["CreateBalancesForAsset"]),
	}
}
