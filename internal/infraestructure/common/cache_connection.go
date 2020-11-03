package common

import (
	"context"
	"net/url"
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
		var resolvedURL = os.Getenv("CACHE_URI")
		var password = ""

		parsedURL, _ := url.Parse(resolvedURL)
		password, _ = parsedURL.User.Password()
		resolvedURL = parsedURL.Host

		redisConnection = redis.NewClient(&redis.Options{
			Addr:         resolvedURL,
			PoolSize:     redisPoolSize,
			MinIdleConns: redisIdleConn,
			Password:     password,
		})
	}

	return redisConnection.Conn(ctx)
}
