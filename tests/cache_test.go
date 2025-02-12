package tests

import (
	"testing"
	"time"

	"github.com/idir-44/mt5-cdn-project/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestRedisCache(t *testing.T) {
	storage.ConnectRedis()
	err := storage.SetCache("test_key", "test_value", 10*time.Second)
	assert.NoError(t, err)

	value, err := storage.GetCache("test_key")
	assert.NoError(t, err)
	assert.Equal(t, "test_value", value)
}
