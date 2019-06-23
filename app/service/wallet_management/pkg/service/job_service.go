package service

import (
  "context"
  "github.com/jinzhu/copier"
  "github.com/gomodule/redigo/redis"
  "trustkeeper-go/app/service/wallet_management/pkg/repository"
)

type JobService interface {
  RedisInstance() *redis.Pool
  CreateMnemonic(ctx context.Context, namespaceID string) error
}

func (b *basicWalletManagementService) RedisInstance() *redis.Pool {
  return b.biz.RedisInstance()
}

func (b *basicWalletManagementService) CreateMnemonic(ctx context.Context, namespaceID string) error {
  chains, err := b.biz.GetChains()
  if err != nil {
    return err
  }

  // https://huangwenwei.com/blogs/how-to-use-slice-capacity-and-length-in-go
  bip44ids := make([]int32, len(chains))
  for i, chain := range chains {
    bip44ids[i] = int32(chain.Bip44id)
  }
  bip44ThirdXpubsForChains, version, err := b.KeySrv.GenerateMnemonic(ctx, namespaceID, bip44ids, 10)
  if err != nil {
    return err
  }
  localBip44ThirdXpubsForChains := []*repository.Bip44ThirdXpubsForChain{}
  copier.Copy(&localBip44ThirdXpubsForChains, &bip44ThirdXpubsForChains)

  return b.biz.Signup(version, localBip44ThirdXpubsForChains)
}
