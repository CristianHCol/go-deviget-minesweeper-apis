package netcommon

import (
	"math/rand"
	"time"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/infrastructure/common/cache"
)

var memoryCache cache.Cache = cache.NewRedisInstance()
var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString() string {
	rand.NewSource(time.Now().UnixNano())
	b := make([]byte, 35)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))] //nolint:gosec
	}
	return string(b)
}

func setCacheInstance(ch cache.Cache) {
	memoryCache = ch
}
