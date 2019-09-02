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
	endpoint "trustkeeper-go/app/service/chains_query/pkg/endpoint"
	service "trustkeeper-go/app/service/chains_query/pkg/service"
	"github.com/afex/hystrix-go/hystrix"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initGRPCHandler(endpoints, g)
	return g
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"BitcoincoreBlock":  {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "BitcoincoreBlock", logger))},
		"ConstructTxBTC":    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ConstructTxBTC", logger))},
		"ConstructTxERC20":  {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ConstructTxERC20", logger))},
		"ConstructTxETH":    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ConstructTxETH", logger))},
		"ConstructTxOmni":   {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ConstructTxOmni", logger))},
		"ERC20TokenInfo":    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ERC20TokenInfo", logger))},
		"QueryBalance":      {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "QueryBalance", logger))},
		"QueryOmniProperty": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "QueryOmniProperty", logger))},
		"SendBTCTx":         {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "SendBTCTx", logger))},
		"SendETHTx":         {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "SendETHTx", logger))},
		"WalletValidate":    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "WalletValidate", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	hystrix.ConfigureCommand("Lorem Request", hystrix.CommandConfig{Timeout: 1000})
	mw["BitcoincoreBlock"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "BitcoincoreBlock")),
		endpoint.InstrumentingMiddleware(duration.With("method", "BitcoincoreBlock")),
		endpoint.Hystrix("BitcoincoreBlock", "Service currently unavailable", logger),
	}
	mw["QueryOmniProperty"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "QueryOmniProperty")),
		endpoint.InstrumentingMiddleware(duration.With("method", "QueryOmniProperty")),
		endpoint.Hystrix("QueryOmniProperty", "Service currently unavailable", logger),
	}
	mw["ERC20TokenInfo"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "ERC20TokenInfo")),
		endpoint.InstrumentingMiddleware(duration.With("method", "ERC20TokenInfo")),
		endpoint.Hystrix("ERC20TokenInfo", "Service currently unavailable", logger),
	}
	mw["ConstructTxBTC"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "ConstructTxBTC")),
		endpoint.InstrumentingMiddleware(duration.With("method", "ConstructTxBTC")),
		endpoint.Hystrix("ConstructTxBTC", "Service currently unavailable", logger),
	}
	mw["SendBTCTx"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "SendBTCTx")),
		endpoint.InstrumentingMiddleware(duration.With("method", "SendBTCTx")),
		endpoint.Hystrix("SendBTCTx", "Service currently unavailable", logger),
	}
	mw["ConstructTxOmni"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "ConstructTxOmni")),
		endpoint.InstrumentingMiddleware(duration.With("method", "ConstructTxOmni")),
		endpoint.Hystrix("ConstructTxOmni", "Service currently unavailable", logger),
	}
	mw["ConstructTxETH"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "ConstructTxETH")),
		endpoint.InstrumentingMiddleware(duration.With("method", "ConstructTxETH")),
		endpoint.Hystrix("ConstructTxETH", "Service currently unavailable", logger),
	}
	mw["SendETHTx"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "SendETHTx")),
		endpoint.InstrumentingMiddleware(duration.With("method", "SendETHTx")),
		endpoint.Hystrix("SendETHTx", "Service currently unavailable", logger),
	}
	mw["ConstructTxERC20"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "ConstructTxERC20")),
		endpoint.InstrumentingMiddleware(duration.With("method", "ConstructTxERC20")),
		endpoint.Hystrix("ConstructTxERC20", "Service currently unavailable", logger),
	}
	mw["QueryBalance"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "QueryBalance")),
		endpoint.InstrumentingMiddleware(duration.With("method", "QueryBalance")),
		endpoint.Hystrix("QueryBalance", "Service currently unavailable", logger),
	}
	mw["WalletValidate"] = []endpoint1.Middleware{
		endpoint.LoggingMiddleware(log.With(logger, "method", "WalletValidate")),
		endpoint.InstrumentingMiddleware(duration.With("method", "WalletValidate")),
		endpoint.Hystrix("WalletValidate", "Service currently unavailable", logger),
	}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"BitcoincoreBlock", "QueryOmniProperty", "ERC20TokenInfo", "ConstructTxBTC", "SendBTCTx", "ConstructTxOmni", "ConstructTxETH", "SendETHTx", "ConstructTxERC20", "QueryBalance", "WalletValidate"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
