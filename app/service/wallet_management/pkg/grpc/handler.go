package grpc

import (
	"fmt"
	"context"
	endpoint "trustkeeper-go/app/service/wallet_management/pkg/endpoint"
	pb "trustkeeper-go/app/service/wallet_management/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
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
