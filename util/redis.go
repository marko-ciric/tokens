// Package redis provides ...
package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

func NewRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	}
}
