package jobs

import (
  "context"
  "github.com/gocraft/work"
)

func (c *Context) CreateMnemonic(job *work.Job) error {
  namespaceid := job.ArgString("namespaceid")
  return c.svc.CreateMnemonic(context.Background(), namespaceid)
}
