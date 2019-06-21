package repository

import (
  "github.com/syndtr/goleveldb/leveldb"
  libLeveldb "trustkeeper-go/Library/database/leveldb"
)

type repo struct {
  ldb *leveldb.DB
}

// New new
func New() (IBiz, error) {
  ldb, err := libLeveldb.New()
  if err != nil {
    return nil, err
  }
  repo := repo{
    ldb: ldb,
  }
  var biz IBiz = &repo
  return biz, nil
}

func (repo *repo) close() error {
  if err := repo.ldb.Close(); err != nil {
    return err
  }
  return nil
}
