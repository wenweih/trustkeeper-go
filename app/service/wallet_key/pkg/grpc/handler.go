package grpc

import (
	"context"
	endpoint "trustkeeper-go/app/service/wallet_key/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_key/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeGenerateMnemonicHandler creates the handler logic
func makeGenerateMnemonicHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GenerateMnemonicEndpoint, decodeGenerateMnemonicRequest, encodeGenerateMnemonicResponse, options...)
}

// decodeGenerateMnemonicResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GenerateMnemonic request.
// TODO implement the decoder
func decodeGenerateMnemonicRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GenerateMnemonicRequest)
	return endpoint.GenerateMnemonicRequest{Namespaceid: req.Namespaceid, Bip44ids: req.Bip44Ids, Bip44accountSize: int(req.Bip44AccountSize)}, nil
}

// encodeGenerateMnemonicResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeGenerateMnemonicResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GenerateMnemonicResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}
	pbXpubs := []*pb.Bip44ThirdXpubsForChain{}
	for _, chainxpub := range resp.ChainsXpubs {
		pbPubkeys := []*pb.Bip44AccountKey{}
		for _, key := range chainxpub.Xpubs {
			pbPubkeys = append(pbPubkeys, &pb.Bip44AccountKey{Account: int32(key.Account), Key: key.Key})
		}
		pbXpubs = append(pbXpubs, &pb.Bip44ThirdXpubsForChain{Chain: int32(chainxpub.Chain), Xpubs: pbPubkeys})
	}
	return &pb.GenerateMnemonicReply{Chainsxpubs: pbXpubs}, nil
}

func (g *grpcServer) GenerateMnemonic(ctx context1.Context, req *pb.GenerateMnemonicRequest) (*pb.GenerateMnemonicReply, error) {
	_, rep, err := g.generateMnemonic.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GenerateMnemonicReply), nil
}
