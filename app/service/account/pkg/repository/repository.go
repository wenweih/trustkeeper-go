package repository

import(
  "github.com/jinzhu/gorm"
  "github.com/gomodule/redigo/redis"
  "trustkeeper-go/app/service/account/pkg/model"
  "trustkeeper-go/library/casbin"
)

type repo struct {
  db *gorm.DB
  redisPool *redis.Pool
  iCasbinRepo
  iAccountRepo
  iNamespaceRepo
}

// New new
func New(redisPool *redis.Pool, db *gorm.DB, jwtKey string) IBiz {
  db.AutoMigrate(
    model.Account{},
    model.Namespace{})
  repo := repo{
    db,
    redisPool,
    &casbinRepo{casbin.NewCasbinRepo(db)},
    &accountRepo{},
    &namespaceRepo{}}
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
