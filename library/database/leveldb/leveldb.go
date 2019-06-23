package leveldb

import (
  "os"
  "fmt"
  "strings"
  homedir "github.com/mitchellh/go-homedir"
  "github.com/syndtr/goleveldb/leveldb"
  "trustkeeper-go/library/util"
)

func New() (*leveldb.DB, error) {
  home, err := homedir.Dir()
  if err != nil {
    return nil, err
  }

  dir := strings.Join([]string{home, common.LevelDBDir}, "/")
  if err := os.MkdirAll(dir, 0755); err != nil {
    return nil, fmt.Errorf("NewLDB %s", err)
  }

  db, err := leveldb.OpenFile(dir, nil)
  if err != nil {
    return nil, fmt.Errorf("NewLDB %s", err)
  }
  return db, nil
}
