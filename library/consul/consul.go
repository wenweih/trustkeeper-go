package consul

import (
  "os"
  "fmt"
  "time"
  consulsd "github.com/go-kit/kit/sd/consul"
  "github.com/hashicorp/consul/api"
  "github.com/go-kit/kit/log"
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

  consulConfig := api.DefaultConfig()
  consulConfig.Address = r.ConsulAddress
  consulClient, err := api.NewClient(consulConfig)
  if err != nil {
    return nil, err
  }
  client := consulsd.NewClient(consulClient)
  // agent := consulClient.Agent()

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

  // if err := agent.ServiceRegister(reg); err != nil {
  //   return err
  // }
  // return nil
  return consulsd.NewRegistrar(client, reg, logger), nil
}