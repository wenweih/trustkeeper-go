package repository

import (
  "errors"
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/account/pkg/model"
  stdcasbin "github.com/casbin/casbin"
)

type accountRepo struct {
  db *gorm.DB
  *stdcasbin.Enforcer  // authorization service
}

type iAccountRepo interface {
  Create(m *model.Account) error
  Query(query map[string]interface{}) ([]*model.Account, error)
  Update(acc *model.Account, colums map[string]interface{}) error
  Roles(tokenID string) (roles []string, err error)
}

func (repo *accountRepo)Query(query map[string]interface{}) (accounts []*model.Account, err error) {
  err = repo.db.Where(query).Find(&accounts).Error
  if len(accounts) < 1 {
    return nil, errors.New("Empty records")
  }
  return
}

func (repo *accountRepo)Create(m *model.Account) error {
  repo.Enforcer.AddRoleForUser(m.UUID, "admin")
  return repo.db.Create(m).Error
}

func (repo *accountRepo) findByTokenID(id string) (*model.Account, error) {
  var acc model.Account
  if err := repo.db.Find(&acc, "token_id = ?", id).Error; err != nil {
    return nil, err
  }
  return &acc, nil
}

func (repo *accountRepo) Update(acc *model.Account, colums map[string]interface{}) error {
  return repo.db.Model(acc).Update(colums).Error
}

func (repo *accountRepo) GetRoles(acc *model.Account) ([]string) {
  return repo.Enforcer.GetRolesForUser(acc.UUID)
}

func (repo *accountRepo)Roles(tokenID string) ([]string, error) {
  acc, err := repo.findByTokenID(tokenID)
  if err != nil {
    return nil, err
  }
  return repo.Enforcer.GetRolesForUser(acc.UUID), nil
}
