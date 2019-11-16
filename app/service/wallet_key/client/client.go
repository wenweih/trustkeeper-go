package client

import (
  "io"
  "fmt"
  "time"
  "github.com/go-kit/kit/log"
  "trustkeeper-go/app/service/wallet_key/pkg/service"
  walletKeyEndpoint "trustkeeper-go/app/service/wallet_key/pkg/endpoint"
  "github.com/go-kit/kit/endpoint"

  sdconsul "github.com/go-kit/kit/sd/consul"
  "github.com/go-kit/kit/sd/lb"
  "github.com/go-kit/kit/sd"
  "google.golang.org/grpc"
  grpctransport "github.com/go-kit/kit/transport/grpc"
  "trustkeeper-go/library/util"
  libconsule "trustkeeper-go/library/consul"
)

func New(consulAddr string, logger log.Logger) (service.WalletKeyService, error) {
  var (
    consulTags = []string{"key"}
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
    registerSrvName = fmt.Sprintf("grpc.health.v1.%v", common.WalletKeySrv)
    instancer = sdconsul.NewInstancer(sdClient, logger, registerSrvName, consulTags, passingOnly)
    endpoints = walletKeyEndpoint.Endpoints{}
  )
  {
    factory := factoryFor(walletKeyEndpoint.MakeGenerateMnemonicEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.GenerateMnemonicEndpoint = retry
  }
  {
    factory := factoryFor(walletKeyEndpoint.MakeSignedBitcoincoreTxEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.SignedBitcoincoreTxEndpoint = retry
  }
  {
    factory := factoryFor(walletKeyEndpoint.MakeSignedEthereumTxEndpoint)
    endpointer := sd.NewEndpointer(instancer, factory, logger)
    balancer := lb.NewRoundRobin(endpointer)
    retry := lb.Retry(retryMax, retryTimeout, balancer)
    endpoints.SignedEthereumTxEndpoint = retry
  }
  return endpoints, nil
}

func factoryFor(makeEndpoint func(service.WalletKeyService) endpoint.Endpoint) sd.Factory {
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
