package cachemanager

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	gocache "github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

func generateKey(cacheName string, key any) string {
	return cacheName + "-" + fmt.Sprintf("%v", key)
}

func BuildCacheInterface(storeType string, cacheAddresses ...string) (cache.CacheInterface, error) {

	var cacheStore store.StoreInterface
	switch storeType {
	case store.MemcacheType:
		memcachedStore := store.NewMemcache(memcache.New(cacheAddresses...), &store.Options{Expiration: 10 * time.Second})
		cacheStore = memcachedStore

	case store.RedisType:
		redisStore := store.NewRedis(redis.NewClient(&redis.Options{Addr: cacheAddresses[0]}), &store.Options{Expiration: 10 * time.Second})
		cacheStore = redisStore

	case store.GoCacheType:
		goCacheStore := store.NewGoCache(gocache.New(5*time.Minute, 10*time.Minute), &store.Options{Expiration: 10 * time.Second})
		cacheStore = goCacheStore

	default:
		return nil, errors.New("storeType not defined")
	}

	return cache.New(cacheStore), nil
}
