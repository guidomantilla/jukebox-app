package config

import (
	"fmt"
	cachemanager "jukebox-app/pkg/cache-manager"
	"jukebox-app/pkg/environment"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"go.uber.org/zap"
)

func InitCache(environment environment.Environment) cachemanager.CacheManager {

	zap.L().Info("server starting up - setting up cache manager")

	var err error
	var cacheInterface cache.CacheInterface

	if cacheInterface, err = cachemanager.BuildCacheInterface(store.GoCacheType, environment); err != nil {
		zap.L().Fatal(fmt.Sprintf("server starting up - error setting up cache manager: %s", err.Error()))
	}

	cacheManager := cachemanager.NewDefaultCacheManager(cacheInterface)
	return cacheManager
}
