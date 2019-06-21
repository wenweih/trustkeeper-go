package jobs

import (
  "github.com/gocraft/work"
  "trustkeeper-go/library/database/redis"
  "trustkeeper-go/Library/common"
  service "trustkeeper-go/app/service/dashboard/pkg/service"
)

type Context struct {
  svc service.JobService
}

// New new jobs
func New(address string, svc service.JobService) *work.WorkerPool {
  redisPool := redis.NewPool(address)
  pool := work.NewWorkerPool(Context{}, 10, redis.Namespace, redisPool)
  pool.Job(common.SignUpJobs, (*Context).Signup)
  // https://github.com/gocraft/work/issues/106
  pool.Middleware(func (c *Context, job *work.Job, next work.NextMiddlewareFunc) error {
    c.svc = svc
    return next()
  })
  return pool
}
