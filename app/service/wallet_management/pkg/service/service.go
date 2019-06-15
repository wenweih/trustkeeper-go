package service

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// WalletManagementService describes the service.
type WalletManagementService interface {
	// Add your methods here
	CreateChain(ctx context.Context, symbol, bit44ID string, status bool) (err error)
}

type basicWalletManagementService struct{}

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

func (b *basicWalletManagementService) CreateChain(ctx context.Context, symbol string, bit44ID string, status bool) (err error) {
	// TODO implement the business logic of CreateChain
	return err
}

// NewBasicWalletManagementService returns a naive, stateless implementation of WalletManagementService.
func NewBasicWalletManagementService() WalletManagementService {
	return &basicWalletManagementService{}
}

// New returns a WalletManagementService with all of the expected middleware wired in.
func New(middleware []Middleware) WalletManagementService {
	var svc WalletManagementService = NewBasicWalletManagementService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
