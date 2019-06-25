package repository

import (
  "fmt"
  "strconv"
  "errors"
  "github.com/jinzhu/gorm"
  "trustkeeper-go/app/service/account/pkg/model"
  stdcasbin "github.com/casbin/casbin"
)

type accountRepo struct {
  *stdcasbin.Enforcer  // authorization service
}

type iAccountRepo interface {
  Create(tx *gorm.DB, m *model.Account) *gorm.DB
  Query(tx *gorm.DB, query map[string]interface{}) ([]*model.Account, error)
  Update(tx *gorm.DB, acc *model.Account, colums map[string]interface{}) error
  Roles(tx *gorm.DB, tokenID string) (roles []string, err error)
  AddRoleForUserInDomain(accountUID, NamespaceID, role string) (result bool)
}

func (repo *accountRepo)Query(tx *gorm.DB, query map[string]interface{}) (accounts []*model.Account, err error) {
  err = tx.Preload("Namespace").Where(query).Find(&accounts).Error
  if len(accounts) < 1 {
    return nil, errors.New("Empty records")
  }
  return
}

func (repo *accountRepo)Create(tx *gorm.DB, m *model.Account) *gorm.DB {
  return tx.Create(m)
}

func (repo *accountRepo) findByTokenID(tx *gorm.DB, id string) (*model.Account, error) {
  var acc model.Account
  if err := tx.Preload("Namespace").Find(&acc, "token_id = ?", id).Error; err != nil {
    return nil, err
  }
  return &acc, nil
}

func (repo *accountRepo) Update(tx *gorm.DB, acc *model.Account, colums map[string]interface{}) error {
  return tx.Model(acc).Update(colums).Error
}

func (repo *accountRepo) GetRoles(tx *gorm.DB, acc *model.Account) (roles []string) {
  roles = repo.Enforcer.GetRolesForUserInDomain(acc.UUID, strconv.FormatUint(uint64(acc.Namespace.ID), 10))
  fmt.Println("roles: ", roles, " namespace: ", acc.Namespace.ID)
  return
  // return repo.Enforcer.GetRolesForUser(acc.UUID)
}

// https://github.com/casbin/casbin/blob/master/rbac/default-role-manager/role_manager_test.go
func (repo *accountRepo) AddRoleForUserInDomain(accountUID, NamespaceID, role string) (result bool) {
  // https://github.com/casbin/casbin/blob/master/rbac_api_with_domains.go#L36
  result = repo.Enforcer.AddRoleForUserInDomain(accountUID,
    role,
    NamespaceID)
  return
}

func (repo *accountRepo)Roles(tx *gorm.DB, tokenID string) ([]string, error) {
  acc, err := repo.findByTokenID(tx, tokenID)
  if err != nil {
    return nil, err
  }
  roles := repo.Enforcer.GetRolesForUserInDomain(acc.UUID, strconv.FormatUint(uint64(acc.Namespace.ID), 10))
  fmt.Println("roles: ", roles, " namespace: ", acc.Namespace.ID)
  return roles, nil
}
