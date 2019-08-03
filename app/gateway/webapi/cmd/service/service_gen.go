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
		"CreateWallet":      {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateWallet", logger))},
		"GetGroupAssets":    {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroupAssets", logger))},
		"GetGroups":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetGroups", logger))},
		"GetRoles":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetRoles", logger))},
		"GetWallets":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetWallets", logger))},
		"QueryOmniProperty": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "QueryOmniProperty", logger))},
		"QueryToken":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "QueryToken", logger))},
		"Signin":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Signin", logger))},
		"Signout":           {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Signout", logger))},
		"Signup":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Signup", logger))},
		"UpdateGroup":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateGroup", logger))},
		"UserInfo":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UserInfo", logger))},
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
	mw["QueryToken"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "QueryToken")), endpoint.InstrumentingMiddleware(duration.With("method", "QueryToken"))}
	mw["QueryOmniProperty"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "QueryOmniProperty")), endpoint.InstrumentingMiddleware(duration.With("method", "QueryOmniProperty"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Signup", "Signin", "Signout", "GetRoles", "UserInfo", "GetGroups", "CreateGroup", "UpdateGroup", "GetGroupAssets", "ChangeGroupAssets", "CreateWallet", "GetWallets", "QueryToken", "QueryOmniProperty"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
