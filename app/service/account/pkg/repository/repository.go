package repository

import(
  // "regexp"
  "github.com/jinzhu/gorm"
  "trustkeeper-go/library/database/orm"
  "trustkeeper-go/app/service/account/pkg/model"
)

// AccoutRepo account obj
type AccoutRepo struct {
  db *gorm.DB
}

func DB(dbInfo string) *gorm.DB {
  db, err := orm.Connect(dbInfo)
  if err != nil {
    panic(err.Error())
  }
  return db
}

// New new
func New(db *gorm.DB) AccoutRepo {
  acc := AccoutRepo{db}
  acc.db.AutoMigrate(
    &model.Account{})

  return acc
}

// Create save repo
func (repo AccoutRepo) Create(acc *model.Account) error {
  return repo.db.Create(acc).Error
  // return repo.db.Raw("INSERT INTO accounts (uuid,email,password) VALUES ($1,$2,$3)",
    // acc.UUID, acc.Email, acc.Password).Error
    // `INSERT INTO "accounts" ("uuid", email, password) VALUES (?,?,?)`,
    // acc.UUID, acc.Email, acc.Password).Error
}
