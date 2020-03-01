package jobs

import (
	service "trustkeeper-go/app/service/wallet_management/pkg/service"
	"trustkeeper-go/library/database/redis"
	common "trustkeeper-go/library/util"

	log "github.com/go-kit/kit/log"
	"github.com/gocraft/work"
)

type Context struct {
	svc service.JobService
}

// New new jobs
func New(logger log.Logger, address string, svc service.JobService) *work.WorkerPool {
	pool := work.NewWorkerPool(Context{}, 10, redis.Namespace, svc.RedisInstance())
	pool.Job(common.WalletMnemonicJob, (*Context).CreateMnemonic)
	// https://github.com/gocraft/work/issues/106
	pool.Middleware(func(c *Context, job *work.Job, next work.NextMiddlewareFunc) error {
		logger.Log("Starting job:", job.Name)
		c.svc = svc
		return next()
	})
	return pool
}
