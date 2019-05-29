package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/library/database/orm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

// DashboardRepo account obj
type DashboardRepo struct {
  db *gorm.DB
}

// DB database connect
func DB(dbInfo string) *gorm.DB {
  db, err := orm.Connect(dbInfo)
  if err != nil {
    panic(err.Error())
  }
  return db
}

// New new
func New(db *gorm.DB) DashboardRepo {
  acc := DashboardRepo{db: db}
  acc.db.AutoMigrate(
    model.Group{},
    model.Role{})
  return acc
}
