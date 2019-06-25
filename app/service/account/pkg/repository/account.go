package repository

import (
  "strconv"
  "context"
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
  Query(ctx context.Context, tx *gorm.DB, query *model.Account) ([]*model.Account, error)
  Update(tx *gorm.DB, acc *model.Account, colums map[string]interface{}) error
  Roles(tx *gorm.DB, account *model.Account) (roles []string, err error)
  AddRoleForUserInDomain(accountUID, NamespaceID, role string) (result bool)
}

func (repo *accountRepo) Query(ctx context.Context, tx *gorm.DB, query *model.Account) (accounts []*model.Account, err error) {
  result := tx.Set("gorm:auto_preload", true).Find(&accounts, query)
  if result.Error != nil {
    return nil, errors.New("Account Query error: " + err.Error())
  }
  if len(accounts) < 1 {
    return nil, errors.New("Empty records")
  }
  return accounts, nil
}

func (repo *accountRepo)Create(tx *gorm.DB, m *model.Account) *gorm.DB {
  return tx.Create(m)
}

func (repo *accountRepo) Update(tx *gorm.DB, acc *model.Account, colums map[string]interface{}) error {
  return tx.Model(acc).Update(colums).Error
}

// https://github.com/casbin/casbin/blob/master/rbac/default-role-manager/role_manager_test.go
func (repo *accountRepo) AddRoleForUserInDomain(accountUID, NamespaceID, role string) (result bool) {
  // https://github.com/casbin/casbin/blob/master/rbac_api_with_domains.go#L36
  result = repo.Enforcer.AddRoleForUserInDomain(accountUID,
    role,
    NamespaceID)
  return
}

func (repo *accountRepo) Roles(tx *gorm.DB, account *model.Account) ([]string, error) {
  namespaceID := strconv.FormatUint(uint64(account.Namespace.ID), 10)
  roles := repo.Enforcer.GetRolesForUserInDomain(account.UUID, namespaceID)
  return roles, nil
}
