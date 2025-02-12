package middlewares

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func CacheMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := "cache:" + c.Request().RequestURI
		ctx := context.Background()

		cachedResponse, err := rdb.Get(ctx, key).Result()
		if err == nil {
			c.Response().WriteHeader(http.StatusOK)
			c.Response().Write([]byte(cachedResponse))
			return nil
		}

		responseWriter := &responseCacheWriter{ResponseWriter: c.Response().Writer, buffer: new(bytes.Buffer)}
		c.Response().Writer = responseWriter

		err = next(c)
		if err != nil {
			return err
		}

		rdb.Set(ctx, key, responseWriter.buffer.String(), 10*time.Minute)
		return nil
	}
}

// ✅ Correcte définition de responseCacheWriter
type responseCacheWriter struct {
	http.ResponseWriter
	buffer *bytes.Buffer
}

func (w *responseCacheWriter) Write(b []byte) (int, error) {
	w.buffer.Write(b)
	return w.ResponseWriter.Write(b)
}
