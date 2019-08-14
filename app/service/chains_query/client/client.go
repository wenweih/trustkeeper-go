package client

import (
  "io"
  "fmt"
  "time"
  "context"
  "github.com/go-kit/kit/log"
  "trustkeeper-go/app/service/chains_query/pkg/service"
  chainsquerytEndpoint "trustkeeper-go/app/service/chains_query/pkg/endpoint"
  "github.com/go-kit/kit/endpoint"
  sdconsul "github.com/go-kit/kit/sd/consul"
  "github.com/go-kit/kit/sd/lb"
  "github.com/go-kit/kit/sd"
  "google.golang.org/grpc"
  grpctransport "github.com/go-kit/kit/transport/grpc"
  libconsule "trustkeeper-go/library/consul"
  "trustkeeper-go/library/util"
  "google.golang.org/grpc/metadata"
)

func New(consulAddr string, logger log.Logger) (service.ChainsQueryService, error) {
  var (
    consulTags = []string{"chainsquery"}
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
    registerSrvName = fmt.Sprintf("grpc.health.v1.%v", common.ChainsQuerySrv)
    instancer = sdconsul.NewInstancer(sdClient, logger, registerSrvName, consulTags, passingOnly)
    endpoints = chainsquerytEndpoint.Endpoints{}
  )
  {
    factory := factoryFor(chainsquerytEndpoint.MakeBitcoincoreBlockEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.BitcoincoreBlockEndpoint = retry
  }
  {
    factory := factoryFor(chainsquerytEndpoint.MakeQueryOmniPropertyEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.QueryOmniPropertyEndpoint = retry
  }
  {
    factory := factoryFor(chainsquerytEndpoint.MakeERC20TokenInfoEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.ERC20TokenInfoEndpoint = retry
  }
  {
    factory := factoryFor(chainsquerytEndpoint.MakeConstructTxBTCEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.ConstructTxBTCEndpoint = retry
  }
  return endpoints, nil
}

func factoryFor(makeEndpoint func(service.ChainsQueryService) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
    conn, err := grpc.Dial(instance, grpc.WithInsecure())
    if err != nil {
      return nil, nil, err
    }

    srv, err := newGRPCClient(conn, []grpctransport.ClientOption{grpctransport.ClientBefore(contextToGRPC())})
		if err != nil {
			return nil, nil, err
		}

		endpoints := makeEndpoint(srv)
    return endpoints, conn, err
	}
}

func contextToGRPC() grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		if authinfo, ok := ctx.Value("auth").(struct{Roles []string; UID string; NID string}); ok {
      // capital "Key" is illegal in HTTP/2.
      (*md)["roles"] = authinfo.Roles
      (*md)["uid"] = []string{authinfo.UID}
      (*md)["nid"] = []string{authinfo.NID}
    }
		return ctx
	}
}
