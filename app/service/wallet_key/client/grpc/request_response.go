package grpc

import (
  "errors"
  "context"
  pb "trustkeeper-go/app/service/wallet_key/pkg/grpc/pb"
  "trustkeeper-go/app/service/wallet_key/pkg/endpoint"
)

// encodeGenerateMnemonicRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GenerateMnemonic request to a gRPC request.
func encodeGenerateMnemonicRequest(_ context.Context, request interface{}) (interface{}, error) {
  r := request.(endpoint.GenerateMnemonicRequest)
  return &pb.GenerateMnemonicRequest{Uuid: r.Uuid}, nil
}

// decodeGenerateMnemonicResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGenerateMnemonicResponse(_ context.Context, reply interface{}) (interface{}, error) {
  r, found := reply.(*pb.GenerateMnemonicReply)
  if !found{
    e := errors.New("'GenerateMnemonicReply' Decoder is not impelemented")
    return &endpoint.GenerateMnemonicResponse{Err: e}, e
  }
  return &endpoint.GenerateMnemonicResponse{Xpub: r.Xpub}, nil
}
