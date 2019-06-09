package grpc

import (
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	endpoint1 "trustkeeper-go/app/service/wallet_key/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_key/pkg/grpc/pb"
	service "trustkeeper-go/app/service/wallet_key/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options []grpc1.ClientOption) (service.WalletKeyService, error) {
	var generateMnemonicEndpoint endpoint.Endpoint
	{
		generateMnemonicEndpoint = grpc1.NewClient(conn, "pb.WalletKey", "GenerateMnemonic", encodeGenerateMnemonicRequest, decodeGenerateMnemonicResponse, pb.GenerateMnemonicReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{GenerateMnemonicEndpoint: generateMnemonicEndpoint}, nil
}
