package repository

import (
  "trustkeeper-go/library/casbin"
)

type casbinRepo struct {
  *casbin.CasbinRepo
}

type iCasbinRepo interface {
  AddResoucreCreatePolicyForRole(role, domain string, resources ...string)
}

func (repo *casbinRepo) AddResoucreCreatePolicyForRole(role, domain string, resources ...string) {
  for _, r := range resources {
    repo.Enforcer.AddPolicy(role, r, "create")
  }
}
