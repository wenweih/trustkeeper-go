package client

import (
	endpoint1 "trustkeeper-go/app/service/account/pkg/endpoint"
	pb "trustkeeper-go/app/service/account/pkg/grpc/pb"
	service "trustkeeper-go/app/service/account/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"

	"context"
	"errors"
	"fmt"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
// 查阅 https://github.com/kujtimiihoxha/kit#generate-the-client-library
func newGRPCClient(conn *grpc.ClientConn, options []grpc1.ClientOption) (service.AccountService, error) {
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

	var userInfoEndpoint endpoint.Endpoint
	{
		userInfoEndpoint = grpc1.NewClient(conn, "pb.Account", "UserInfo", encodeUserInfoRequest, decodeUserInfoResponse, pb.UserInfoReply{}, options...).Endpoint()
	}

	var authEndpoint endpoint.Endpoint
	{
		authEndpoint = grpc1.NewClient(conn, "pb.Account", "Auth", encodeAuthRequest, decodeAuthResponse, pb.AuthReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
		AuthEndpoint:     authEndpoint,
		CreateEndpoint:   createEndpoint,
		RolesEndpoint:    rolesEndpoint,
		SigninEndpoint:   signinEndpoint,
		SignoutEndpoint:  signoutEndpoint,
		UserInfoEndpoint: userInfoEndpoint,
	}, nil
}

// encodeCreateRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Create request to a gRPC request.
func encodeCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	r := request.(endpoint1.CreateRequest)
	return &pb.CreateRequest{
		Email:    r.Email,
		Password: r.Password,
		Orgname:  r.OrgName}, nil
}

// decodeCreateResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCreateResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp, found := reply.(*pb.CreateReply)
	if !found {
		return nil, fmt.Errorf("pb CreateReply type assertion error")
	}
	return endpoint1.CreateResponse{E1: nil, UUID: resp.Uuid}, nil
}

// encodeSigninRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Signin request to a gRPC request.
func encodeSigninRequest(_ context.Context, request interface{}) (interface{}, error) {
	r := request.(endpoint1.SigninRequest)
	return &pb.SigninRequest{
		Email:    r.Email,
		Password: r.Password}, nil
}

// decodeSigninResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeSigninResponse(_ context.Context, reply interface{}) (interface{}, error) {
	r, found := reply.(*pb.SigninReply)
	if !found {
		e := fmt.Errorf("pb CreateReply type assertion error")
		return &endpoint1.SigninResponse{
			E1: e,
		}, e
	}
	return endpoint1.SigninResponse{S0: r.Token}, nil
}

// encodeSignoutRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Signout request to a gRPC request.
func encodeSignoutRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

// decodeSignoutResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeSignoutResponse(_ context.Context, reply interface{}) (interface{}, error) {
	r := reply.(*pb.SignoutReply)
	if r.Result {
		return &endpoint1.SignoutResponse{E0: nil}, nil
	}
	return nil, errors.New("'Account' Decoder is not impelemented")
}

// encodeRolesRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Roles request to a gRPC request.
func encodeRolesRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

// decodeRolesResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeRolesResponse(_ context.Context, reply interface{}) (interface{}, error) {
	r, ok := reply.(*pb.RolesReply)
	if !ok {
		e := errors.New("'Account' Decoder is not impelemented")
		return endpoint1.RolesResponse{E1: e}, e
	}
	return endpoint1.RolesResponse{S0: r.Roles}, nil
}

// encodeUserInfoRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain UserInfo request to a gRPC request.
func encodeUserInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &pb.UserInfoRequest{}, nil
	// return nil, errors.New("'Account' Encoder is not impelemented")
}

// decodeUserInfoResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeUserInfoResponse(_ context.Context, reply interface{}) (interface{}, error) {
	r, ok := reply.(*pb.UserInfoReply)
	if !ok {
		e := errors.New("pb UserInfoReply type assertion error")
		return nil, e
	}
	return endpoint1.UserInfoResponse{Roles: r.Roles, OrgName: r.OrgName}, nil
}

// encodeAuthRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Auth request to a gRPC request.
func encodeAuthRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &pb.AuthRequest{}, nil
}

// decodeAuthResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeAuthResponse(_ context.Context, reply interface{}) (interface{}, error) {
	r, ok := reply.(*pb.AuthReply)
	if !ok {
		e := errors.New("'AuthReply' Decoder is not impelemented")
		return endpoint1.AuthResponse{Err: e}, e
	}
	return endpoint1.AuthResponse{Uuid: r.Uuid, NamespaceID: r.NamespaceID, Roles: r.Roles}, nil
}
