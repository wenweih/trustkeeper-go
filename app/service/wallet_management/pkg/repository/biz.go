package repository

import (
  "trustkeeper-go/app/service/wallet_management/pkg/model"
  "github.com/gomodule/redigo/redis"
)

type IBiz interface {
  Signup(key string, chainID, groupID int8) error
  Close() error
  RedisInstance() *redis.Pool
  GetChains() (chains []*model.Chain, err error)
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

func (repo *repo)GetChains() (chains []*model.Chain, err error) {
  return repo.iChainRepo.Query(repo.db, map[string]interface{}{})
}
