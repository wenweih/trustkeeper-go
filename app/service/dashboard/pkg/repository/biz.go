package repository

import (
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

type IBiz interface {
  Signup(uuid, email, name, xpub string) error
  Group(m *model.Group) error
  Close() error
  GetGroups(query map[string]interface{}) (groups []*GetGroupsResp, err error)
}

func (repo *repo) Signup(uuid, email, name, xpub string) error {
  return nil
}

func (repo *repo) Group(m *model.Group) error {
  tx := repo.db.Begin()
  repo.iGroupRepo.Create(tx, m)
  return tx.Commit().Error
}

func (repo *repo) Close() error{
  return repo.close()
}

type GetGroupsResp struct {
	Name  string  `json:"name"`
}

func (repo *repo) GetGroups(query map[string]interface{}) (groupsResp []*GetGroupsResp, err error) {
  groups, err := repo.iGroupRepo.Query(repo.db, query)
  if err != nil {
    return nil, err
  }
  groupsResp = make([]*GetGroupsResp, len(groups))
  for i, group := range groups {
    groupsResp[i] = &GetGroupsResp{Name: group.Name}
  }
  return groupsResp, nil
}
