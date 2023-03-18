package repository

import (
	"context"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/guidomantilla/go-feather-commons/pkg/encodingjson"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"jukebox-app/internal/model"
	"jukebox-app/pkg/cachemanager"
)

type CachedUserRepository struct {
	delegateUserRepository UserRepository
	cacheManager           cachemanager.CacheManager
	cacheName              string
	marshalFunc            encodingjson.MarshalFunc
	unmarshalFunc          encodingjson.UnmarshalFunc
}

func (repository *CachedUserRepository) Create(ctx context.Context, user *model.User) error {

	var err error
	if err = repository.delegateUserRepository.Create(ctx, user); err != nil {
		return err
	}

	if err = repository.cacheManager.Set(ctx, repository.cacheName, user.Id, user); err != nil {
		zap.L().Error(fmt.Sprintf("error caching the created user: %s", err.Error()))
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
		zap.L().Error("error caching the updated user")
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
			zap.L().Error("error deleting the user from cache")
		}
		return nil
	}

	return nil
}

func (repository *CachedUserRepository) FindById(ctx context.Context, id int64) (*model.User, error) {

	var err error
	var user *model.User

	if err = repository.cacheManager.Get(ctx, repository.cacheName, id, &user); err != nil {
		zap.L().Error(fmt.Sprintf("error getting the user from cache: %s", err.Error()))
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
		zap.L().Error("error caching the user")
		return user, nil
	}

	return user, nil
}

func (repository *CachedUserRepository) findAllAndSet(ctx context.Context) (*[]model.User, error) {

	var err error
	var users *[]model.User

	if users, err = repository.delegateUserRepository.FindAll(ctx); err != nil {
		return nil, err
	}

	if err = repository.cacheManager.Set(ctx, repository.cacheName, "all", users); err != nil {
		zap.L().Error("error caching the users")
		return users, nil
	}

	return users, nil
}

//

func NewCachedUserRepository(delegateUserRepository UserRepository, cacheManager cachemanager.CacheManager,
	marshalFunc encodingjson.MarshalFunc, unmarshalFunc encodingjson.UnmarshalFunc) *CachedUserRepository {
	return &CachedUserRepository{
		delegateUserRepository: delegateUserRepository,
		cacheManager:           cacheManager,
		cacheName:              "users",
		marshalFunc:            marshalFunc,
		unmarshalFunc:          unmarshalFunc,
	}
}
