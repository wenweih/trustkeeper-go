package service

import (
	"flag"
	"fmt"
	"net"
	http1 "net/http"
	"os"
	"os/signal"
	"syscall"
	endpoint "trustkeeper-go/app/gateway/webapi/pkg/endpoint"
	http "trustkeeper-go/app/gateway/webapi/pkg/http"
	service "trustkeeper-go/app/gateway/webapi/pkg/service"

	"github.com/rs/cors"

	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	lightsteptracergo "github.com/lightstep/lightstep-tracer-go"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	prometheus1 "github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	appdash "sourcegraph.com/sourcegraph/appdash"
	opentracing "sourcegraph.com/sourcegraph/appdash/opentracing"

	stdjwt "github.com/go-kit/kit/auth/jwt"
	httptransport "github.com/go-kit/kit/transport/http"
)

var tracer opentracinggo.Tracer
var logger log.Logger

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("webapi", flag.ExitOnError)
var debugAddr = fs.String("debug.addr", ":7979", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
var lightstepToken = fs.String("lightstep-token", "", "Enable LightStep tracing via a LightStep access token")
var appdashAddr = fs.String("appdash-addr", "", "Enable Appdash tracing via an Appdash server host:port")

func Run() {
	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	//  Determine which tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency
	if *lightstepToken != "" {
		logger.Log("tracer", "LightStep")
		tracer = lightsteptracergo.NewTracer(lightsteptracergo.Options{AccessToken: *lightstepToken})
		defer lightsteptracergo.FlushLightStepTracer(tracer)
	} else if *appdashAddr != "" {
		logger.Log("tracer", "Appdash", "addr", *appdashAddr)
		collector := appdash.NewRemoteCollector(*appdashAddr)
		tracer = opentracing.NewTracer(collector)
		defer collector.Close()
	} else {
		logger.Log("tracer", "none")
		tracer = opentracinggo.GlobalTracer()
	}

	svc, err := service.New(logger, getServiceMiddleware(logger))
	if err != nil {
		logger.Log("err: ", err.Error())
		os.Exit(1)
	}
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger, tracer)

	// Add your http options here
	// 为了添加 Auth endpoint middleware, 并且 middleware 中拿不到 raw request 的信息的情况下
	// 避免网关过滤鉴权接口都要在 decodeXXXRequest 时 extract request header Authorization
	// 所以需要把请求头的 Authorization 参数到请求上下文中
	// https://github.com/go-kit/kit/blob/master/auth/jwt/README.md  HTTPToContext
	options["GetRoles"] = append(options["GetRoles"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["UserInfo"] = append(options["UserInfo"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["Signout"] = append(options["Signout"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["GetGroups"] = append(options["GetGroups"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["CreateGroup"] = append(options["CreateGroup"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["UpdateGroup"] = append(options["UpdateGroup"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["GetGroupAssets"] = append(options["GetGroupAssets"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["ChangeGroupAssets"] = append(options["ChangeGroupAssets"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["CreateWallet"] = append(options["CreateWallet"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["GetWallets"] = append(options["GetWallets"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["QueryOmniProperty"] = append(options["QueryOmniProperty"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["CreateToken"] = append(options["CreateToken"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["EthToken"] = append(options["EthToken"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["SendBTCTx"] = append(options["SendBTCTx"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["QueryBalance"] = append(options["QueryBalance"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["WalletValidate"] = append(options["WalletValidate"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["SendETHTx"] = append(options["SendETHTx"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["SendERC20Tx"] = append(options["SendERC20Tx"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	options["SendOmniTx"] = append(options["SendOmniTx"], httptransport.ServerBefore(stdjwt.HTTPToContext()))
	httpHandler := http.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})
	handler := c.Handler(httpHandler)
	// handler := cors.Default().Handler(httpHandler)
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http1.Serve(httpListener, handler)
	}, func(error) {
		httpListener.Close()
	})
}
func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = addDefaultServiceMiddleware(logger, mw)
	// Append your middleware here

	return
}
func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	duration := prometheus.NewSummaryFrom(prometheus1.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: "example",
		Subsystem: "webapi",
	}, []string{"method", "success"})
	addDefaultEndpointMiddleware(logger, duration, mw)

	// 添加 Auth endpoint middleware
	mw["GetRoles"] = append(mw["GetRoles"], endpoint.AuthMiddleware())
	mw["Signout"] = append(mw["Signout"], endpoint.AuthMiddleware())
	return mw
}
func initMetricsEndpoint(g *group.Group) {
	http1.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", *debugAddr)
		return http1.Serve(debugListener, http1.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}
func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}
