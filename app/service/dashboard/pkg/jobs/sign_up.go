package jobs

import (
	"context"

	"github.com/gocraft/work"
)

func (c *Context) Signup(job *work.Job) error {
	uuid := job.ArgString("uuid")
	email := job.ArgString("email")
	orgname := job.ArgString("orgname")
	return c.svc.Signup(context.Background(), uuid, email, orgname)
}
