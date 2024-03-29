package client

import (
	"fmt"
	"io"
	"time"
	transactionEndpoint "trustkeeper-go/app/service/transaction/pkg/endpoint"
	"trustkeeper-go/app/service/transaction/pkg/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"context"
	libconsule "trustkeeper-go/library/consul"
	common "trustkeeper-go/library/util"

	"github.com/go-kit/kit/sd"
	sdconsul "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func New(consulAddr string, logger log.Logger) (service.TransactionService, error) {
	var (
		consulTags   = []string{"transaction"}
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
		registerSrvName = fmt.Sprintf("grpc.health.v1.%v", common.TxSrv)
		instancer       = sdconsul.NewInstancer(sdClient, logger, registerSrvName, consulTags, passingOnly)
		endpoints       = transactionEndpoint.Endpoints{}
	)
	{
		factory := factoryFor(transactionEndpoint.MakeAssignAssetsToWalletEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.AssignAssetsToWalletEndpoint = retry
	}
	{
		factory := factoryFor(transactionEndpoint.MakeCreateBalancesForAssetEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CreateBalancesForAssetEndpoint = retry
	}
	return endpoints, nil
}

func factoryFor(makeEndpoint func(service.TransactionService) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		srv, err := newGRPCClient(
			conn,
			[]grpctransport.ClientOption{grpctransport.ClientBefore(contextToGRPC())})
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
