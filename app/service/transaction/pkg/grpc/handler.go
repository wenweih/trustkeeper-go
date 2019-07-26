package grpc

import (
	"fmt"
	"context"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "trustkeeper-go/app/service/transaction/pkg/endpoint"
	pb "trustkeeper-go/app/service/transaction/pkg/grpc/pb"

	"github.com/jinzhu/copier"
	"trustkeeper-go/app/service/transaction/pkg/repository"
)

// makeAssignAssetsToWalletHandler creates the handler logic
func makeAssignAssetsToWalletHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AssignAssetsToWalletEndpoint, decodeAssignAssetsToWalletRequest, encodeAssignAssetsToWalletResponse, options...)
}

// decodeAssignAssetsToWalletResponse is a transport/grpc.DecodeRequestFunc that converts a
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

// encodeAssignAssetsToWalletResponse is a transport/grpc.EncodeResponseFunc that converts
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
