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

  var updateGroupEndpoint endpoint.Endpoint
	{
		updateGroupEndpoint = grpc1.NewClient(conn, "pb.Dashboard", "UpdateGroup", encodeUpdateGroupRequest, decodeUpdateGroupResponse, pb.UpdateGroupReply{}, options...).Endpoint()
	}

  var getGroupAssetEndpoint endpoint.Endpoint
	{
		getGroupAssetEndpoint = grpc1.NewClient(conn, "pb.Dashboard", "GetGroupAsset", encodeGetGroupAssetRequest, decodeGetGroupAssetResponse, pb.GetGroupAssetReply{}, options...).Endpoint()
	}

	return endpoint1.Endpoints{
		CreateGroupEndpoint: createGroupEndpoint,
		GetGroupsEndpoint:   getGroupsEndpoint,
    UpdateGroupEndpoint: updateGroupEndpoint,
    GetGroupAssetsEndpoint: getGroupAssetEndpoint,
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
  return endpoint1.CreateGroupResponse{Name: resp.Group.Name, Desc: resp.Group.Desc, ID: resp.Group.Id}, nil
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
    groupsResp[i] = &repository.GetGroupsResp{Name: g.Name, Desc: g.Desc, ID: g.Id}
  }

  return endpoint1.GetGroupsResponse{Groups: groupsResp}, nil
}

// encodeUpdateGroupRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain UpdateGroup request to a gRPC request.
func encodeUpdateGroupRequest(_ context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(endpoint1.UpdateGroupRequest)
  if !ok {
    return nil, fmt.Errorf("endpoint UpdateGroupRequest type assertion error")
  }
  return &pb.UpdateGroupRequest{Groupid: r.GroupID, Name: r.Name, Desc: r.Desc}, nil
}

// decodeUpdateGroupResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeUpdateGroupResponse(_ context.Context, reply interface{}) (interface{}, error) {
  _, ok := reply.(*pb.UpdateGroupReply)
  if !ok {
    return nil, fmt.Errorf("pb UpdateGroupReply type assertion error")
  }
	return endpoint1.UpdateGroupResponse{}, nil
}

// encodeGetGroupAssetRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetGroupAsset request to a gRPC request.
func encodeGetGroupAssetRequest(_ context.Context, request interface{}) (interface{}, error) {
  r, ok := request.(endpoint1.GetGroupAssetRequest)
  if !ok {
    return nil, fmt.Errorf("request interface to endpoint GetGroupAssetRequest type assertion error")
  }
  return &pb.GetGroupAssetRequest{Groupid: r.GroupID}, nil
}

// decodeGetGroupAssetResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetGroupAssetResponse(_ context.Context, reply interface{}) (interface{}, error) {
  resp, ok := reply.(*pb.GetGroupAssetReply)
  if !ok{
    return nil, fmt.Errorf("pb GetGroupAssetReply type assertion error")
  }

  chainAssetsResp := make([]*repository.ChainAsset, len(resp.Chainassets))
  for i, ca := range resp.Chainassets {
    simpleTokens := make([]*repository.SimpleToken, len(ca.Simpletokens))
    for si, token := range ca.Simpletokens {
      simpleTokens[si] = &repository.SimpleToken{
        TokenID: token.Tokenid,
        Symbol: token.Symbol,
        Status: token.Status}
    }
    chainAssetsResp[i] = &repository.ChainAsset{
      ChainID: ca.Chainid,
      Name: ca.Name,
      Coin: ca.Coin,
      Status: ca.Status,
      SimpleTokens: simpleTokens}
  }

  return endpoint1.GetGroupAssetResponse{ChainAssets: chainAssetsResp}, nil
}
