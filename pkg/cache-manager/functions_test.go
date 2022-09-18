package cachemanager

import (
	"testing"

	"github.com/eko/gocache/v2/store"
	"github.com/stretchr/testify/assert"
)

func Test_generateKey(t *testing.T) {

	key := generateKey("sample", 1)
	assert.NotEmpty(t, key)
	assert.Equal(t, "sample-1", key)

}
func Test_NewCache(t *testing.T) {

	cache, _ := BuildCacheInterface(store.GoCacheType, "some_address")
	assert.NotNil(t, cache)

	cache, _ = BuildCacheInterface(store.MemcacheType, "some_address")
	assert.NotNil(t, cache)

	cache, _ = BuildCacheInterface(store.RedisType, "some_address")
	assert.NotNil(t, cache)

	cache, err := BuildCacheInterface("", "some_address")
	assert.Nil(t, cache)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "storeType not defined", err.Error())

}
