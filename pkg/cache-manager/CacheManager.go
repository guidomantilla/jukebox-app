package cachemanager

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	gocache "github.com/patrickmn/go-cache"

	"jukebox-app/pkg/environment"
)

const (
	CACHE_ADDRESS                   = "CACHE_ADDRESS"
	MEMCACHED_ADDRESS_DEFAULT_VALUE = "localhost:11211"
	REDIS_ADDRESS_DEFAULT_VALUE     = "localhost:6379"
)

type CacheManager interface {
	Get(ctx context.Context, cacheName string, key any) (any, error)
	Set(ctx context.Context, cacheName string, key any, value any) error
	Delete(ctx context.Context, cacheName string, key any) error
	Invalidate(ctx context.Context, options store.InvalidateOptions) error
	Clear(ctx context.Context) error
	GetType() string
}

type DefaultCacheManager struct {
	cacheStore store.StoreInterface
	cache      *cache.Cache
	storeType  string
}

func (cacheManager *DefaultCacheManager) Get(ctx context.Context, cacheName string, key any) (any, error) {
	return nil, nil
}

func (cacheManager *DefaultCacheManager) Set(ctx context.Context, cacheName string, key any, value any) error {
	var err error
	var valueToCache []byte

	keyToCache := generateKey(cacheName, key)
	if valueToCache, err = json.Marshal(value); err != nil {
		return err
	}

	if err = cacheManager.cache.Set(ctx, keyToCache, valueToCache, &store.Options{}); err != nil {
		return err
	}
	return nil
}

func (cacheManager *DefaultCacheManager) Delete(ctx context.Context, cacheName string, key any) error {
	return nil
}

func (cacheManager *DefaultCacheManager) Invalidate(ctx context.Context, options store.InvalidateOptions) error {
	return nil
}

func (cacheManager *DefaultCacheManager) Clear(ctx context.Context) error {
	return nil
}

func (cacheManager *DefaultCacheManager) GetType() string {
	return cacheManager.storeType
}

//

func NewDefaultCacheManager(storeType string, environment environment.Environment) *DefaultCacheManager {

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
		_ = store.NewRedis(redis.NewClient(&redis.Options{Addr: pair[0]}), &store.Options{Expiration: 10 * time.Second})
		//cacheStore = memcachedStore

	default:
		goCacheStore := store.NewGoCache(gocache.New(5*time.Minute, 10*time.Minute), &store.Options{Expiration: 10 * time.Second})
		cacheStore = goCacheStore
	}

	return &DefaultCacheManager{
		cacheStore: cacheStore,
		cache:      cache.New(cacheStore),
	}
}
