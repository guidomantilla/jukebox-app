package cachemanager

import (
	"context"
	"encoding/json"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"

	"jukebox-app/pkg/encodingjson"
)

var _ CacheManager = (*DefaultCacheManager)(nil)

type CacheManager interface {
	Set(ctx context.Context, cacheName string, key any, value any) error
	Get(ctx context.Context, cacheName string, key any, value any) error
	Delete(ctx context.Context, cacheName string, key any) error
	Invalidate(ctx context.Context, cacheName string) error
	Clear(ctx context.Context) error
	GetType() string
}

type DefaultCacheManager struct {
	cache         cache.CacheInterface
	marshalFunc   encodingjson.MarshalFunc
	unmarshalFunc encodingjson.UnmarshalFunc
}

func (cacheManager *DefaultCacheManager) Set(ctx context.Context, cacheName string, key any, value any) error {
	var err error
	var valueToCache []byte

	keyToCache := generateKey(cacheName, key)
	if valueToCache, err = cacheManager.marshalFunc(value); err != nil {
		return err
	}

	if err = cacheManager.cache.Set(ctx, keyToCache, valueToCache, &store.Options{Tags: []string{cacheName}}); err != nil {
		return err
	}
	return nil
}

func (cacheManager *DefaultCacheManager) Get(ctx context.Context, cacheName string, key any, value any) error {
	var err error
	var data any

	keyToCache := generateKey(cacheName, key)
	if data, err = cacheManager.cache.Get(ctx, keyToCache); err != nil {
		return err
	}

	byteSlice := data.([]byte)
	if err = cacheManager.unmarshalFunc(byteSlice, value); err != nil {
		return err
	}

	return nil
}

func (cacheManager *DefaultCacheManager) Delete(ctx context.Context, cacheName string, key any) error {
	var err error

	keyToCache := generateKey(cacheName, key)
	if err = cacheManager.cache.Delete(ctx, keyToCache); err != nil {
		return err
	}

	return nil
}

func (cacheManager *DefaultCacheManager) Invalidate(ctx context.Context, cacheName string) error {
	var err error

	if err = cacheManager.cache.Invalidate(ctx, store.InvalidateOptions{Tags: []string{cacheName}}); err != nil {
		return err
	}

	return nil
}

func (cacheManager *DefaultCacheManager) Clear(ctx context.Context) error {

	var err error

	if err = cacheManager.cache.Clear(ctx); err != nil {
		return err
	}

	return nil
}

func (cacheManager *DefaultCacheManager) GetType() string {
	return cacheManager.cache.GetType()
}

//

func NewDefaultCacheManager(cache cache.CacheInterface) *DefaultCacheManager {
	return &DefaultCacheManager{
		cache:         cache,
		marshalFunc:   json.Marshal,
		unmarshalFunc: json.Unmarshal,
	}
}
