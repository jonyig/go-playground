package pkg

import (
	"github.com/go-redis/redis/v8"
	"go-playground/config"
	"time"
)

func NewRedis(config config.RedisCluster) *redis.ClusterClient {

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    config.Addrs,
		Password: config.Password,
		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		DialTimeout:   time.Duration(config.DialTimeout) * time.Second,  // 设置连接超时
		ReadTimeout:   time.Duration(config.ReadTimeout) * time.Second,  // 设置读取超时
		WriteTimeout:  time.Duration(config.WriteTimeout) * time.Second, // 设置写入超时
		RouteRandomly: true,
	})

	return rdb
}
