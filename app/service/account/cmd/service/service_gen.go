// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	grpc "github.com/go-kit/kit/transport/grpc"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	endpoint "trustkeeper-go/app/service/account/pkg/endpoint"
	service "trustkeeper-go/app/service/account/pkg/service"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initGRPCHandler(endpoints, g)
	return g
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"Create":  {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "Create", logger))},
		"Signin":  {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "Signin", logger))},
		"Signout": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "Signout", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Create"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Create")), endpoint.InstrumentingMiddleware(duration.With("method", "Create"))}
	mw["Signin"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Signin")), endpoint.InstrumentingMiddleware(duration.With("method", "Signin"))}
	mw["Signout"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Signout")), endpoint.InstrumentingMiddleware(duration.With("method", "Signout"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Create", "Signin", "Signout"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
