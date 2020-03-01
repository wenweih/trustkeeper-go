package jobs

import (
	"context"
	"fmt"

	"github.com/gocraft/work"
)

// CreateMnemonic create mnemonic job
func (c *Context) CreateMnemonic(job *work.Job) error {
	nid := job.ArgString("namespaceid")
	uid := job.ArgString("uid")
	if err := job.ArgError(); err != nil {
		fmt.Println("arg error", err.Error())
		return err
	}
	err := c.svc.CreateMnemonic(context.Background(), uid, nid)
	if err != nil {
		fmt.Println("CreateMnemonic err: ", err.Error())
		return err
	}
	return nil
}
