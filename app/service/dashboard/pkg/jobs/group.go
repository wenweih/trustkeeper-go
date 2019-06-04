package jobs

import (
  "github.com/gocraft/work"
)

const groupJobs = "group"

func (c *Context)RootGroup(job *work.Job) error {
  return nil
}
