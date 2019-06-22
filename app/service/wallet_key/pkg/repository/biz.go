package repository

import (
  "fmt"
  "time"
  "strings"
)

// IBiz repository bussiness logic
type IBiz interface {
  SaveMnemonic(namespaceID string, mnemonic []byte) (version string, err error)
  Close() error
}

func (repo *repo) Close() error{
  return repo.close()
}

func (repo * repo)SaveMnemonic(namespaceID string, mnemonic []byte) (string ,error) {
  // https://stackoverflow.com/questions/33119748/convert-time-time-to-string
  t := time.Now().Format("2006-01-02 15:04:05.000000")
  version := strings.Join([]string{namespaceID, t}, "/")
  err := repo.ldb.Put([]byte(version), mnemonic, nil)
  if err != nil {
    return "", fmt.Errorf("Save privite key to leveldb %s", err)
  }
  return version, nil
}
