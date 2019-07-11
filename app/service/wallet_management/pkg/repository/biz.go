package repository

import (
  "fmt"
  "strconv"
  "context"
  "google.golang.org/grpc/metadata"
  "github.com/gomodule/redigo/redis"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
  libctx "trustkeeper-go/library/context"
)

type IBiz interface {
  // bip44ThirdXpubsForChains 参数是 slice 的引用传递 [] 里面有具体数字是数组的值传递
  // https://leileiluoluo.com/posts/golang-method-calling-value-copy-or-reference-copy.html
  Signup(version string, bip44ThirdXpubsForChains []*Bip44ThirdXpubsForChain) error
  Close() error
  RedisInstance() *redis.Pool
  GetChains(ctx context.Context, query map[string]interface{}) (chains []*SimpleChain, err error)
  UpdateXpubState(ctx context.Context, from, to, groupid string) error
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

func (repo *repo)GetChains(ctx context.Context, query map[string]interface{}) ([]*SimpleChain, error) {
  _, _, _, err := libctx.ExtractAuthInfoFromContext(ctx)
  if err != nil {
    return nil, err
  }
  chains, err := repo.iChainRepo.Query(repo.db, query)
  if err != nil {
    return nil, err
  }
  simpleChains := make([]*SimpleChain, len(chains))
  for i, c := range chains {
    simpleChains[i] = &SimpleChain{
      ID: strconv.FormatUint(uint64(c.ID), 10),
      Name: c.Name,
      Coin: c.Coin,
      Bip44id: c.Bip44id,
      Status: c.Status}
  }
  return simpleChains, nil
}

func (repo *repo) UpdateXpubState(ctx context.Context, from, to, groupid string) error {
  md, ok := metadata.FromIncomingContext(ctx)
  if !ok {
    return fmt.Errorf("fail to extract auth info from ctx")
  }
  // uid := md["uid"][0]
  nid := md["nid"][0]

  mnemonicVs, err := repo.imnemonicVersionRepo.VersionLikeQuery(repo.db, nid)
  if err != nil {
    return err
  }

  if len(mnemonicVs) != 1 {
    return fmt.Errorf("records error")
  }
  xpub := model.Xpub{}
  switch from {
  case Idle:
    repo.db.Where("state = ? AND mnemonic_version_id = ?", Idle, uint(mnemonicVs[0].ID)).First(&xpub)
  case Assigned:
    repo.db.Where("state = ? AND mnemonic_version_id = ?", Assigned, uint(mnemonicVs[0].ID)).First(&xpub)
  case Abandon:
    repo.db.Where("state = ? AND mnemonic_version_id = ?", Abandon, uint(mnemonicVs[0].ID)).First(&xpub)
  default:
    return fmt.Errorf("invalid state:" + from)
  }
  if err := repo.iXpubRepo.UpdateState(repo.db, &xpub, to); err != nil {
    return err
  }
  if err := repo.db.Model(&xpub).Updates(map[string]interface{}{"state": to, "group_id": groupid}).Error; err != nil {
    return err
  }
  return nil
}
