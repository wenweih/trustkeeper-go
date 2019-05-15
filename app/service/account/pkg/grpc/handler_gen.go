// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "trustkeeper-go/app/service/account/pkg/endpoint"
	pb "trustkeeper-go/app/service/account/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	create  grpc.Handler
	signin  grpc.Handler
	signout grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.AccountServer {
	return &grpcServer{
		create:  makeCreateHandler(endpoints, options["Create"]),
		signin:  makeSigninHandler(endpoints, options["Signin"]),
		signout: makeSignoutHandler(endpoints, options["Signout"]),
	}
}
