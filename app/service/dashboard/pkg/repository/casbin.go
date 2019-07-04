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
  GetObjForUserInDomain(uid, domain, action string) []interface{}
}

func (repo *casbinRepo) AddResoucreCreatePolicyForRole(role, domain string, resources ...string) {
  for _, r := range resources {
    repo.Enforcer.AddPolicy(role, r, "create")
  }
}

func (repo *casbinRepo) AddReadWriteForRoleInDomain (uid, domain, resourceID string) {
  repo.Enforcer.AddPolicy(uid, domain, resourceID, "read")
  repo.Enforcer.AddPolicy(uid, domain, resourceID, "write")
}

func (repo *casbinRepo) HasPolicy(policy []string) bool {
  return repo.Enforcer.HasPolicy(policy)
}

func (repo *casbinRepo) GetObjForUserInDomain(uid, domain, action string) []interface{} {
  result := repo.Enforcer.GetFilteredPolicy(0, uid, domain)
  matchResult := make([][]string, 0)
  for _, record := range result {
    if record[len(record)-1] == action {
      matchResult = append(matchResult, record)
    }
  }
  objids := make([]interface{}, 0, len(matchResult))
  for _, mr := range matchResult {
    objids = append(objids, mr[len(mr)-2])
  }
  return objids
}
