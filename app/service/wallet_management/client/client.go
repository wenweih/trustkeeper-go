package client

import (
  "io"
  "fmt"
  "time"
  "github.com/go-kit/kit/log"
  "trustkeeper-go/app/service/wallet_management/pkg/service"
  walletmanagementEndpoint "trustkeeper-go/app/service/wallet_management/pkg/endpoint"
  "github.com/go-kit/kit/endpoint"
  sdconsul "github.com/go-kit/kit/sd/consul"
  "github.com/go-kit/kit/sd/lb"
  "github.com/go-kit/kit/sd"
  "google.golang.org/grpc"
  grpctransport "github.com/go-kit/kit/transport/grpc"
  libconsule "trustkeeper-go/library/consul"
  "trustkeeper-go/library/util"
)

// https://github.com/go-kit/kit/blob/master/examples/apigateway/main.go
// https://github.com/go-kit/kit/blob/master/examples/profilesvc/client/client.go
// https://github.com/go-kit/kit/blob/e2d71a06a40aa95cb82ccd72e854893612c02db7/sd/consul/integration_test.go
func New(consulAddr string, logger log.Logger) (service.WalletManagementService, error) {
  var (
    consulTags = []string{"wallet"}
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
    registerSrvName = fmt.Sprintf("grpc.health.v1.%v", common.WalletManagementSrv)
    instancer = sdconsul.NewInstancer(sdClient, logger, registerSrvName, consulTags, passingOnly)
    endpoints = walletmanagementEndpoint.Endpoints{}
  )
  {
    factory := factoryFor(walletmanagementEndpoint.MakeCreateChainEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.CreateChainEndpoint = retry
  }
  return endpoints, nil
}

func factoryFor(makeEndpoint func(service.WalletManagementService) endpoint.Endpoint) sd.Factory {
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
