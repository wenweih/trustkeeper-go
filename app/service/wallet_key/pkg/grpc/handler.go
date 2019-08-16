package grpc

import (
	"context"
	"fmt"
	endpoint "trustkeeper-go/app/service/wallet_key/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_key/pkg/grpc/pb"
	"github.com/jinzhu/copier"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	"trustkeeper-go/app/service/wallet_key/pkg/repository"
)

func makeGenerateMnemonicHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GenerateMnemonicEndpoint, decodeGenerateMnemonicRequest, encodeGenerateMnemonicResponse, options...)
}

func decodeGenerateMnemonicRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GenerateMnemonicRequest)
	return endpoint.GenerateMnemonicRequest{Namespaceid: req.Namespaceid, Bip44ids: req.Bip44Ids, Bip44accountSize: int(req.Bip44AccountSize)}, nil
}

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
	return &pb.GenerateMnemonicReply{Chainsxpubs: pbXpubs, Version: resp.Version}, nil
}

func (g *grpcServer) GenerateMnemonic(ctx context1.Context, req *pb.GenerateMnemonicRequest) (*pb.GenerateMnemonicReply, error) {
	_, rep, err := g.generateMnemonic.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GenerateMnemonicReply), nil
}

func makeSignedBitcoincoreTxHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SignedBitcoincoreTxEndpoint, decodeSignedBitcoincoreTxRequest, encodeSignedBitcoincoreTxResponse, options...)
}

func decodeSignedBitcoincoreTxRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.SignedBitcoincoreTxRequest)
	if !ok {
		return nil, fmt.Errorf("pb SignedBitcoincoreTxRequest type assersion error")
	}
	walletHD := repository.WalletHD{}
	if err := copier.Copy(&walletHD, req.WalletHD); err != nil {
		return nil, err
	}
	return endpoint.SignedBitcoincoreTxRequest{WalletHD: walletHD, VinAmount: req.VinAmount, TxHex: req.TxHex}, nil
}

func encodeSignedBitcoincoreTxResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.SignedBitcoincoreTxResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint SignedBitcoincoreTxResponse type assertion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.SignedBitcoincoreTxReply{SignedTxHex: resp.SignedTxHex}, nil
}
func (g *grpcServer) SignedBitcoincoreTx(ctx context1.Context, req *pb.SignedBitcoincoreTxRequest) (*pb.SignedBitcoincoreTxReply, error) {
	_, rep, err := g.signedBitcoincoreTx.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SignedBitcoincoreTxReply), nil
}
