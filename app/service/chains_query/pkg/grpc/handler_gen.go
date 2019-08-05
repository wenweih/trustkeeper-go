// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "trustkeeper-go/app/service/chains_query/pkg/endpoint"
	pb "trustkeeper-go/app/service/chains_query/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	bitcoincoreBlock  grpc.Handler
	queryOmniProperty grpc.Handler
	eRC20TokenInfo    grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.ChainsQueryServer {
	return &grpcServer{
		bitcoincoreBlock:  makeBitcoincoreBlockHandler(endpoints, options["BitcoincoreBlock"]),
		eRC20TokenInfo:    makeERC20TokenInfoHandler(endpoints, options["ERC20TokenInfo"]),
		queryOmniProperty: makeQueryOmniPropertyHandler(endpoints, options["QueryOmniProperty"]),
	}
}
