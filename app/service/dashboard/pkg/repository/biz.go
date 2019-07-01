package repository

import (
  "strconv"
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
  if err := repo.iGroupRepo.Create(tx, m).Error; err != nil {
    tx.Rollback()
    return err
  }
  return tx.Commit().Error
}

func (repo *repo) Close() error{
  return repo.close()
}

type GetGroupsResp struct {
  ID    string  `json:"id"`
	Name  string  `json:"name"`
  Desc  string  `json:"desc"`
}

func (repo *repo) GetGroups(query map[string]interface{}) (groupsResp []*GetGroupsResp, err error) {
  groups, err := repo.iGroupRepo.Query(repo.db, query)
  if err != nil {
    return nil, err
  }
  groupsResp = make([]*GetGroupsResp, len(groups))
  for i, group := range groups {
    groupsResp[i] = &GetGroupsResp{Name: group.Name, Desc: group.Desc, ID: strconv.FormatUint(uint64(group.ID), 10)}
  }
  return groupsResp, nil
}
