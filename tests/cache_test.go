package tests

import (
	"context"
	"testing"
	"time"

	"github.com/idir-44/mt5-cdn-project/internal/storage"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisCache(t *testing.T) {
	storage.RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	err := storage.RedisClient.Set(context.Background(), "test_key", "test_value", 10*time.Second).Err()
	assert.NoError(t, err)

	value, err := storage.RedisClient.Get(context.Background(), "test_key").Result()
	assert.NoError(t, err)
	assert.Equal(t, "test_value", value)
}
