package config

import (
	"fmt"
	"strings"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"go.uber.org/zap"

	"jukebox-app/pkg/cachemanager"
	"jukebox-app/pkg/environment"
)

const (
	CACHE_ADDRESS                   = "CACHE_ADDRESS"
	CACHE_STORE_TYPE                = "CACHE_STORE_TYPE"
	MEMCACHED_ADDRESS_DEFAULT_VALUE = "localhost:11211"
	REDIS_ADDRESS_DEFAULT_VALUE     = "localhost:6379"
)

func InitCache(environment environment.Environment) cachemanager.CacheManager {

	zap.L().Info("server starting up - setting up cache manager")

	var cacheAddresses []string
	storeType := environment.GetValueOrDefault(CACHE_STORE_TYPE, store.GoCacheType).AsString()

	switch storeType {
	case store.MemcacheType:
		memcachedAddresses := environment.GetValueOrDefault(CACHE_ADDRESS, MEMCACHED_ADDRESS_DEFAULT_VALUE).AsString()
		cacheAddresses = strings.SplitN(memcachedAddresses, ",", 2)

	case store.RedisType:
		redisAddresses := environment.GetValueOrDefault(CACHE_ADDRESS, REDIS_ADDRESS_DEFAULT_VALUE).AsString()
		cacheAddresses = strings.SplitN(redisAddresses, ",", 2)
	}

	var err error
	var cacheInterface cache.CacheInterface

	if cacheInterface, err = cachemanager.BuildCacheInterface(storeType, cacheAddresses...); err != nil {
		zap.L().Fatal(fmt.Sprintf("server starting up - error setting up cache manager: %s", err.Error()))
	}

	cacheManager := cachemanager.NewDefaultCacheManager(cacheInterface)
	return cacheManager
}
