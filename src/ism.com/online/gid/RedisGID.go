package gid

import (
	"fmt"

	"github.com/go-redis/redis"
	"ism.com/common/rediscache"
)

type RedisChecker struct {
	GidCheckerInterface
}

func (gidChecker *RedisChecker) CheckGID(gid string) bool {
	println("RedidChecker ...")
	_, err := rediscache.Get(fmt.Sprint("GID:", gid))
	if err != nil {

		if err == redis.Nil {
			rediscache.SetExpire(fmt.Sprint("GID:", gid), gid, 10)
			println("return true")
			return true
		} else {
			panic(err.Error())
		}
	}

	println("return false")
	return false
}
