package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
}

func GetFromCache(key string) (string, error) {
	ctx := context.Background()
	return RedisClient.Get(ctx, key).Result()
}

func SetToCache(key, value string) error {
	ctx := context.Background()
	return RedisClient.Set(ctx, key, value, 0).Err()
}
