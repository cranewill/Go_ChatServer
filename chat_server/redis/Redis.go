package redis

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

var Redisdb *redis.Client

func init() {
	log.Println("Init Redis ...")
	Redisdb = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

// IsNil returns if the result of Redisdb 'Get' operation is Nil
func IsNil(err error) bool {
	return err == redis.Nil
}
