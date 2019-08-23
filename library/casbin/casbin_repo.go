package casbin

type casbinRepo struct {
  *CasbinRepo
}

type ICasbinRepo interface {
  AddResoucreCreatePolicyForRole(role, domain string, resources ...string)
  HasPolicy(policy []string) bool
  AddActionForRoleInDomain (uid, domain, resourceID, action string)
  GetObjForUserInDomain(uid, domain, action string) []interface{}
}

func (repo *CasbinRepo) AddResoucreCreatePolicyForRole(role, domain string, resources ...string) {
  for _, r := range resources {
    repo.Enforcer.AddPolicy(role, r, "create")
  }
}

func (repo *CasbinRepo) AddActionForRoleInDomain (uid, domain, resourceID, action string) {
  repo.Enforcer.AddPolicy(uid, domain, resourceID, action)
}

func (repo *CasbinRepo) HasPolicy(policy []string) bool {
  repo.Enforcer.LoadPolicy()
  return repo.Enforcer.HasPolicy(policy)
}

func (repo *CasbinRepo) GetObjForUserInDomain(uid, domain, action string) []interface{} {
  repo.Enforcer.LoadPolicy()
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
