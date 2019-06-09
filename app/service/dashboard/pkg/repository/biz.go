package repository

import (
  "trustkeeper-go/app/service/dashboard/pkg/model"
  uuidlib "github.com/satori/go.uuid"
)

type IBiz interface {
  Signup(uuid, email, name, xpub string) error
  Group(m *model.Group) error
}

func (repo *repo) Signup(uuid, email, name, xpub string) error {
  xpubM := &model.Xpub{
    Key: xpub,
    UUID: uuidlib.NewV4().String(),
    Status: true,
  }
  if err := repo.iXpubRepo.Create(xpubM); err != nil {
    return err
  }

  nsM := &model.Namespace{
    Name: name,
    CreatorID: uuid,
    DefaultKey: xpub,
  }
  if err := repo.iNamespaceRepo.Create(nsM); err != nil {
    return err
  }
  return nil
}

func (repo *repo) Group(m *model.Group) error {
  return repo.iGroupRepo.Create(m)
}
