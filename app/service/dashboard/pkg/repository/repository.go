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
  iXpubRepo
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
    model.Namespace{},
    model.Xpub{})
  repo := repo{
    &namespaceRepo{db},
    &groupRepo{db},
    &xpubRepo{db}}
  var biz IBiz = &repo
  return biz
}
