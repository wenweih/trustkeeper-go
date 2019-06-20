package service

import (
  "context"
)


type JobService interface {
  CreateMnemonic(ctx context.Context, namespaceID string) error
}

func (b *basicWalletManagementService) CreateMnemonic(ctx context.Context, namespaceID string) error {
  // TODO wallet_key service generate xpub
  // xpub, err := b.KeySrv.GenerateMnemonic(ctx, uuid)
  // if err != nil {
  //   return err
  // }
  //
  // return b.biz.Signup(uuid, email, orgName, xpub)
  _, err := b.KeySrv.GenerateMnemonic(ctx, namespaceID)
  return err
}
