package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/library/database/orm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

// Repo repo obj
type repo struct {
  iNamespaceRepo
  iGroupRepo
}

// DB database connect
func DB(dbInfo string) *gorm.DB {
  db, err := orm.Connect(dbInfo)
  if err != nil {
    panic(err.Error())
  }
  return db
}

// New new repo
func New(db *gorm.DB) IBiz {
  db.AutoMigrate(
    model.Group{},
    model.Role{},
    model.Namespace{})
  repo := repo{
    &namespaceRepo{db},
    &groupRepo{db}}
  var biz IBiz = &repo
  return biz
}
