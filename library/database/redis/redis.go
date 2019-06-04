package redis

import (
  "github.com/gomodule/redigo/redis"
)

// Namespace redis namespace
const Namespace = "trustkeeper"

// NewPool redis pool
func NewPool(address string) *redis.Pool {
  return &redis.Pool{
    MaxActive: 5,
    MaxIdle: 5,
    Wait: true,
    Dial: func () (redis.Conn, error) {
      return redis.Dial("tcp", address)
    },
  }
}
