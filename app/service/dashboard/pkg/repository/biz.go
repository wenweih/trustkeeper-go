package repository

import (
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

type IBiz interface {
  Signup(uuid, email, name, xpub string) error
  Group(m *model.Group) error
}

func (repo *repo) Signup(uuid, email, name, xpub string) error {
  return nil
}

func (repo *repo) Group(m *model.Group) error {
  return repo.iGroupRepo.Create(m)
}
