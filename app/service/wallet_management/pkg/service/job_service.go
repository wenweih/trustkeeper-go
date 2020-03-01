package service

import (
	"context"
	"trustkeeper-go/app/service/wallet_management/pkg/repository"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/copier"
)

type JobService interface {
	RedisInstance() *redis.Pool
	CreateMnemonic(ctx context.Context, uid, nid string) error
}

func (b *basicWalletManagementService) RedisInstance() *redis.Pool {
	return b.biz.RedisInstance()
}

func (b *basicWalletManagementService) CreateMnemonic(ctx context.Context, uid, nid string) error {
	chains, err := b.biz.GetChains(ctx, map[string]interface{}{})
	if err != nil {
		return err
	}

	// https://huangwenwei.com/blogs/how-to-use-slice-capacity-and-length-in-go
	bip44ids := make([]int32, len(chains))
	for i, chain := range chains {
		bip44ids[i] = int32(chain.Bip44id)
	}
	bip44ThirdXpubsForChains, version, err := b.KeySrv.GenerateMnemonic(ctx, nid, bip44ids, 50)
	if err != nil {
		return err
	}
	localBip44ThirdXpubsForChains := []*repository.Bip44ThirdXpubsForChain{}
	if err := copier.Copy(&localBip44ThirdXpubsForChains, &bip44ThirdXpubsForChains); err != nil {
		return err
	}

	return b.biz.Signup(uid, nid, version, localBip44ThirdXpubsForChains)
}
