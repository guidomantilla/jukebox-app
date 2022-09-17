package cachemanager

import (
	"testing"

	"github.com/eko/gocache/v2/store"
	"github.com/stretchr/testify/assert"

	"jukebox-app/internal/config"
)

func Test_generateKey(t *testing.T) {

	key := generateKey("sample", 1)
	assert.NotEmpty(t, key)
	assert.Equal(t, "sample-1", key)

}
func Test_NewCache(t *testing.T) {

	var args []string
	environment := config.InitConfig(&args)
	defer config.StopConfig()

	cache, _ := BuildCacheInterface(store.GoCacheType, environment)
	assert.NotNil(t, cache)

	cache, _ = BuildCacheInterface(store.MemcacheType, environment)
	assert.NotNil(t, cache)

	cache, _ = BuildCacheInterface(store.RedisType, environment)
	assert.NotNil(t, cache)

	cache, err := BuildCacheInterface("", environment)
	assert.Nil(t, cache)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "storeType not defined", err.Error())

}
