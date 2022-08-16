package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"jukebox-app/internal/core/model"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type CachedUserRepository struct {
	delegateUserRepository UserRepository
	memcachedClient        *memcache.Client
}

func (repository *CachedUserRepository) Create(ctx context.Context, user *model.User) error {

	var err error
	if err = repository.delegateUserRepository.Create(ctx, user); err != nil {
		return err
	}

	var valueToCache []byte
	if valueToCache, err = json.Marshal(&user); err != nil {
		zap.L().Error("Error marshalling to json the created user")
		return nil
	}
	keyToCache := fmt.Sprintf("Users-Id(%s)", strconv.FormatInt(user.Id, 10))

	userToCache := &memcache.Item{
		Key:   keyToCache,
		Value: valueToCache,
	}
	if err = repository.memcachedClient.Set(userToCache); err != nil {
		zap.L().Error("Error caching the created user")
		return nil
	}

	return nil
}

func (repository *CachedUserRepository) Update(ctx context.Context, user *model.User) error {

	var err error
	if err = repository.delegateUserRepository.Update(ctx, user); err != nil {
		return err
	}

	var valueToCache []byte
	if valueToCache, err = json.Marshal(&user); err != nil {
		zap.L().Error("Error marshalling to json the updated user")
		return nil
	}
	keyToCache := fmt.Sprintf("Users-Id(%s)", strconv.FormatInt(user.Id, 10))

	userToCache := &memcache.Item{
		Key:   keyToCache,
		Value: valueToCache,
	}
	if err = repository.memcachedClient.Set(userToCache); err != nil {
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

	keyToCache := fmt.Sprintf("Users-Id(%s)", strconv.FormatInt(id, 10))
	if err = repository.memcachedClient.Delete(keyToCache); err != nil {
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
	var cachedUser *memcache.Item

	keyToCache := fmt.Sprintf("Users-Id(%s)", strconv.FormatInt(id, 10))

	cachedUser, err = repository.memcachedClient.Get(keyToCache)
	if err != nil {
		if !errors.Is(err, memcache.ErrCacheMiss) {
			zap.L().Error("Error getting the user from cache")
		}
		return repository.findByIdAndSet(ctx, id)
	}

	user = &model.User{}
	if err = json.Unmarshal(cachedUser.Value, user); err != nil {
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

	var valueToCache []byte
	if valueToCache, err = json.Marshal(&user); err != nil {
		zap.L().Error("Error marshalling to json the user")
		return user, nil
	}
	keyToCache := fmt.Sprintf("Users-Id(%s)", strconv.FormatInt(id, 10))

	userToCache := &memcache.Item{
		Key:   keyToCache,
		Value: valueToCache,
	}
	if err = repository.memcachedClient.Set(userToCache); err != nil {
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

	var valueToCache []byte
	if valueToCache, err = json.Marshal(&users); err != nil {
		zap.L().Error("Error marshalling to json the users")
		return users, nil
	}
	keyToCache := "Users-FindAll()"

	usersToCache := &memcache.Item{
		Key:   keyToCache,
		Value: valueToCache,
	}
	if err = repository.memcachedClient.Set(usersToCache); err != nil {
		zap.L().Error("Error caching the users")
		return users, nil
	}

	return users, nil
}

//

func NewCachedUserRepository(delegateUserRepository UserRepository, memcachedClient *memcache.Client) *CachedUserRepository {
	return &CachedUserRepository{
		delegateUserRepository: delegateUserRepository,
		memcachedClient:        memcachedClient,
	}
}
