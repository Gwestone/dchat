package session

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type StorageContext struct {
	Rdb *redis.Client
	ctx context.Context
}

func NewStorageContext(Rdb *redis.Client, ctx context.Context) *StorageContext {
	return &StorageContext{
		Rdb: Rdb,
		ctx: ctx,
	}
}
