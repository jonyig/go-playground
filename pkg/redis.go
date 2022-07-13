package pkg

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedis(name string, port string, password string, db int) *redis.Client {

	addr := fmt.Sprintf("%s:%s", name, port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return rdb
}
