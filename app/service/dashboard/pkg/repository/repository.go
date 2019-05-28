package repository

import(
  "github.com/jinzhu/gorm"
  "trustkeeper-go/library/database/orm"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

// https://dzone.com/articles/go-microservices-blog-series-part-13-data-consiste
// AccoutRepo account obj
type DashboardRepo struct {
  db *gorm.DB
  // *stdcasbin.Enforcer						// authorization service
}

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

// // Create save repo
// func (repo AccoutRepo) Create(acc *model.Account) error {
//   repo.Enforcer.AddRoleForUser(acc.UUID, "admin")
//   // repo.Enforcer.AddRoleForUser(acc.UUID, "normal")
//   // repo.Enforcer.AddPolicy(acc.UUID, "asset_a", "GET")
//   // repo.Enforcer.AddPermissionForUser(acc.UUID, "read")
//   // repo.Enforcer.Enforce()
//   return repo.db.Create(acc).Error
// }
//
// func (repo AccoutRepo) FindByEmail(email string) (*model.Account, error) {
//   var acc model.Account
//   if err := repo.db.Find(&acc, "email = ?", email).Error; err != nil {
//     return nil, err
//   }
//   return &acc, nil
// }
//
// func (repo AccoutRepo) FindByTokenID(id string) (*model.Account, error) {
//   var acc model.Account
//   if err := repo.db.Find(&acc, "token_id = ?", id).Error; err != nil {
//     return nil, err
//   }
//   return &acc, nil
// }
//
// func (repo AccoutRepo) Update(acc *model.Account, colums map[string]interface{}) error {
//   return repo.db.Model(acc).Update(colums).Error
// }
//
// func (repo AccoutRepo) GetRoles(acc *model.Account) ([]string) {
//   return repo.Enforcer.GetRolesForUser(acc.UUID)
// }