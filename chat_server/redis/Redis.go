package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var Redisdb *redis.Client

func init() {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	fmt.Println("Redis init!")
}

// IsNil returns if the result of Redisdb 'Get' operation is Nil
func IsNil(err error) bool {
	return err == redis.Nil
}


