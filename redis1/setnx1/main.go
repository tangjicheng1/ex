package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient 全局 Redis 客户端
var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// AcquireLock 尝试获取分布式锁
func AcquireLock(ctx context.Context, lockKey string, acquireTimeout, lockTimeout time.Duration) bool {
	endTime := time.Now().Add(acquireTimeout)
	for time.Now().Before(endTime) {
		// 使用 SETNX 尝试获取锁
		set, err := RedisClient.SetNX(ctx, lockKey, "locked", lockTimeout).Result()
		if err != nil {
			fmt.Printf("Error setting lock: %v\n", err)
			return false
		}
		if set {
			return true
		}
		// 短暂休眠后重试
		time.Sleep(100 * time.Millisecond)
	}
	return false
}

// ReleaseLock 释放分布式锁
func ReleaseLock(ctx context.Context, lockKey string) {
	// 删除锁对应的键
	err := RedisClient.Del(ctx, lockKey).Err()
	if err != nil {
		fmt.Printf("Error releasing lock: %v\n", err)
	}
}

func main() {
	ctx := context.Background()
	lockKey := "my_distributed_lock"
	acquireTimeout := 10 * time.Second
	lockTimeout := 10 * time.Second

	// 尝试获取锁
	if AcquireLock(ctx, lockKey, acquireTimeout, lockTimeout) {
		defer ReleaseLock(ctx, lockKey)
		fmt.Println("获取到锁，开始执行临界区代码")
		// 模拟临界区代码执行
		time.Sleep(5 * time.Second)
		fmt.Println("临界区代码执行完毕")
	} else {
		fmt.Println("获取锁失败")
	}
}
