package client

import (
  "io"
  "fmt"
  "time"
  "github.com/go-kit/kit/log"
  "trustkeeper-go/app/service/dashboard/pkg/service"
  dashboardEndpoint "trustkeeper-go/app/service/dashboard/pkg/endpoint"
  "github.com/go-kit/kit/endpoint"

  sdconsul "github.com/go-kit/kit/sd/consul"
  "github.com/go-kit/kit/sd/lb"
  "github.com/go-kit/kit/sd"
  "google.golang.org/grpc"
  grpctransport "github.com/go-kit/kit/transport/grpc"
  "trustkeeper-go/Library/common"
  libconsule "trustkeeper-go/library/consul"
)

// https://github.com/go-kit/kit/blob/master/examples/apigateway/main.go
// https://github.com/go-kit/kit/blob/master/examples/profilesvc/client/client.go
// https://github.com/go-kit/kit/blob/e2d71a06a40aa95cb82ccd72e854893612c02db7/sd/consul/integration_test.go
func New(consulAddr string, logger log.Logger) (service.DashboardService, error) {
  var (
    consulTags = []string{"dashboard"}
    passingOnly = true
    retryMax = 3
    retryTimeout = 500 * time.Millisecond
  )

  consulClient, err := libconsule.NewClient(consulAddr)
  if err != nil {
    return nil, err
  }

  var (
    sdClient = sdconsul.NewClient(consulClient)
    registerSrvName = fmt.Sprintf("grpc.health.v1.%v", common.DashboardSrv)
    instancer = sdconsul.NewInstancer(sdClient, logger, registerSrvName, consulTags, passingOnly)
    endpoints = dashboardEndpoint.Endpoints{}
  )
  {
    factory := factoryFor(dashboardEndpoint.MakeGetGroupsEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.GetGroupsEndpoint = retry
  }
  {
    factory := factoryFor(dashboardEndpoint.MakeCreateGroupEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.CreateGroupEndpoint = retry
  }
  return endpoints, nil
}

func factoryFor(makeEndpoint func(service.DashboardService) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
    conn, err := grpc.Dial(instance, grpc.WithInsecure())
    if err != nil {
      return nil, nil, err
    }

    srv, err := newGRPCClient(conn, []grpctransport.ClientOption{})
		if err != nil {
			return nil, nil, err
		}

		endpoints := makeEndpoint(srv)
    return endpoints, conn, err
	}
}
