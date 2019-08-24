package client

import (
  "io"
  "fmt"
  "time"
  "github.com/go-kit/kit/log"
  "trustkeeper-go/app/service/account/pkg/service"
  accountEndpoint "trustkeeper-go/app/service/account/pkg/endpoint"
  "github.com/go-kit/kit/endpoint"

  sdconsul "github.com/go-kit/kit/sd/consul"
  stdjwt "github.com/go-kit/kit/auth/jwt"
  "github.com/go-kit/kit/sd/lb"
  "github.com/go-kit/kit/sd"
  "google.golang.org/grpc"
  grpctransport "github.com/go-kit/kit/transport/grpc"
  "trustkeeper-go/library/util"
  libconsule "trustkeeper-go/library/consul"
)

// https://github.com/go-kit/kit/blob/master/examples/apigateway/main.go
// https://github.com/go-kit/kit/blob/master/examples/profilesvc/client/client.go
// https://github.com/go-kit/kit/blob/e2d71a06a40aa95cb82ccd72e854893612c02db7/sd/consul/integration_test.go
func New(consulAddr string, logger log.Logger) (service.AccountService, error) {
  var (
    consulTags = []string{"account"}
    passingOnly = true
    retryMax = 1
    retryTimeout = 500 * time.Millisecond
  )

  consulClient, err := libconsule.NewClient(consulAddr)
  if err != nil {
    return nil, err
  }

  var (
    sdClient = sdconsul.NewClient(consulClient)
    registerSrvName = fmt.Sprintf("grpc.health.v1.%v", common.AccountSrv)
    instancer = sdconsul.NewInstancer(sdClient, logger, registerSrvName, consulTags, passingOnly)
    endpoints = accountEndpoint.Endpoints{}
  )
  {
    factory := factoryFor(accountEndpoint.MakeAuthEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.AuthEndpoint = retry
  }
  {
    factory := factoryFor(accountEndpoint.MakeRolesEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.RolesEndpoint = retry
  }
  {
    factory := factoryFor(accountEndpoint.MakeUserInfoEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.UserInfoEndpoint = retry
  }
  {
    factory := factoryFor(accountEndpoint.MakeCreateEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.CreateEndpoint = retry
  }
  {
    factory := factoryFor(accountEndpoint.MakeSigninEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.SigninEndpoint = retry
  }
  {
    factory := factoryFor(accountEndpoint.MakeSignoutEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.SignoutEndpoint = retry
  }
  return endpoints, nil
}

func factoryFor(makeEndpoint func(service.AccountService) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
    conn, err := grpc.Dial(instance, grpc.WithInsecure())
    if err != nil {
      return nil, nil, err
    }

    // https://github.com/go-kit/kit/blob/master/auth/jwt/README.md ContextToGRPC
    srv, err := newGRPCClient(conn, []grpctransport.ClientOption{(grpctransport.ClientBefore(stdjwt.ContextToGRPC()))})
		if err != nil {
			return nil, nil, err
		}

		endpoints := makeEndpoint(srv)
    return endpoints, conn, err
	}
}
