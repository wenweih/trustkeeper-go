// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	endpoint "trustkeeper-go/app/gateway/webapi/pkg/endpoint"
	http1 "trustkeeper-go/app/gateway/webapi/pkg/http"
	service "trustkeeper-go/app/gateway/webapi/pkg/service"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"ChangeGroupAssets": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ChangeGroupAssets", logger))},
		"CreateGroup":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateGroup", logger))},
		"CreateToken":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateToken", logger))},
		"CreateWallet":      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateWallet", logger))},
		"EthToken":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "EthToken", logger))},
		"GetGroupAssets":    {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroupAssets", logger))},
		"GetGroups":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroups", logger))},
		"GetRoles":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetRoles", logger))},
		"GetWallets":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetWallets", logger))},
		"QueryBalance":      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "QueryBalance", logger))},
		"QueryOmniProperty": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "QueryOmniProperty", logger))},
		"SendBTCTx":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SendBTCTx", logger))},
		"SendETHTx":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SendETHTx", logger))},
		"Signin":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Signin", logger))},
		"Signout":           {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Signout", logger))},
		"Signup":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Signup", logger))},
		"UpdateGroup":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateGroup", logger))},
		"UserInfo":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UserInfo", logger))},
		"WalletValidate":    {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "WalletValidate", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Signup"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Signup")), endpoint.InstrumentingMiddleware(duration.With("method", "Signup"))}
	mw["Signin"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Signin")), endpoint.InstrumentingMiddleware(duration.With("method", "Signin"))}
	mw["Signout"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Signout")), endpoint.InstrumentingMiddleware(duration.With("method", "Signout"))}
	mw["GetRoles"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetRoles")), endpoint.InstrumentingMiddleware(duration.With("method", "GetRoles"))}
	mw["UserInfo"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "UserInfo")), endpoint.InstrumentingMiddleware(duration.With("method", "UserInfo"))}
	mw["GetGroups"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetGroups")), endpoint.InstrumentingMiddleware(duration.With("method", "GetGroups"))}
	mw["CreateGroup"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CreateGroup")), endpoint.InstrumentingMiddleware(duration.With("method", "CreateGroup"))}
	mw["UpdateGroup"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "UpdateGroup")), endpoint.InstrumentingMiddleware(duration.With("method", "UpdateGroup"))}
	mw["GetGroupAssets"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetGroupAssets")), endpoint.InstrumentingMiddleware(duration.With("method", "GetGroupAssets"))}
	mw["ChangeGroupAssets"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "ChangeGroupAssets")), endpoint.InstrumentingMiddleware(duration.With("method", "ChangeGroupAssets"))}
	mw["CreateWallet"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CreateWallet")), endpoint.InstrumentingMiddleware(duration.With("method", "CreateWallet"))}
	mw["GetWallets"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetWallets")), endpoint.InstrumentingMiddleware(duration.With("method", "GetWallets"))}
	mw["QueryOmniProperty"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "QueryOmniProperty")), endpoint.InstrumentingMiddleware(duration.With("method", "QueryOmniProperty"))}
	mw["EthToken"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "EthToken")), endpoint.InstrumentingMiddleware(duration.With("method", "EthToken"))}
	mw["CreateToken"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CreateToken")), endpoint.InstrumentingMiddleware(duration.With("method", "CreateToken"))}
	mw["SendBTCTx"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SendBTCTx")), endpoint.InstrumentingMiddleware(duration.With("method", "SendBTCTx"))}
	mw["SendETHTx"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SendETHTx")), endpoint.InstrumentingMiddleware(duration.With("method", "SendETHTx"))}
	mw["QueryBalance"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "QueryBalance")), endpoint.InstrumentingMiddleware(duration.With("method", "QueryBalance"))}
	mw["WalletValidate"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "WalletValidate")), endpoint.InstrumentingMiddleware(duration.With("method", "WalletValidate"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Signup", "Signin", "Signout", "GetRoles", "UserInfo", "GetGroups", "CreateGroup", "UpdateGroup", "GetGroupAssets", "ChangeGroupAssets", "CreateWallet", "GetWallets", "QueryOmniProperty", "EthToken", "CreateToken", "SendBTCTx", "SendETHTx", "QueryBalance", "WalletValidate"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
