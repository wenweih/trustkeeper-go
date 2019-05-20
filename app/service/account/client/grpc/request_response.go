package grpc

import (
  "fmt"
  "context"
  "errors"
  pb "trustkeeper-go/app/service/account/pkg/grpc/pb"
  "trustkeeper-go/app/service/account/pkg/endpoint"
)

// encodeCreateRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Create request to a gRPC request.
func encodeCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	r := request.(endpoint.CreateRequest)
	return &pb.CreateRequest{
		Email: r.Email,
		Password: r.Password}, nil
}

// decodeCreateResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCreateResponse(_ context.Context, reply interface{}) (interface{}, error) {
  _, found := reply.(*pb.CreateReply)
  if !found{
    return nil, fmt.Errorf("pb CreateReply type assertion error")
  }
  return &endpoint.CreateResponse{E1: nil}, nil
}

// encodeSigninRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Signin request to a gRPC request.
func encodeSigninRequest(_ context.Context, request interface{}) (interface{}, error) {
  r := request.(endpoint.SigninRequest)
  return &pb.SigninRequest{
    Email: r.Email,
    Password: r.Password}, nil
}

// decodeSigninResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeSigninResponse(_ context.Context, reply interface{}) (interface{}, error) {
  r, found := reply.(*pb.SigninReply)
  if !found {
    e := fmt.Errorf("pb CreateReply type assertion error")
    return &endpoint.SigninResponse{
      E1: e,
    }, e
  }
  return &endpoint.SigninResponse{S0: r.Token,}, nil
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
  if r.Result{
    return &endpoint.SignoutResponse{E0: nil}, nil
  }
	return nil, errors.New("'Account' Decoder is not impelemented")
}

// encodeRolesRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Roles request to a gRPC request.
func encodeRolesRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Account' Encoder is not impelemented")
}

// decodeRolesResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeRolesResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Account' Decoder is not impelemented")
}

// encodeFindByTokenIDRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain FindByTokenID request to a gRPC request.
func encodeFindByTokenIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Account' Encoder is not impelemented")
}

// decodeFindByTokenIDResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeFindByTokenIDResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Account' Decoder is not impelemented")
}
