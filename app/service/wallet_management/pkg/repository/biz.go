package repository

import (
  // "trustkeeper-go/app/service/wallet_management/pkg/model"
  // uuidlib "github.com/satori/go.uuid"
)

type IBiz interface {
  Signup(key string, chainID, groupID int8) error
}

func (repo *repo) Signup(key string, chainID, groupID int8) error {
  return nil
}
