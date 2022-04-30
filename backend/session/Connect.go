package session

import (
	"github.com/go-redis/redis/v8"
)

func Connect(addr, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	return rdb
}
