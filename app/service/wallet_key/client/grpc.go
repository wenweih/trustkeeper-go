package client

import (
  "errors"
  "context"
  pb "trustkeeper-go/app/service/wallet_key/pkg/grpc/pb"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	endpoint1 "trustkeeper-go/app/service/wallet_key/pkg/endpoint"
	service "trustkeeper-go/app/service/wallet_key/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func newGRPCClient(conn *grpc.ClientConn, options []grpc1.ClientOption) (service.WalletKeyService, error) {
	var generateMnemonicEndpoint endpoint.Endpoint
	{
		generateMnemonicEndpoint = grpc1.NewClient(conn, "pb.WalletKey", "GenerateMnemonic", encodeGenerateMnemonicRequest, decodeGenerateMnemonicResponse, pb.GenerateMnemonicReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{GenerateMnemonicEndpoint: generateMnemonicEndpoint}, nil
}

// encodeGenerateMnemonicRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GenerateMnemonic request to a gRPC request.
func encodeGenerateMnemonicRequest(_ context.Context, request interface{}) (interface{}, error) {
  r := request.(endpoint1.GenerateMnemonicRequest)
  return &pb.GenerateMnemonicRequest{Namespaceid: r.Namespaceid, Bip44Ids: r.Bip44ids, Bip44AccountSize: int32(r.Bip44accountSize)}, nil
}

// decodeGenerateMnemonicResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGenerateMnemonicResponse(_ context.Context, reply interface{}) (interface{}, error) {
  r, found := reply.(*pb.GenerateMnemonicReply)
  if !found{
    e := errors.New("request interface to *pb.GenerateMnemonicReply type assertion error")
    return endpoint1.GenerateMnemonicResponse{Err: e}, e
  }

  xpubs := []*service.Bip44ThirdXpubsForChain{}
  for _, chainwithxpubs := range r.Chainsxpubs {
    bip44AccountKeys := []*service.Bip44AccountKey{}
    for _,  bip44AccountKey := range chainwithxpubs.Xpubs {
      bip44AccountKeys = append(bip44AccountKeys, &service.Bip44AccountKey{
        Account: int(bip44AccountKey.Account), Key: bip44AccountKey.Key})
    }
    xpubs = append(xpubs, &service.Bip44ThirdXpubsForChain{Chain: int(chainwithxpubs.Chain), Xpubs: bip44AccountKeys})
  }
  return endpoint1.GenerateMnemonicResponse{ChainsXpubs: xpubs, Version: r.Version}, nil
}
