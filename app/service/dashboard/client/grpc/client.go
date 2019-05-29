package grpc

import (
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	endpoint1 "trustkeeper-go/app/service/dashboard/pkg/endpoint"
	pb "trustkeeper-go/app/service/dashboard/pkg/grpc/pb"
	service "trustkeeper-go/app/service/dashboard/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.DashboardService, error) {
	var createGroupEndpoint endpoint.Endpoint
	{
		createGroupEndpoint = grpc1.NewClient(conn, "pb.Dashboard", "CreateGroup", encodeCreateGroupRequest, decodeCreateGroupResponse, pb.CreateGroupReply{}, options["CreateGroup"]...).Endpoint()
	}

	var getGroupsEndpoint endpoint.Endpoint
	{
		getGroupsEndpoint = grpc1.NewClient(conn, "pb.Dashboard", "GetGroups", encodeGetGroupsRequest, decodeGetGroupsResponse, pb.GetGroupsReply{}, options["GetGroups"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		CreateGroupEndpoint: createGroupEndpoint,
		GetGroupsEndpoint:   getGroupsEndpoint,
	}, nil
}
