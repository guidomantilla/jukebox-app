package cachemanager

import (
	"fmt"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	gocache "github.com/patrickmn/go-cache"
	"github.com/pkg/errors"

	"jukebox-app/pkg/environment"
)

func generateKey(cacheName string, key any) string {
	return cacheName + "-" + fmt.Sprintf("%v", key)
}

func BuildCacheInterface(storeType string, environment environment.Environment) (cache.CacheInterface, error) {

	var cacheStore store.StoreInterface
	switch storeType {
	case store.MemcacheType:
		cacheAddresses := environment.GetValueOrDefault(CACHE_ADDRESS, MEMCACHED_ADDRESS_DEFAULT_VALUE).AsString()
		pair := strings.SplitN(cacheAddresses, ",", 2)
		memcachedStore := store.NewMemcache(memcache.New(pair...), &store.Options{Expiration: 10 * time.Second})
		cacheStore = memcachedStore

	case store.RedisType:
		cacheAddresses := environment.GetValueOrDefault(CACHE_ADDRESS, REDIS_ADDRESS_DEFAULT_VALUE).AsString()
		pair := strings.SplitN(cacheAddresses, ",", 2)
		redisStore := store.NewRedis(redis.NewClient(&redis.Options{Addr: pair[0]}), &store.Options{Expiration: 10 * time.Second})
		cacheStore = redisStore

	case store.GoCacheType:
		goCacheStore := store.NewGoCache(gocache.New(5*time.Minute, 10*time.Minute), &store.Options{Expiration: 10 * time.Second})
		cacheStore = goCacheStore

	default:
		return nil, errors.New("storeType not defined")
	}

	return cache.New(cacheStore), nil
}
