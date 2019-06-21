package service

import (
  "context"
  "github.com/gomodule/redigo/redis"
)

type JobService interface {
  RedisInstance() *redis.Pool
  CreateMnemonic(ctx context.Context, namespaceID string) error
}

func (b *basicWalletManagementService) RedisInstance() *redis.Pool {
  return b.biz.RedisInstance()
}

func (b *basicWalletManagementService) CreateMnemonic(ctx context.Context, namespaceID string) error {
  // TODO wallet_key service generate xpub
  // xpub, err := b.KeySrv.GenerateMnemonic(ctx, uuid)
  // if err != nil {
  //   return err
  // }
  //
  // return b.biz.Signup(uuid, email, orgName, xpub)
  // params: namespaceid as levelDB key, slices for default bip44id of chain, default size for response xpub
  // response:
  // [
  //   {
  //     "chainid": xxx,
  //     [
  //       {
  //         "bip44account": xxx,
  //         "key": xxx
  //       }
  //       ...
  //     ]
  //   },
  //   ...
  // ]
  // _, err := b.KeySrv.GenerateMnemonic(ctx, namespaceID)
  return nil
}
