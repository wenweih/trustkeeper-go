package service

import (
  "context"
)


type JobService interface {
  Signup(ctx context.Context, uuid, email, orgName string) error
}

func (b *basicDashboardService) Signup(ctx context.Context, uuid, email, orgName string) error {
  // TODO wallet_key service generate xpub
  xpub, err := b.walletSrv.GenerateMnemonic(ctx, uuid)
  if err != nil {
    return err
  }

  return b.biz.Signup(uuid, email, orgName, xpub)
}
