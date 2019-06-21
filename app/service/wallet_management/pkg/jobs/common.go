package jobs

import (
  "github.com/gocraft/work"
  "trustkeeper-go/Library/database/redis"
  "trustkeeper-go/Library/common"
  service "trustkeeper-go/app/service/wallet_management/pkg/service"
)

type Context struct {
  svc service.JobService
}

// New new jobs
func New(address string, svc service.JobService) *work.WorkerPool {
  pool := work.NewWorkerPool(Context{}, 10, redis.Namespace, svc.RedisInstance())
  pool.Job(common.WalletMnemonicJob, (*Context).CreateMnemonic)
  // https://github.com/gocraft/work/issues/106
  pool.Middleware(func (c *Context, job *work.Job, next work.NextMiddlewareFunc) error {
    c.svc = svc
    return next()
  })
  return pool
}
