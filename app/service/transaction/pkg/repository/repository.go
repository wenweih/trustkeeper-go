package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

// Repo repo obj
type repo struct {
  db *gorm.DB
}

// New new repo
func New(db *gorm.DB) IBiz {
  db.AutoMigrate(
    model.Asset{})
  repo := repo{db}
  var biz IBiz = &repo
  return biz
}


func (repo *repo) close() error {
  if err := repo.db.Close(); err != nil {
    return err
  }
  return nil
}
