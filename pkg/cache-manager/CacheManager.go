package cachemanager

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"

	"jukebox-app/pkg/environment"
)

const (
	CACHE_ADDRESS                   = "CACHE_ADDRESS"
	MEMCACHED_ADDRESS_DEFAULT_VALUE = "localhost:11211"
)

type CacheManager interface {
	Get(cacheName string, key any) (any, error)
	Set(cacheName string, key any, value any) error
	Delete(cacheName string, key any) error
	Invalidate(options store.InvalidateOptions) error
	Clear() error
	GetType() string
}

type DefaultCacheManager struct {
	cacheStore store.StoreInterface
	cache      *cache.Cache
	storeType  string
}

func (cacheManager *DefaultCacheManager) Get(cacheName string, key any) (any, error) {
	return nil, nil
}

func (cacheManager *DefaultCacheManager) Set(cacheName string, key any, value any) error {
	var err error
	var valueToCache []byte

	keyToCache := generateKey(cacheName, key)
	if valueToCache, err = json.Marshal(value); err != nil {
		return err
	}

	if err = cacheManager.cache.Set(keyToCache, valueToCache, &store.Options{}); err != nil {
		return err
	}
	return nil
}

func (cacheManager *DefaultCacheManager) Delete(cacheName string, key any) error {
	return nil
}

func (cacheManager *DefaultCacheManager) Invalidate(options store.InvalidateOptions) error {
	return nil
}

func (cacheManager *DefaultCacheManager) Clear() error {
	return nil
}

func (cacheManager *DefaultCacheManager) GetType() string {
	return cacheManager.storeType
}

//

func NewDefaultCacheManager(storeType string, environment environment.Environment) *DefaultCacheManager {

	cacheAddresses := environment.GetValueOrDefault(CACHE_ADDRESS, MEMCACHED_ADDRESS_DEFAULT_VALUE).AsString()
	pair := strings.SplitN(cacheAddresses, ",", 2)

	var cacheStore *store.MemcacheStore
	if storeType == store.MemcacheType {
		cacheStore = store.NewMemcache(memcache.New(pair...), &store.Options{Expiration: 10 * time.Second})
	}

	internalCacheManager := cache.New(cacheStore)

	return &DefaultCacheManager{
		cacheStore: cacheStore,
		cache:      internalCacheManager,
	}
}
