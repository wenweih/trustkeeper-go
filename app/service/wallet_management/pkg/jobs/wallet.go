package jobs

import (
  "fmt"
  "context"
  "github.com/gocraft/work"
)

// CreateMnemonic create mnemonic job
func (c *Context) CreateMnemonic(job *work.Job) error {
  namespaceid := job.ArgString("namespaceid")
  if err := job.ArgError(); err != nil {
		return err
	}
  err := c.svc.CreateMnemonic(context.Background(), namespaceid)
  if err != nil {
    fmt.Println("CreateMnemonic err: ", err.Error())
    return err
  }
  return nil
}
