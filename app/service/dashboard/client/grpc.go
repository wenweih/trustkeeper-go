package client


import (
  "fmt"
  "context"
  "trustkeeper-go/app/service/dashboard/pkg/repository"
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
func newGRPCClient(conn *grpc.ClientConn, options []grpc1.ClientOption) (service.DashboardService, error) {
	var createGroupEndpoint endpoint.Endpoint
	{
		createGroupEndpoint = grpc1.NewClient(conn, "pb.Dashboard", "CreateGroup", encodeCreateGroupRequest, decodeCreateGroupResponse, pb.CreateGroupReply{}, options...).Endpoint()
	}

	var getGroupsEndpoint endpoint.Endpoint
	{
		getGroupsEndpoint = grpc1.NewClient(conn, "pb.Dashboard", "GetGroups", encodeGetGroupsRequest, decodeGetGroupsResponse, pb.GetGroupsReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
		CreateGroupEndpoint: createGroupEndpoint,
		GetGroupsEndpoint:   getGroupsEndpoint,
	}, nil
}

// encodeCreateGroupRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain CreateGroup request to a gRPC request.
func encodeCreateGroupRequest(_ context.Context, request interface{}) (interface{}, error) {
  r, ok := request.(endpoint1.CreateGroupRequest)
  if !ok {
    return nil, fmt.Errorf("request interface to endpoint.CreateGroupRequest type assertion error")
  }
  return &pb.CreateGroupRequest{Uuid: r.UUID, Name: r.Name, Desc: r.Desc, NamespaceID: r.NamespaceID}, nil
}

// decodeCreateGroupResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCreateGroupResponse(_ context.Context, reply interface{}) (interface{}, error) {
  resp, ok := reply.(*pb.CreateGroupReply)
  if !ok{
    return nil, fmt.Errorf("pb CreateReply type assertion error")
  }
  return endpoint1.CreateGroupResponse{Name: resp.Group.Name, Desc: resp.Group.Desc}, nil
}

// encodeGetGroupsRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetGroups request to a gRPC request.
func encodeGetGroupsRequest(_ context.Context, request interface{}) (interface{}, error) {
  r, ok := request.(endpoint1.GetGroupsRequest)
  if !ok {
    return nil, fmt.Errorf("request interface to endpoint.GetGroupsRequest type assertion error")
  }
  return &pb.GetGroupsRequest{NamespaceID: r.NamespaceID}, nil
}

// decodeGetGroupsResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetGroupsResponse(_ context.Context, reply interface{}) (interface{}, error) {
  resp, ok := reply.(*pb.GetGroupsReply)
  if !ok{
    return nil, fmt.Errorf("pb GetGroupsReply type assertion error")
  }

  groupsResp := make([]*repository.GetGroupsResp, len(resp.Groups))
  for i, g := range resp.Groups {
    groupsResp[i] = &repository.GetGroupsResp{Name: g.Name, Desc: g.Desc}
  }

  return endpoint1.GetGroupsResponse{Groups: groupsResp}, nil
}
