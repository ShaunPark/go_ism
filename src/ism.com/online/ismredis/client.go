package ismredis

import (
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func createRedisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
}

func Get(key string) (string, error) {
	if client == nil {
		createRedisClient()
	}
	val, err := client.Get(key).Result()
	return val, err
}

func Set(key string, val string) {
	if client == nil {
		createRedisClient()
	}
	client.Set(key, val, 0)
}

func SetExpire(key string, val string, sec int) {
	if client == nil {
		createRedisClient()
	}
	client.Set(key, val, time.Duration(sec)*time.Second)
}
