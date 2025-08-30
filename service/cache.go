package service

import (
	"sync"
	"time"

	"github.com/dgraph-io/ristretto/v2"
)

const (
	cacheCapacity = 1 << 30 // 1 GB
)

var (
	Cache *ristretto.Cache[string, string]
	mut   sync.Mutex
)

func InitCache() {
	tmpCache, err := ristretto.NewCache(&ristretto.Config[string, string]{
		NumCounters: 1e6,           // 追踪 100w key 的频率
		MaxCost:     cacheCapacity, // 1 GB 容量
		BufferItems: 64,            // 推荐值
	})
	if err != nil {
		panic(err)
	}
	Cache = tmpCache
}

func GetFromCache(key string) (string, bool) {
	if Cache == nil {
		return "", false
	}
	if value, found := Cache.Get(key); found {
		return value, true
	}
	return "", false
}

func PutCache(key, value string, ttl time.Duration) {
	if Cache == nil {
		return
	}
	if ttl > 0 {
		Cache.SetWithTTL(key, value, int64(len(value)+len(key)), ttl)
		return
	}
	Cache.Set(key, value, int64(len(value)+len(key)))
}
