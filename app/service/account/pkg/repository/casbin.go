package repository

import (
  "trustkeeper-go/library/casbin"
)

type casbinRepo struct {
  *casbin.CasbinRepo
}

type iCasbinRepo interface {
  GetRoles(accountUID, namespaceID string) (roles []string, err error)
  AddResoucreCreatePolicyForRole(role, domain string, resources ...string)
  AddRoleForUserInDomain(accountUID, NamespaceID, role string) (result bool)
}


func (repo *casbinRepo) GetRoles(accountUID, namespaceID string) ([]string, error) {
  roles := repo.Enforcer.GetRolesForUserInDomain(accountUID, namespaceID)
  return roles, nil
}

// https://github.com/casbin/casbin/blob/master/rbac/default-role-manager/role_manager_test.go
func (repo *casbinRepo) AddRoleForUserInDomain(accountUID, NamespaceID, role string) (result bool) {
  // https://github.com/casbin/casbin/blob/master/rbac_api_with_domains.go#L36
  result = repo.Enforcer.AddRoleForUserInDomain(accountUID,
    role,
    NamespaceID)
  return
}

func (repo *casbinRepo) AddResoucreCreatePolicyForRole(role, domain string, resources ...string) {
  for _, r := range resources {
    repo.Enforcer.AddPolicy(role, r, "create")
  }
}
