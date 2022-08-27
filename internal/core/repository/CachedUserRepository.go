package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"jukebox-app/internal/core/model"
	"jukebox-app/pkg/cache-manager"
)

const (
	USER_ID = "Users-Id(%s)"
)

type CachedUserRepository struct {
	delegateUserRepository UserRepository
	cacheManager           cachemanager.CacheManager
	cacheName              string
}

func (repository *CachedUserRepository) Create(ctx context.Context, user *model.User) error {

	var err error
	if err = repository.delegateUserRepository.Create(ctx, user); err != nil {
		return err
	}

	if err = repository.cacheManager.Set(ctx, repository.cacheName, user.Id, user); err != nil {
		zap.L().Error(fmt.Sprintf("Error caching the created user: %s", err.Error()))
		return nil
	}

	return nil
}

func (repository *CachedUserRepository) Update(ctx context.Context, user *model.User) error {

	var err error
	if err = repository.delegateUserRepository.Update(ctx, user); err != nil {
		return err
	}

	if err = repository.cacheManager.Set(ctx, repository.cacheName, user.Id, user); err != nil {
		zap.L().Error("Error caching the updated user")
		return nil
	}

	return nil
}

func (repository *CachedUserRepository) DeleteById(ctx context.Context, id int64) error {

	var err error
	if err = repository.delegateUserRepository.DeleteById(ctx, id); err != nil {
		return err
	}

	if err = repository.cacheManager.Delete(ctx, repository.cacheName, id); err != nil {
		if !errors.Is(err, memcache.ErrCacheMiss) {
			zap.L().Error("Error deleting the user from cache")
		}
		return nil
	}

	return nil
}

func (repository *CachedUserRepository) FindById(ctx context.Context, id int64) (*model.User, error) {

	var err error
	var user *model.User
	var valueFromCache any

	err = repository.cacheManager.Get(ctx, repository.cacheName, id, &valueFromCache)
	if err != nil {
		if !errors.Is(err, memcache.ErrCacheMiss) {
			zap.L().Error("Error getting the user from cache")
		}
		return repository.findByIdAndSet(ctx, id)
	}

	user = &model.User{}
	if err = json.Unmarshal([]byte(valueFromCache.(string)), user); err != nil {
		zap.L().Error("Error unmarshalling from json the user")
		return repository.findByIdAndSet(ctx, id)
	}

	return user, nil
}

func (repository *CachedUserRepository) FindAll(ctx context.Context) (*[]model.User, error) {
	return nil, nil
}

func (repository *CachedUserRepository) FindByCode(ctx context.Context, code int64) (*model.User, error) {
	return nil, nil
}

func (repository *CachedUserRepository) FindByName(ctx context.Context, name string) (*model.User, error) {
	return nil, nil
}

func (repository *CachedUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}

//

func (repository *CachedUserRepository) findByIdAndSet(ctx context.Context, id int64) (*model.User, error) {

	var err error
	var user *model.User

	if user, err = repository.delegateUserRepository.FindById(ctx, id); err != nil {
		return nil, err
	}

	if err = repository.cacheManager.Set(ctx, repository.cacheName, id, user); err != nil {
		zap.L().Error("Error caching the user")
		return user, nil
	}

	return user, nil
}

func (repository *CachedUserRepository) FindAllAndSet(ctx context.Context) (*[]model.User, error) {

	var err error
	var users *[]model.User

	if users, err = repository.delegateUserRepository.FindAll(ctx); err != nil {
		return nil, err
	}

	if err = repository.cacheManager.Set(ctx, repository.cacheName, "", users); err != nil {
		zap.L().Error("Error caching the users")
		return users, nil
	}

	return users, nil
}

//

func NewCachedUserRepository(delegateUserRepository UserRepository, cacheManager cachemanager.CacheManager) *CachedUserRepository {
	return &CachedUserRepository{
		delegateUserRepository: delegateUserRepository,
		cacheManager:           cacheManager,
		cacheName:              "users",
	}
}
