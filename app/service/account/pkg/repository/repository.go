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
    model.Account{})
  return acc
}

// Create save repo
func (repo AccoutRepo) Create(acc *model.Account) error {
  return repo.db.Create(acc).Error
}

func (repo AccoutRepo) FindByEmail(email string) (*model.Account, error) {
  var acc model.Account
  if err := repo.db.Find(&acc, "email = ?", email).Error; err != nil {
    return nil, err
  }
  return &acc, nil
}

func (repo AccoutRepo) Update(acc *model.Account, colums map[string]interface{}) error {
  return repo.db.Model(acc).Update(colums).Error
}
