package client

import (
	"context"
	"fmt"
	"io"
	"time"
	chainsquerytEndpoint "trustkeeper-go/app/service/chains_query/pkg/endpoint"
	"trustkeeper-go/app/service/chains_query/pkg/service"
	libconsule "trustkeeper-go/library/consul"
	common "trustkeeper-go/library/util"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	sdconsul "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func New(consulAddr string, logger log.Logger) (service.ChainsQueryService, error) {
	var (
		consulTags   = []string{"chainsquery"}
		passingOnly  = true
		retryMax     = 3
		retryTimeout = 500 * time.Millisecond
	)

	consulClient, err := libconsule.NewClient(consulAddr)
	if err != nil {
		return nil, err
	}

	var (
		sdClient        = sdconsul.NewClient(consulClient)
		registerSrvName = fmt.Sprintf("grpc.health.v1.%v", common.ChainsQuerySrv)
		instancer       = sdconsul.NewInstancer(sdClient, logger, registerSrvName, consulTags, passingOnly)
		endpoints       = chainsquerytEndpoint.Endpoints{}
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
	{
		factory := factoryFor(chainsquerytEndpoint.MakeSendBTCTxEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.SendBTCTxEndpoint = retry
	}
	{
		factory := factoryFor(chainsquerytEndpoint.MakeQueryBalanceEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.QueryBalanceEndpoint = retry
	}
	{
		factory := factoryFor(chainsquerytEndpoint.MakeWalletValidateEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.WalletValidateEndpoint = retry
	}
	{
		factory := factoryFor(chainsquerytEndpoint.MakeConstructTxETHEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.ConstructTxETHEndpoint = retry
	}
	{
		factory := factoryFor(chainsquerytEndpoint.MakeSendETHTxEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.SendETHTxEndpoint = retry
	}
	{
		factory := factoryFor(chainsquerytEndpoint.MakeConstructTxERC20Endpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.ConstructTxERC20Endpoint = retry
	}
	{
		factory := factoryFor(chainsquerytEndpoint.MakeConstructTxOmniEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.ConstructTxOmniEndpoint = retry
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
		if authinfo, ok := ctx.Value("auth").(struct {
			Roles []string
			UID   string
			NID   string
		}); ok {
			// capital "Key" is illegal in HTTP/2.
			(*md)["roles"] = authinfo.Roles
			(*md)["uid"] = []string{authinfo.UID}
			(*md)["nid"] = []string{authinfo.NID}
		}
		return ctx
	}
}
