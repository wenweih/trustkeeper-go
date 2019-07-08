package service

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"trustkeeper-go/app/service/account/pkg/configure"
	endpoint "trustkeeper-go/app/service/account/pkg/endpoint"
	grpc "trustkeeper-go/app/service/account/pkg/grpc"
	pb "trustkeeper-go/app/service/account/pkg/grpc/pb"
	service "trustkeeper-go/app/service/account/pkg/service"

	stdjwt "github.com/go-kit/kit/auth/jwt"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	"trustkeeper-go/library/consul"
	common "trustkeeper-go/library/util"

	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	lightsteptracergo "github.com/lightstep/lightstep-tracer-go"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	zipkingoopentracing "github.com/openzipkin/zipkin-go-opentracing"
	prometheus1 "github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	appdash "sourcegraph.com/sourcegraph/appdash"
	opentracing "sourcegraph.com/sourcegraph/appdash/opentracing"
)

var (
	logger log.Logger
	tracer opentracinggo.Tracer
	conf   configure.Conf
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("account", flag.ExitOnError)
var zipkinURL = fs.String("zipkin-url", "", "Enable Zipkin tracing via a collector URL e.g. http://localhost:9411/api/v1/spans")
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
	if *zipkinURL != "" {
		logger.Log("tracer", "Zipkin", "URL", *zipkinURL)
		collector, err := zipkingoopentracing.NewHTTPCollector(*zipkinURL)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		defer collector.Close()
		recorder := zipkingoopentracing.NewRecorder(collector, false, "localhost:80", "account")
		tracer, err = zipkingoopentracing.NewTracer(recorder)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
	} else if *lightstepToken != "" {
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

	c, err := configure.New()
	if err != nil {
		logger.Log(err)
	}
	conf = *c

	svc, err := service.New(conf, getServiceMiddleware(logger))
	if err != nil {
		logger.Log("err:", err.Error())
		os.Exit(1)
	}
	eps := endpoint.New(svc, getEndpointMiddleware(logger, svc))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(svc, g)
	logger.Log("exit", g.Run())
}
func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)
	// Add your GRPC options here
	// https://github.com/go-kit/kit/blob/master/auth/jwt/README.md GRPCToContext
	options["Roles"] = append(options["Roles"], grpctransport.ServerBefore(stdjwt.GRPCToContext()))
	options["UserInfo"] = append(options["UserInfo"], grpctransport.ServerBefore(stdjwt.GRPCToContext()))
	options["Signout"] = append(options["Signout"], grpctransport.ServerBefore(stdjwt.GRPCToContext()))
	options["Auth"] = append(options["Auth"], grpctransport.ServerBefore(stdjwt.GRPCToContext()))

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", common.LocalIP()+":0")
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	port := grpcListener.Addr().(*net.TCPAddr).Port
	consulReg := consul.NewConsulRegister(conf.ConsulAddress, common.AccountSrv, common.LocalIP(), port, []string{"account"})
	register, err := consulReg.NewConsulGRPCRegister()
	if err != nil {
		logger.Log("Get consul grpc register error: ", err.Error())
		os.Exit(1)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", grpcListener.Addr().String())
		baseServer := grpc1.NewServer()
		pb.RegisterAccountServer(baseServer, grpcServer)
		grpc_health_v1.RegisterHealthServer(baseServer, &consul.HealthImpl{})
		register.Register()
		return baseServer.Serve(grpcListener)
	}, func(error) {
		register.Deregister()
		grpcListener.Close()
	})

}
func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = addDefaultServiceMiddleware(logger, mw)
	// Append your middleware here

	return mw
}
func getEndpointMiddleware(logger log.Logger, s service.AccountService) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	duration := prometheus.NewSummaryFrom(prometheus1.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: "example",
		Subsystem: "account",
	}, []string{"method", "success"})
	addDefaultEndpointMiddleware(logger, duration, mw)
	// Add you endpoint middleware here
	mw["Roles"] = append(mw["Roles"], endpoint.AuthMiddleware())
	mw["Signout"] = append(mw["Signout"], endpoint.AuthMiddleware())
	mw["Auth"] = append(mw["Auth"], endpoint.AuthMiddleware())

	return mw
}
func initMetricsEndpoint(g *group.Group) {
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", "[::1]:0")
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", debugListener.Addr().String())
		return http.Serve(debugListener, http.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}

// This function just sits and waits for ctrl-C.
func initCancelInterrupt(s service.AccountService, g *group.Group) {
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
		s.Close()
		close(cancelInterrupt)
	})
}
