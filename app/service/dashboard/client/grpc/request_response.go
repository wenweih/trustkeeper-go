package grpc

import (
  "fmt"
  "context"
	"errors"
  pb "trustkeeper-go/app/service/dashboard/pkg/grpc/pb"
  "trustkeeper-go/app/service/dashboard/pkg/endpoint"
)

// encodeCreateGroupRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain CreateGroup request to a gRPC request.
func encodeCreateGroupRequest(_ context.Context, request interface{}) (interface{}, error) {
  r, ok := request.(endpoint.CreateGroupRequest)
  if !ok {
    return nil, fmt.Errorf("request interface to endpoint.CreateGroupRequest type assertion error")
  }
  return &pb.CreateGroupRequest{Uuid: r.UUID, Name: r.Name, Desc: r.Desc}, nil
}

// decodeCreateGroupResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCreateGroupResponse(_ context.Context, reply interface{}) (interface{}, error) {
  resp, ok := reply.(*pb.CreateGroupReply)
  if !ok{
    return nil, fmt.Errorf("pb CreateReply type assertion error")
  }
  return &endpoint.CreateGroupResponse{Result: resp.Result}, nil
}

// encodeGetGroupsRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetGroups request to a gRPC request.
func encodeGetGroupsRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Dashboard' Encoder is not impelemented")
}

// decodeGetGroupsResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetGroupsResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Dashboard' Decoder is not impelemented")
}
