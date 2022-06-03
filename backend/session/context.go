package session

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type CacheContext struct {
	Rdb *redis.Client
	ctx context.Context
}

func NewCacheContext(Rdb *redis.Client, ctx context.Context) *CacheContext {
	return &CacheContext{
		Rdb: Rdb,
		ctx: ctx,
	}
}
