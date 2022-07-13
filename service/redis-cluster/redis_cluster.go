package redis_cluster

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func Start() {

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"docker-own_redis-node-0_1:6379",
			"docker-own_redis-node-1_1:6379",
			"docker-own_redis-node-2_1:6379",
			"docker-own_redis-node-3_1:6379",
			"docker-own_redis-node-4_1:6379",
			"docker-own_redis-node-5_1:6379",
		},
		Password: "jonny",
		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		DialTimeout:   50 * time.Second, // 设置连接超时
		ReadTimeout:   50 * time.Second, // 设置读取超时
		WriteTimeout:  50 * time.Second, // 设置写入超时
		RouteRandomly: true,
	})
	rdb.Set(context.Background(), "k2", "123", -1)
	result := rdb.Get(context.Background(), "k2")
	log.Print(result)

}
