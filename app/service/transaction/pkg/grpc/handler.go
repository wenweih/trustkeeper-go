package grpc

import (
	"context"
	"fmt"
	endpoint "trustkeeper-go/app/service/transaction/pkg/endpoint"
	pb "trustkeeper-go/app/service/transaction/pkg/grpc/pb"
	"trustkeeper-go/app/service/transaction/pkg/repository"

	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/jinzhu/copier"
	context1 "golang.org/x/net/context"
)

func makeAssignAssetsToWalletHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AssignAssetsToWalletEndpoint, decodeAssignAssetsToWalletRequest, encodeAssignAssetsToWalletResponse, options...)
}

func decodeAssignAssetsToWalletRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.AssignAssetsToWalletRequest)
	if !ok {
		return nil, fmt.Errorf("interface{} to pb AssignAssetsToWalletRequest type assertion error")
	}
	simpleAssets := []*repository.SimpleAsset{}
	if err := copier.Copy(&simpleAssets, req.SimpleAssets); err != nil {
		return nil, err
	}
	return endpoint.AssignAssetsToWalletRequest{Address: req.Address, Assets: simpleAssets}, nil
}

func encodeAssignAssetsToWalletResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.AssignAssetsToWalletResponse)
	if !ok {
		return nil, fmt.Errorf("interface{} to endpoint AssignAssetsToWalletResponse type assertion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.AssignAssetsToWalletReply{}, nil
}
func (g *grpcServer) AssignAssetsToWallet(ctx context1.Context, req *pb.AssignAssetsToWalletRequest) (*pb.AssignAssetsToWalletReply, error) {
	_, rep, err := g.assignAssetsToWallet.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AssignAssetsToWalletReply), nil
}

func makeCreateBalancesForAssetHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateBalancesForAssetEndpoint, decodeCreateBalancesForAssetRequest, encodeCreateBalancesForAssetResponse, options...)
}

func decodeCreateBalancesForAssetRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.CreateBalancesForAssetRequest)
	if !ok {
		return nil, fmt.Errorf("interface{} to pb CreateBalancesForAssetRequest type assertion error")
	}
	simpleAsset := repository.SimpleAsset{}
	if err := copier.Copy(&simpleAsset, req.Asset); err != nil {
		return nil, err
	}

	wallets := make([]*repository.Wallet, 0)
	if err := copier.Copy(&wallets, req.Wallets); err != nil {
		return nil, err
	}
	return endpoint.CreateBalancesForAssetRequest{Wallets: wallets, Asset: &simpleAsset}, nil
}

func encodeCreateBalancesForAssetResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.CreateBalancesForAssetResponse)
	if !ok {
		return nil, fmt.Errorf("interface{} to endpoint CreateBalancesForAssetResponse type assertion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.CreateBalancesForAssetReply{}, nil
}

func (g *grpcServer) CreateBalancesForAsset(ctx context1.Context, req *pb.CreateBalancesForAssetRequest) (*pb.CreateBalancesForAssetReply, error) {
	_, rep, err := g.createBalancesForAsset.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateBalancesForAssetReply), nil
}
