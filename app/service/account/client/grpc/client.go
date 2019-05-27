package grpc

import (
	grpc1 "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "google.golang.org/grpc"
	endpoint1 "trustkeeper-go/app/service/account/pkg/endpoint"
	pb "trustkeeper-go/app/service/account/pkg/grpc/pb"
	service "trustkeeper-go/app/service/account/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
// 查阅 https://github.com/kujtimiihoxha/kit#generate-the-client-library
func New(conn *grpc.ClientConn, options []grpc1.ClientOption) (service.AccountService, error) {
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = grpc1.NewClient(conn, "pb.Account", "Create", encodeCreateRequest, decodeCreateResponse, pb.CreateReply{}, options...).Endpoint()
	}

	var signinEndpoint endpoint.Endpoint
	{
		signinEndpoint = grpc1.NewClient(conn, "pb.Account", "Signin", encodeSigninRequest, decodeSigninResponse, pb.SigninReply{}, options...).Endpoint()
	}

	var signoutEndpoint endpoint.Endpoint
	{
		signoutEndpoint = grpc1.NewClient(conn, "pb.Account", "Signout", encodeSignoutRequest, decodeSignoutResponse, pb.SignoutReply{}, options...).Endpoint()
	}

	var rolesEndpoint endpoint.Endpoint
	{
		rolesEndpoint = grpc1.NewClient(conn, "pb.Account", "Roles", encodeRolesRequest, decodeRolesResponse, pb.RolesReply{}, options...).Endpoint()
	}

	var authEndpoint endpoint.Endpoint
	{
		authEndpoint = grpc1.NewClient(conn, "pb.Account", "Auth", encodeAuthRequest, decodeAuthResponse, pb.AuthReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
		AuthEndpoint:    authEndpoint,
		CreateEndpoint:        createEndpoint,
		RolesEndpoint:         rolesEndpoint,
		SigninEndpoint:        signinEndpoint,
		SignoutEndpoint:       signoutEndpoint,
	}, nil
}
