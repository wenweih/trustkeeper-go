package repository

import (
  // "trustkeeper-go/app/service/wallet_management/pkg/model"
  // uuidlib "github.com/satori/go.uuid"
  "github.com/gomodule/redigo/redis"
)

type IBiz interface {
  Signup(key string, chainID, groupID int8) error
  Close() error
  RedisInstance() *redis.Pool
}

func (repo *repo) Signup(key string, chainID, groupID int8) error {
  return nil
}

func (repo *repo)Close() error{
  return repo.close()
}

func (repo *repo) RedisInstance() *redis.Pool {
  return repo.redisPool
}
