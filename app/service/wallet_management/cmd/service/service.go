package service

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"trustkeeper-go/app/service/wallet_management/pkg/configure"
	endpoint "trustkeeper-go/app/service/wallet_management/pkg/endpoint"
	grpc "trustkeeper-go/app/service/wallet_management/pkg/grpc"
	pb "trustkeeper-go/app/service/wallet_management/pkg/grpc/pb"
	"trustkeeper-go/app/service/wallet_management/pkg/jobs"
	service "trustkeeper-go/app/service/wallet_management/pkg/service"
	"trustkeeper-go/library/consul"
	common "trustkeeper-go/library/util"

	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	lightsteptracergo "github.com/lightstep/lightstep-tracer-go"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	prometheus1 "github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	appdash "sourcegraph.com/sourcegraph/appdash"
	opentracing "sourcegraph.com/sourcegraph/appdash/opentracing"
)

var (
	conf   configure.Conf
	tracer opentracinggo.Tracer
	logger log.Logger
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("wallet_management", flag.ExitOnError)
var thriftFramed = fs.Bool("thrift-framed", false, "true to enable framing")
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

	c, err := configure.New()
	if err != nil {
		logger.Log("configure err: ", err.Error())
		os.Exit(1)
	}
	conf = *c

	svc, err := service.New(conf, logger, getServiceMiddleware(logger))
	if err != nil {
		logger.Log("svc error: ", err.Error())
		os.Exit(1)
	}
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initJobs(logger, g)
	initMetricsEndpoint(g)
	initCancelInterrupt(svc, g)
	logger.Log("exit", g.Run())

}
func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)
	// Add your GRPC options here
	grpcServer := grpc.NewGRPCServer(endpoints, options)

	// call grpc listen should not use Loopback address
	// https://stackoverflow.com/questions/43911793/cannot-connect-to-go-grpc-server-running-in-local-docker-container
	grpcListener, err := net.Listen("tcp", common.LocalIP()+":0")
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	port := grpcListener.Addr().(*net.TCPAddr).Port
	consulReg := consul.NewConsulRegister(conf.ConsulAddress, common.WalletManagementSrv, common.LocalIP(), port, []string{"wallet"})
	register, err := consulReg.NewConsulGRPCRegister()
	if err != nil {
		logger.Log("Get consul grpc register error: ", err.Error())
		os.Exit(1)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", grpcListener.Addr().String())
		baseServer := grpc1.NewServer()
		pb.RegisterWalletManagementServer(baseServer, grpcServer)
		grpc_health_v1.RegisterHealthServer(baseServer, &consul.HealthImpl{})
		register.Register()
		return baseServer.Serve(grpcListener)
	}, func(error) {
		register.Deregister()
		grpcListener.Close()
	})
}

func initJobs(logger log.Logger, g *group.Group) {
	jobSvc, err := service.NewJobsService(conf, logger)
	if err != nil {
		logger.Log("init job service error: ", err.Error())
		os.Exit(1)
	}
	workPool := jobs.New(logger, conf.Redis, jobSvc)
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		workPool.Start()
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(err error) {
		workPool.Stop()
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
		Subsystem: "wallet_management",
	}, []string{"method", "success"})
	addDefaultEndpointMiddleware(logger, duration, mw)
	// Add you endpoint middleware here

	return
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
func initCancelInterrupt(s service.WalletManagementService, g *group.Group) {
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
