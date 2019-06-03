package grpc

import (
	"context"
	endpoint "trustkeeper-go/app/service/account/pkg/endpoint"
	pb "trustkeeper-go/app/service/account/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

func makeCreateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)
}

func decodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateRequest)
	return endpoint.CreateRequest{Email: req.Email, Password: req.Password}, nil
}

func encodeCreateResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.CreateResponse)
	if resp.E1 != nil {
		return &pb.CreateReply{Uuid: resp.UUID}, resp.E1
	}
	return &pb.CreateReply{Uuid: resp.UUID}, nil
}
func (g *grpcServer) Create(ctx context1.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	_, rep, err := g.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateReply), nil
}

func makeSigninHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SigninEndpoint, decodeSigninRequest, encodeSigninResponse, options...)
}

func decodeSigninRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SigninRequest)
	return endpoint.SigninRequest{Email: req.Email, Password: req.Password}, nil
}

func encodeSigninResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.SigninResponse)
	if resp.E1 != nil {
		return &pb.SigninReply{Token: resp.S0}, resp.E1
	}
	return &pb.SigninReply{Token: resp.S0}, nil
}
func (g *grpcServer) Signin(ctx context1.Context, req *pb.SigninRequest) (*pb.SigninReply, error) {
	_, rep, err := g.signin.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SigninReply), nil
}

func makeSignoutHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SignoutEndpoint, decodeSignoutRequest, encodeSignoutResponse, options...)
}

func decodeSignoutRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SignoutRequest)
	return endpoint.SignoutRequest{Token: req.Token}, nil
}

func encodeSignoutResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.SignoutResponse)
	if resp.E0 != nil {
		return nil, resp.E0
	}
	return &pb.SignoutReply{Result: true}, nil
}
func (g *grpcServer) Signout(ctx context1.Context, req *pb.SignoutRequest) (*pb.SignoutReply, error) {
	_, rep, err := g.signout.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SignoutReply), nil
}

func makeRolesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RolesEndpoint, decodeRolesRequest, encodeRolesResponse, options...)
}

func decodeRolesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return endpoint.RolesRequest{}, nil
}

func encodeRolesResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.RolesResponse)
	if resp.E1 != nil {
		return nil, resp.E1
	}
	return &pb.RolesReply{Roles: resp.S0}, nil
}
func (g *grpcServer) Roles(ctx context1.Context, req *pb.RolesRequest) (*pb.RolesReply, error) {
	_, rep, err := g.roles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RolesReply), nil
}

func makeAuthHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AuthEndpoint, decodeAuthRequest, encodeAuthResponse, options...)
}

func decodeAuthRequest(_ context.Context, r interface{}) (interface{}, error) {
	return endpoint.AuthRequest{}, nil
}

func encodeAuthResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.AuthResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.AuthReply{Uuid: resp.Uuid}, nil
}

func (g *grpcServer) Auth(ctx context1.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	_, rep, err := g.auth.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AuthReply), nil
}
