package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

// Repo repo obj
type repo struct {
  iGroupRepo
}

// New new repo
func New(db *gorm.DB) IBiz {
  db.AutoMigrate(
    model.Group{})
  repo := repo{
    &groupRepo{db}}
  var biz IBiz = &repo
  return biz
}
