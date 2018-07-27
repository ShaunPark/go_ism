package gid

import (
  "fmt"
  "ism.com/online/ismredis"
  "github.com/go-redis/redis"
)
type RedisChecker struct {
  GidCheckerInterface
}

func (gidChecker *RedisChecker) CheckGID(gid string) bool {
  _, err := ismredis.Get(fmt.Sprint("GID:",gid))
  if err == redis.Nil {
    ismredis.Set(fmt.Sprint("GID:",gid), gid)
    return false
  }

  return true
}
