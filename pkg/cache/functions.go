package cache

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"

	"jukebox-app/pkg/environment"
)

func NewCacheManagerStore(storeType string, environment environment.Environment) *store.MemcacheStore {

	if storeType == "memcached" {
		return store.NewMemcache(
			memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212"),
			&store.Options{
				Expiration: 10 * time.Second,
			},
		)
	}
	return nil
}

func NewCacheManager(cacheStore store.StoreInterface) *cache.Cache {

	return cache.New(cacheStore)
}
