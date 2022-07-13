package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-playground/config"
	"go-playground/pkg"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	lockKey              = "key1"
	lockKeyTime          = 5 * time.Second
	inventory            = "inventory"
	inventoryTime        = 10 * time.Minute
	watchDogIntervalTime = time.Second * 3
)

var (
	unlockCh = make(chan int, 0) //用户解锁通知通道
)

type RedisRepo struct {
	client *redis.ClusterClient
}

func NewRedis(config config.RedisCluster) *RedisRepo {
	rdb := pkg.NewRedis(config)
	return &RedisRepo{client: rdb}
}

func (r *RedisRepo) SetInventory(amount int) error {
	var ctx = context.Background()
	result := r.client.Set(ctx, inventory, amount, inventoryTime)
	_, err := result.Result()
	return err
}

func (r *RedisRepo) SubInventory() {
	var ctx = context.Background()
	script := redis.NewScript(`
    if tonumber(redis.call('get', KEYS[1])) > 0 then
		redis.call('decr', KEYS[1])
		return 1 
	end
	return 0
  	`)
	resp := script.Run(ctx, r.client, []string{inventory})
	result, err := resp.Result()
	if err != nil || result == int64(0) {
		fmt.Println("sub inventory failed:", err)
	}
}
func (r *RedisRepo) Lock() {
	var ctx = context.Background()
	for {
		nx := r.client.SetNX(ctx, lockKey, GetCurrentGoroutineId(), lockKeyTime)
		result, err := nx.Result()
		if err == nil && result {
			log.Print("lock", result, err, GetCurrentGoroutineId())
			go r.watchDog(GetCurrentGoroutineId())
			return
		}
	}
}
func (r *RedisRepo) Unlock() {
	var ctx = context.Background()
	nx := r.client.Get(ctx, lockKey)
	value, err := nx.Result()
	if err != nil || value == "" {
		return
	}

	if value != strconv.Itoa(GetCurrentGoroutineId()) {
		return
	}
	time.Sleep(2 * time.Second)
	if value, err := r.client.Get(ctx, lockKey).Result(); err == nil {
		log.Printf("%d ==> %s", GetCurrentGoroutineId(), value)
	}
	// 此版會有解到別人 key 的問題
	result := r.client.Del(ctx, lockKey)
	unlockSuccess, err := result.Result()
	if err == nil && unlockSuccess > 0 {
		log.Println("unlock success! ", GetCurrentGoroutineId())
	} else {
		log.Println("unlock failed!", err)
	}
}

func GetCurrentGoroutineId() int {
	buf := make([]byte, 128)

	buf = buf[:runtime.Stack(buf, false)]

	stackInfo := string(buf)

	goIdStr := strings.TrimSpace(strings.Split(strings.Split(stackInfo, "[running]")[0], "goroutine")[1])

	goId, err := strconv.Atoi(goIdStr)

	if err != nil {

		fmt.Println("err=", err)

		return 0

	}

	return goId
}

func (r *RedisRepo) UnlockUseLua() {
	var ctx = context.Background()
	script := redis.NewScript(`
    if redis.call('get', KEYS[1]) == ARGV[1] then
		redis.call('del', KEYS[1])
		return 1
	else
		return 0
	end
  	`)

	resp := script.Run(ctx, r.client, []string{lockKey}, GetCurrentGoroutineId())
	result, err := resp.Result()
	if err != nil || result == int64(0) {
		fmt.Println("unlock failed:", err, GetCurrentGoroutineId())
	}
	log.Print("unlock", result, err, GetCurrentGoroutineId())
	unlockCh <- 1
}

func (r *RedisRepo) watchDog(goId int) {
	var ctx = context.Background()
	expTicker := time.NewTicker(watchDogIntervalTime)

	script := redis.NewScript(`
		if redis.call('get', KEYS[1]) == ARGV[1] then
      		return redis.call('expire', KEYS[1], ARGV[2])
    	else
      		return 0 
		end
   `)
	for {
		select {
		case <-expTicker.C:
			resp := script.Run(ctx, r.client, []string{lockKey}, goId, lockKeyTime)
			if result, err := resp.Result(); err != nil || result == int64(0) {
				log.Println("expire lock failed", err)
			}
		case <-unlockCh:
			return
		}
	}
}
