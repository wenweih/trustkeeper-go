package repository

import (
  "trustkeeper-go/library/casbin"
)

type casbinRepo struct {
  *casbin.CasbinRepo
}

type iCasbinRepo interface {
  AddResoucreCreatePolicyForRole(role, domain string, resources ...string)
  HasPolicy(policy []string) bool
  AddReadWriteForRoleInDomain (role, domain, resourceID string)
}

func (repo *casbinRepo) AddResoucreCreatePolicyForRole(role, domain string, resources ...string) {
  for _, r := range resources {
    repo.Enforcer.AddPolicy(role, r, "create")
  }
}

func (repo *casbinRepo) AddReadWriteForRoleInDomain (role, domain, resourceID string) {
  repo.Enforcer.AddPolicy(role, domain, resourceID, "read")
  repo.Enforcer.AddPolicy(role, domain, resourceID, "write")
}

func (repo *casbinRepo) HasPolicy(policy []string) bool {
  return repo.Enforcer.HasPolicy(policy)
}
