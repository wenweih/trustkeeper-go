package repository

import (
  "trustkeeper-go/app/service/wallet_management/pkg/model"
  "github.com/gomodule/redigo/redis"
)

type IBiz interface {
  // bip44ThirdXpubsForChains 参数是 slice 的引用传递 [] 里面有具体数字是数组的值传递
  // https://leileiluoluo.com/posts/golang-method-calling-value-copy-or-reference-copy.html
  Signup(version string, bip44ThirdXpubsForChains []*Bip44ThirdXpubsForChain) error
  Close() error
  RedisInstance() *redis.Pool
  GetChains() (chains []*model.Chain, err error)
}

type Bip44AccountKey struct {
	Account int     `json:"account"`
	Key     string  `json:"key"`
}

type Bip44ThirdXpubsForChain struct {
	Chain  int           `json:"chain"`
	Xpubs   []*Bip44AccountKey	`json:"xpubs"`
}

func (repo *repo) Signup(version string, bip44ThirdXpubsForChains []*Bip44ThirdXpubsForChain) error {
  tx := repo.db.Begin()
  mnemonicVersion := model.MnemonicVersion{Version: version}
  if err := repo.imnemonicVersionRepo.Create(tx, &mnemonicVersion).Error; err != nil {
    tx.Rollback()
    return err
  }
  for _, bip44ThirdXpubsForChain := range bip44ThirdXpubsForChains {
    for _, xpub := range bip44ThirdXpubsForChain.Xpubs {
      if err := repo.iXpubRepo.Create(tx, &model.Xpub{
        Key: xpub.Key,
        Bip44ChainID: bip44ThirdXpubsForChain.Chain,
        BIP44Account: xpub.Account,
        MnemonicVersionID: mnemonicVersion.ID}).Error; err != nil {
          tx.Rollback()
          return err
        }
    }
  }
  return tx.Commit().Error
}

func (repo *repo)Close() error{
  return repo.close()
}

func (repo *repo) RedisInstance() *redis.Pool {
  return repo.redisPool
}

func (repo *repo)GetChains() (chains []*model.Chain, err error) {
  return repo.iChainRepo.Query(repo.db, map[string]interface{}{})
}
