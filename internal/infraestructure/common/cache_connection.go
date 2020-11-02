package common

import (
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var redisConnection *redis.Client

var redisPoolSize, _ = strconv.Atoi(os.Getenv("CACHE_POOL"))
var redisIdleConn, _ = strconv.Atoi(os.Getenv("CACHE_IDLE_CONN"))

// GetRedisConnection returns a redis connection
func GetRedisConnection(ctx context.Context) *redis.Conn {
	if redisConnection == nil {
		redisConnection = redis.NewClient(&redis.Options{
			Addr:         os.Getenv("CACHE_URI"),
			PoolSize:     redisPoolSize,
			MinIdleConns: redisIdleConn,
			Password:     os.Getenv("CACHE_PASSWORD"),
		})
	}

	return redisConnection.Conn(ctx)
}
