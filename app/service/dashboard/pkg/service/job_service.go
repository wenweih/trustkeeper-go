package service

import (
  "context"
  "fmt"
)


type JobService interface {
  Signup(ctx context.Context, uuid, email, orgName string) error
}

func (b *basicDashboardService) Signup(ctx context.Context, uuid, email, orgName string) error {
  fmt.Println("uuid: ", uuid, " email: ", email, " orgName: ", orgName)
  return nil
}
