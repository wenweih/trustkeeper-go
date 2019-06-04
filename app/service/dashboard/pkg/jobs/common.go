package jobs

import (
  "github.com/gocraft/work"
  "trustkeeper-go/library/database/redis"
)

type Context struct {
  UUID  string
}

func New(address string) *work.WorkerPool {
  redisPool := redis.NewPool(address)
  pool := work.NewWorkerPool(Context{}, 10, redis.Namespace, redisPool)
  pool.Job(groupJobs, (*Context).RootGroup)
  return pool
}
