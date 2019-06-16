package consul

import (
  "os"
  "fmt"
  "time"
  "context"
  consulsd "github.com/go-kit/kit/sd/consul"
  "github.com/hashicorp/consul/api"
  "github.com/go-kit/kit/log"
  "google.golang.org/grpc/health/grpc_health_v1"
)

type ConsulRegister struct {
  ConsulAddress   string  // consul address
  ServiceName     string  // service name
  ServiceIP       string
  Tags            []string // consul tags
  ServicePort      int //service port
  DeregisterCriticalServiceAfter time.Duration
  Interval  time.Duration
}

func NewConsulRegister(consulAddress, serviceName, serviceIP string, servicePort int, tags []string ) *ConsulRegister {
  return &ConsulRegister {
    ConsulAddress: consulAddress,
    ServiceName: serviceName,
    ServiceIP: serviceIP,
    Tags: tags,
    ServicePort: servicePort,
    DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
    Interval: time.Duration(10) * time.Second,
  }
}

// https://github.com/ru-rocker/gokit-playground/blob/master/lorem-consul/register.go
// https://github.com/hatlonely/hellogolang/blob/master/sample/addservice/internal/grpcsr/consul_register.go
func (r *ConsulRegister) NewConsulGRPCRegister() (*consulsd.Registrar, error) {
  var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

  consulClient, err := NewClient(r.ConsulAddress)
  if err != nil {
    return nil, err
  }
  client := consulsd.NewClient(consulClient)

  reg := &api.AgentServiceRegistration{
    ID: fmt.Sprintf("%v-%v-%v", r.ServiceName, r.ServiceIP, r.ServicePort),
    Name: fmt.Sprintf("grpc.health.v1.%v", r.ServiceName),
    Tags: r.Tags,
    Port: r.ServicePort,
    Address: r.ServiceIP,
    Check: &api.AgentServiceCheck{
      // 健康检查间隔
      Interval: r.Interval.String(),
      //grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
      GRPC: fmt.Sprintf("%v:%v/%v", r.ServiceIP, r.ServicePort, r.ServiceName),
      // 注销时间，相当于过期时间
      DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),
    },
  }
  return consulsd.NewRegistrar(client, reg, logger), nil
}


func NewClient(consulAddr string) (*api.Client, error) {
  consulConfig := api.DefaultConfig()
  consulConfig.Address = consulAddr
  consulClient, err := api.NewClient(consulConfig)
  if err != nil {
    return nil, err
  }
  return consulClient, nil
}

// HealthImpl grpc 健康检查
// https://studygolang.com/articles/18737
type HealthImpl struct{}
// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
// https://github.com/hashicorp/consul/blob/master/agent/checks/grpc.go
// consul 检查服务器的健康状态，consul 用 google.golang.org/grpc/health/grpc_health_v1.HealthServer 接口，实现了对 grpc健康检查的支持，所以我们只需要实现先这个接口，consul 就能利用这个接口作健康检查了
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
// Watch HealthServer interface 有两个方法
// Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
// Watch(*HealthCheckRequest, Health_WatchServer) error
// 所以在 HealthImpl 结构提不仅要实现 Check 方法，还要实现 Watch 方法
func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}
