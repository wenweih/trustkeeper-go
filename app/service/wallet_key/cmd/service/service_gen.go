// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint "trustkeeper-go/app/service/wallet_key/pkg/endpoint"
	service "trustkeeper-go/app/service/wallet_key/pkg/service"

	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	grpc "github.com/go-kit/kit/transport/grpc"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initGRPCHandler(endpoints, g)
	return g
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"GenerateMnemonic":    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GenerateMnemonic", logger))},
		"SignedBitcoincoreTx": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "SignedBitcoincoreTx", logger))},
		"SignedEthereumTx":    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "SignedEthereumTx", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["GenerateMnemonic"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GenerateMnemonic")), endpoint.InstrumentingMiddleware(duration.With("method", "GenerateMnemonic"))}
	mw["SignedBitcoincoreTx"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SignedBitcoincoreTx")), endpoint.InstrumentingMiddleware(duration.With("method", "SignedBitcoincoreTx"))}
	mw["SignedEthereumTx"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SignedEthereumTx")), endpoint.InstrumentingMiddleware(duration.With("method", "SignedEthereumTx"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"GenerateMnemonic", "SignedBitcoincoreTx", "SignedEthereumTx"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
