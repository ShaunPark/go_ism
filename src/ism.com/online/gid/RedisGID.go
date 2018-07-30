package gid

import (
	"fmt"

	"github.com/go-redis/redis"
	"ism.com/online/ismredis"
)

type RedisChecker struct {
	GidCheckerInterface
}

func (gidChecker *RedisChecker) CheckGID(gid string) bool {
	println("RedidChecker ...")
	_, err := ismredis.Get(fmt.Sprint("GID:", gid))
	if err != nil {

		if err == redis.Nil {
			ismredis.SetExpire(fmt.Sprint("GID:", gid), gid, 10)
			println("return true")
			return true
		} else {
			panic(err.Error())
		}
	}

	println("return false")
	return false
}
