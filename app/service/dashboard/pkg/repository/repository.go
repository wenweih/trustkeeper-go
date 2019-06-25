package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

// Repo repo obj
type repo struct {
  db *gorm.DB
  // redisPool *redis.Pool
  iGroupRepo
}

// New new repo
func New(db *gorm.DB) IBiz {
  db.AutoMigrate(
    model.Group{})
  repo := repo{
    db,
    &groupRepo{}}
  var biz IBiz = &repo
  return biz
}


func (repo *repo) close() error {
  if err := repo.db.Close(); err != nil {
    return err
  }
  // if err := repo.redisPool.Close(); err != nil {
  //   return err
  // }
  return nil
}
