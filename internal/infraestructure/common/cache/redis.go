package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/infraestructure/common"
)

// RedisCache is the implementation for the cache
type RedisCache struct{}

// NewRedisInstance initializes the redis instance
func NewRedisInstance() *RedisCache {
	return &RedisCache{}
}

// Set set values in the cache
func (s *RedisCache) Set(ctx context.Context, key string, value []byte, ttl uint64) error {
	conn := common.GetRedisConnection(ctx)
	rs := conn.Set(ctx, key, value, time.Second*time.Duration(ttl))
	if rs.Err() != nil {
		fmt.Println("error setting value in redis", rs.Err())
		return rs.Err()
	}

	return nil
}

// Get gets a specific value from the cache
func (s *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	conn := common.GetRedisConnection(ctx)
	rs := conn.Get(ctx, key)
	if rs.Err() != nil {
		fmt.Println("error getting value from redis", rs.Err())
		return nil, rs.Err()
	}

	return rs.Bytes()
}

// Remove removes a specific value from the cache
func (s *RedisCache) Remove(ctx context.Context, key string) bool {
	conn := common.GetRedisConnection(ctx)
	rs := conn.Del(ctx, key)
	if rs.Err() != nil {
		fmt.Println("error removing value from redis", rs.Err())
		return false
	}

	return true
}
