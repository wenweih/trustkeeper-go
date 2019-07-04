package repository

import (
  "fmt"
  "context"
  "strconv"
  "google.golang.org/grpc/metadata"
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

type IBiz interface {
  Signup(uuid, email, name, xpub string) error
  CreateGroup(ctx context.Context, m *model.Group) error
  Close() error
  GetGroups(ctx context.Context, query map[string]interface{}) (groups []*GetGroupsResp, err error)
}

func (repo *repo) Signup(uuid, email, name, xpub string) error {
  return nil
}

func (repo *repo) CreateGroup(ctx context.Context, m *model.Group) error {
  md, ok := metadata.FromIncomingContext(ctx)
  if !ok {
    return fmt.Errorf("fail to extract auth info from ctx")
  }
  allowRoles := make([]string, 0, len(md["roles"]))
  for _, role := range md["roles"] {
   allow := repo.iCasbinRepo.HasPolicy([]string{role, "group", "create"})
   if allow {
     allowRoles = append(allowRoles, role)
   }
  }

  if len(allowRoles) <= 0 {
    return fmt.Errorf("not allow")
  }

  tx := repo.db.Begin()
  if err := repo.iGroupRepo.Create(tx, m).Error; err != nil {
    tx.Rollback()
    return err
  }
  if err := tx.Commit().Error; err != nil {
    return err
  }
  repo.iCasbinRepo.AddReadWriteForRoleInDomain(m.CreatorID, m.NamespaceID, strconv.FormatUint(uint64(m.ID), 10))
  return nil
}

func (repo *repo) Close() error{
  return repo.close()
}

type GetGroupsResp struct {
  ID    string  `json:"id"`
	Name  string  `json:"name"`
  Desc  string  `json:"desc"`
}

func (repo *repo) GetGroups(ctx context.Context, query map[string]interface{}) (groupsResp []*GetGroupsResp, err error) {
  md, ok := metadata.FromIncomingContext(ctx)
  if !ok {
    return nil, fmt.Errorf("fail to extract auth info from ctx")
  }
  uid := md["uid"][0]
  nid := md["nid"][0]

  ids := repo.iCasbinRepo.GetObjForUserInDomain(uid, nid, "read")
  groups, err := repo.iGroupRepo.Query(repo.db, ids, query)
  if err != nil {
    return nil, err
  }
  groupsResp = make([]*GetGroupsResp, len(groups))
  for i, group := range groups {
    groupsResp[i] = &GetGroupsResp{Name: group.Name, Desc: group.Desc, ID: strconv.FormatUint(uint64(group.ID), 10)}
  }
  return groupsResp, nil
}
