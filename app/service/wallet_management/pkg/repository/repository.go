package repository

import(
  "github.com/jinzhu/gorm"
  "github.com/qor/transition"
  "github.com/gomodule/redigo/redis"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
)

// Repo repo obj
type repo struct {
  db *gorm.DB
  redisPool *redis.Pool
  iChainRepo
  iWalletRepo
  iXpubRepo
  imnemonicVersionRepo
}

// New new repo
func New(redisPool *redis.Pool, db *gorm.DB) IBiz {
  db.AutoMigrate(
    model.Wallet{},
    model.Xpub{},
    model.MnemonicVersion{},
    &transition.StateChangeLog{})
  repo := repo{
    db,
    redisPool,
    &chainRepo{},
    &walletRepo{db},
    &xpubRepo{},
    &mnemonicVersionRepo{}}
  var biz IBiz = &repo
  return biz
}

func (repo *repo) close() error {
  if err := repo.db.Close(); err != nil {
    return err
  }
  if err := repo.redisPool.Close(); err != nil {
    return err
  }
  return nil
}
