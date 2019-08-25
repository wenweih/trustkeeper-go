package client

import (
  "fmt"
  "errors"
  "context"
  pb "trustkeeper-go/app/service/wallet_key/pkg/grpc/pb"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
  "github.com/jinzhu/copier"
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
		generateMnemonicEndpoint = grpc1.NewClient(conn, "pb.WalletKey", "GenerateMnemonic",
      encodeGenerateMnemonicRequest, decodeGenerateMnemonicResponse, pb.GenerateMnemonicReply{}, options...).Endpoint()
	}

  var signedBitcoincoreTxEndpoint endpoint.Endpoint
	{
		signedBitcoincoreTxEndpoint = grpc1.NewClient(conn, "pb.WalletKey", "SignedBitcoincoreTx",
      encodeSignedBitcoincoreTxRequest, decodeSignedBitcoincoreTxResponse, pb.SignedBitcoincoreTxReply{}, options...).Endpoint()
	}

  var signedEthereumTxEndpoint endpoint.Endpoint
	{
		signedEthereumTxEndpoint = grpc1.NewClient(conn, "pb.WalletKey", "SignedEthereumTx",
      encodeSignedEthereumTxRequest, decodeSignedEthereumTxResponse, pb.SignedEthereumTxReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
    GenerateMnemonicEndpoint: generateMnemonicEndpoint,
    SignedBitcoincoreTxEndpoint: signedBitcoincoreTxEndpoint,
    SignedEthereumTxEndpoint:    signedEthereumTxEndpoint,
  }, nil
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
    xpubs = append(xpubs, &service.Bip44ThirdXpubsForChain{Chain: uint(chainwithxpubs.Chain), Xpubs: bip44AccountKeys})
  }
  return endpoint1.GenerateMnemonicResponse{ChainsXpubs: xpubs, Version: r.Version}, nil
}

// encodeSignedBitcoincoreTxRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain SignedBitcoincoreTx request to a gRPC request.
func encodeSignedBitcoincoreTxRequest(_ context.Context, request interface{}) (interface{}, error) {
  r, ok := request.(endpoint1.SignedBitcoincoreTxRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint SignedBitcoincoreTxRequest type assertion error")
	}
  walletHD := pb.WalletHD{}
	if err := copier.Copy(&walletHD, r.WalletHD); err != nil {
		return nil, err
	}
  return &pb.SignedBitcoincoreTxRequest{TxHex: r.TxHex, VinAmount: r.VinAmount, WalletHD: &walletHD}, nil
}

// decodeSignedBitcoincoreTxResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeSignedBitcoincoreTxResponse(_ context.Context, reply interface{}) (interface{}, error) {
  resp, ok := reply.(*pb.SignedBitcoincoreTxReply)
	if !ok {
		e := fmt.Errorf("pb SignedBitcoincoreTxReply type assertion error")
    return endpoint1.SignedBitcoincoreTxResponse{Err: e}, e
	}
  return endpoint1.SignedBitcoincoreTxResponse{SignedTxHex: resp.SignedTxHex}, nil
}

// encodeSignedEthereumTxRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain SignedEthereumTx request to a gRPC request.
func encodeSignedEthereumTxRequest(_ context.Context, request interface{}) (interface{}, error) {
  r, ok := request.(endpoint1.SignedEthereumTxRequest)
	if !ok {
		return nil, fmt.Errorf("endpoint SignedEthereumTxRequest type assertion error")
	}
  walletHD := pb.WalletHD{}
	if err := copier.Copy(&walletHD, r.WalletHD); err != nil {
		return nil, err
	}
  return &pb.SignedEthereumTxRequest{ChainID: r.ChainID, TxHex: r.TxHex, WalletHD: &walletHD}, nil
}

// decodeSignedEthereumTxResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeSignedEthereumTxResponse(_ context.Context, reply interface{}) (interface{}, error) {
  resp, ok := reply.(*pb.SignedEthereumTxReply)
	if !ok {
		e := fmt.Errorf("pb SignedEthereumTxReply type assertion error")
    return endpoint1.SignedEthereumTxResponse{Err: e}, e
	}
  return endpoint1.SignedEthereumTxResponse{SignedTxHex: resp.SignedTxHex}, nil
}
