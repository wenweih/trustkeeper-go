package etcd

import (
  "context"
  sdetcd "github.com/go-kit/kit/sd/etcdv3"
  log "github.com/go-kit/kit/log"
)

// RegisterService https://dev.to/plutov/packagemain-13-microservices-with-go-kit-part-2-4lgh
func RegisterService(etcdServer, prefix, instance string, logger log.Logger) (*sdetcd.Registrar, error) {
  var key = prefix + instance
	client, err := sdetcd.NewClient(context.Background(), []string{etcdServer}, sdetcd.ClientOptions{})
	if err != nil {
		return nil, err
	}
	registrar := sdetcd.NewRegistrar(client, sdetcd.Service{
		Key:   key,
		Value: instance,
	}, logger)

	registrar.Register()

	return registrar, nil
}
