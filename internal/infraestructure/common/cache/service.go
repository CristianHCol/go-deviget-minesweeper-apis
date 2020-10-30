package cache

import "context"

// Cache interface
type Cache interface {
	// Set a value in the cache, if ttl is 0 then the values does not have any expiration
	Set(ctx context.Context, key string, value []byte, ttl uint64) error
	// Get returns the value stored in the cache along with an error if applies
	Get(ctx context.Context, key string) ([]byte, error)
	// Remove removes a value from cache
	Remove(ctx context.Context, key string) bool
}
