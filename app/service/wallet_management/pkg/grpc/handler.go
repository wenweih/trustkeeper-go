package grpc

import (
	"context"
	"fmt"
	endpoint "trustkeeper-go/app/service/wallet_management/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_management/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/jinzhu/copier"
	context1 "golang.org/x/net/context"
)

func makeCreateChainHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateChainEndpoint, decodeCreateChainRequest, encodeCreateChainResponse, options...)
}

func decodeCreateChainRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateChainRequest)
	return endpoint.CreateChainRequest{Symbol: req.Symbol, Bit44ID: req.Bitid, Status: req.Status}, nil
}

func encodeCreateChainResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.CreateChainResponse)
	if resp.Err != nil {
		return &pb.CreateChainReply{Result: false}, resp.Err
	}
	return &pb.CreateChainReply{Result: true}, nil
}

func (g *grpcServer) CreateChain(ctx context1.Context, req *pb.CreateChainRequest) (*pb.CreateChainReply, error) {
	_, rep, err := g.createChain.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateChainReply), nil
}

func makeAssignedXpubToGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AssignedXpubToGroupEndpoint, decodeAssignedXpubToGroupRequest, encodeAssignedXpubToGroupResponse, options...)
}

func decodeAssignedXpubToGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.AssignedXpubToGroupRequest)
	if !ok {
		return nil, fmt.Errorf("pb AssignedXpubToGroupRequest type assersion error")
	}
	return endpoint.AssignedXpubToGroupRequest{GroupID: req.Groupid}, nil
}

func encodeAssignedXpubToGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	_, ok := r.(endpoint.AssignedXpubToGroupResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint AssignedXpubToGroupResponse type assersion error")
	}
	return &pb.AssignedXpubToGroupReply{}, nil
}
func (g *grpcServer) AssignedXpubToGroup(ctx context1.Context, req *pb.AssignedXpubToGroupRequest) (*pb.AssignedXpubToGroupReply, error) {
	_, rep, err := g.assignedXpubToGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AssignedXpubToGroupReply), nil
}

func makeGetChainsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetChainsEndpoint, decodeGetChainsRequest, encodeGetChainsResponse, options...)
}

func decodeGetChainsRequest(_ context.Context, r interface{}) (interface{}, error) {
	_, ok := r.(*pb.GetChainsRequest)
	if !ok {
		return nil, fmt.Errorf("interface{} to pb GetChainsRequest type assertion error")
	}
	return endpoint.GetChainsRequest{}, nil
}

func encodeGetChainsResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.GetChainsResponse)
	if !ok {
		return nil, fmt.Errorf("interface{} to endpoint GetChainsResponse type assertion error")
	}

	pbChains := make([]*pb.SimpleChain, len(resp.Chains))
	for i, c := range resp.Chains {
		pbChains[i] = &pb.SimpleChain{
			Id:      c.ID,
			Name:    c.Name,
			Coin:    c.Coin,
			Bip44Id: int32(c.Bip44id),
			Status:  c.Status}
	}
	return &pb.GetChainsReply{Chains: pbChains}, nil
}
func (g *grpcServer) GetChains(ctx context1.Context, req *pb.GetChainsRequest) (*pb.GetChainsReply, error) {
	_, rep, err := g.getChains.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetChainsReply), nil
}

func makeCreateWalletHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateWalletEndpoint, decodeCreateWalletRequest, encodeCreateWalletResponse, options...)
}

func decodeCreateWalletRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.CreateWalletRequest)
	if !ok {
		return nil, fmt.Errorf("pb CreateWalletRequest type assersion error")
	}
	return endpoint.CreateWalletRequest{Groupid: req.Groupid, Chainname: req.Chainname, Bip44change: int(req.Bip44Change)}, nil
}

func encodeCreateWalletResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.CreateWalletResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint CreateWalletResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	wallet := pb.Wallet{}
	if err := copier.Copy(&wallet, resp.Wallet); err != nil {
		return nil, err
	}
	return &pb.CreateWalletReply{Wallet: &wallet}, nil
}

func (g *grpcServer) CreateWallet(ctx context1.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletReply, error) {
	_, rep, err := g.createWallet.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateWalletReply), nil
}

func makeGetWalletsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetWalletsEndpoint, decodeGetWalletsRequest, encodeGetWalletsResponse, options...)
}

func decodeGetWalletsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.GetWalletsRequest)
	if !ok {
		return nil, fmt.Errorf("pb GetWalletsRequest type assersion error")
	}
	return endpoint.GetWalletsRequest{Groupid: req.Groupid}, nil
}

func encodeGetWalletsResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.GetWalletsResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint GetWalletsResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	wallets :=[]* pb.Wallet{}
	if err := copier.Copy(&wallets, resp.Wallets); err != nil {
		return nil, err
	}
	return &pb.GetWalletsReply{Wallets: wallets}, nil
}
func (g *grpcServer) GetWallets(ctx context1.Context, req *pb.GetWalletsRequest) (*pb.GetWalletsReply, error) {
	_, rep, err := g.getWallets.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetWalletsReply), nil
}
