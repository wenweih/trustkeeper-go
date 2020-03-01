package client

import (
	"context"
	"fmt"
	"io"
	"time"
	walletmanagementEndpoint "trustkeeper-go/app/service/wallet_management/pkg/endpoint"
	"trustkeeper-go/app/service/wallet_management/pkg/service"
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

// https://github.com/go-kit/kit/blob/master/examples/apigateway/main.go
// https://github.com/go-kit/kit/blob/master/examples/profilesvc/client/client.go
// https://github.com/go-kit/kit/blob/e2d71a06a40aa95cb82ccd72e854893612c02db7/sd/consul/integration_test.go
func New(consulAddr string, logger log.Logger) (service.WalletManagementService, error) {
	var (
		consulTags   = []string{"wallet"}
		passingOnly  = true
		retryMax     = 3
		retryTimeout = 1000 * time.Millisecond
	)

	consulClient, err := libconsule.NewClient(consulAddr)
	if err != nil {
		return nil, err
	}

	var (
		sdClient        = sdconsul.NewClient(consulClient)
		registerSrvName = fmt.Sprintf("grpc.health.v1.%v", common.WalletManagementSrv)
		instancer       = sdconsul.NewInstancer(sdClient, logger, registerSrvName, consulTags, passingOnly)
		endpoints       = walletmanagementEndpoint.Endpoints{}
	)
	{
		factory := factoryFor(walletmanagementEndpoint.MakeCreateChainEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CreateChainEndpoint = retry
	}
	{
		factory := factoryFor(walletmanagementEndpoint.MakeAssignedXpubToGroupEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.AssignedXpubToGroupEndpoint = retry
	}
	{
		factory := factoryFor(walletmanagementEndpoint.MakeGetChainsEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.GetChainsEndpoint = retry
	}
	{
		factory := factoryFor(walletmanagementEndpoint.MakeCreateWalletEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CreateWalletEndpoint = retry
	}
	{
		factory := factoryFor(walletmanagementEndpoint.MakeGetWalletsEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.GetWalletsEndpoint = retry
	}
	{
		factory := factoryFor(walletmanagementEndpoint.MakeQueryWalletsForGroupByChainNameEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.QueryWalletsForGroupByChainNameEndpoint = retry
	}

	{
		factory := factoryFor(walletmanagementEndpoint.MakeQueryWalletHDEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.QueryWalletHDEndpoint = retry
	}
	return endpoints, nil
}

func factoryFor(makeEndpoint func(service.WalletManagementService) endpoint.Endpoint) sd.Factory {
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
