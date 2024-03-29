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
			Status:  c.Status,
			Decimal: c.Decimal}
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
	return endpoint.GetWalletsRequest{Groupid: req.Groupid, Page: req.Page, Limit: req.Limit, Bip44Change: req.Bip44Change}, nil
}

func encodeGetWalletsResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.GetWalletsResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint GetWalletsResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	wallets := []*pb.ChainWithWallets{}
	if err := copier.Copy(&wallets, resp.ChainWithWallets); err != nil {
		return nil, err
	}
	return &pb.GetWalletsReply{ChainWithWallets: wallets}, nil
}
func (g *grpcServer) GetWallets(ctx context1.Context, req *pb.GetWalletsRequest) (*pb.GetWalletsReply, error) {
	_, rep, err := g.getWallets.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetWalletsReply), nil
}

func makeQueryWalletsForGroupByChainNameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.QueryWalletsForGroupByChainNameEndpoint, decodeQueryWalletsForGroupByChainNameRequest, encodeQueryWalletsForGroupByChainNameResponse, options...)
}

func decodeQueryWalletsForGroupByChainNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.QueryWalletsForGroupByChainNameRequest)
	if !ok {
		return nil, fmt.Errorf("pb QueryWalletsForGroupByChainNameRequest type assersion error")
	}
	return endpoint.QueryWalletsForGroupByChainNameRequest{Groupid: req.Groupid, ChainName: req.ChainName}, nil
}

func encodeQueryWalletsForGroupByChainNameResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.QueryWalletsForGroupByChainNameResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint QueryWalletsForGroupByChainNameResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	pbWallets := make([]*pb.Wallet, 0)
	if err := copier.Copy(&pbWallets, resp.Wallets); err != nil {
		return nil, err
	}
	return &pb.QueryWalletsForGroupByChainNameReply{Wallets: pbWallets}, nil
}
func (g *grpcServer) QueryWalletsForGroupByChainName(ctx context1.Context, req *pb.QueryWalletsForGroupByChainNameRequest) (*pb.QueryWalletsForGroupByChainNameReply, error) {
	_, rep, err := g.queryWalletsForGroupByChainName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.QueryWalletsForGroupByChainNameReply), nil
}

func makeQueryWalletHDHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.QueryWalletHDEndpoint, decodeQueryWalletHDRequest, encodeQueryWalletHDResponse, options...)
}

func decodeQueryWalletHDRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*pb.QueryWalletHDRequest)
	if !ok {
		return nil, fmt.Errorf("pb QueryWalletHDRequest type assersion error")
	}
	return endpoint.QueryWalletHDRequest{Address: req.Address}, nil
}

func encodeQueryWalletHDResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp, ok := r.(endpoint.QueryWalletHDResponse)
	if !ok {
		return nil, fmt.Errorf("endpoint QueryWalletHDResponse type assersion error")
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	pbWalletHD := pb.WalletHD{}
	if err := copier.Copy(&pbWalletHD, resp.Hd); err != nil {
		return nil, err
	}
	return &pb.QueryWalletHDReply{WalletHD: &pbWalletHD}, nil
}
func (g *grpcServer) QueryWalletHD(ctx context1.Context, req *pb.QueryWalletHDRequest) (*pb.QueryWalletHDReply, error) {
	_, rep, err := g.queryWalletHD.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.QueryWalletHDReply), nil
}
