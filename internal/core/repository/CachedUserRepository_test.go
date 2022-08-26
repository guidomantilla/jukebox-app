package repository

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"jukebox-app/internal/config"
	"jukebox-app/internal/core/model"
	"jukebox-app/pkg/cache"
)

func Test_CachedUserRepository_Create_Ok(t *testing.T) {

	var args []string
	environment := config.InitConfig(&args)
	defer config.StopConfig()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//

	ctx := context.TODO()
	user := &model.User{
		Id:    -1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockCreateFunction := func(ctx context.Context, user *model.User) error {
		user.Id = 1
		return nil
	}
	delegateRepository.EXPECT().Create(ctx, user).DoAndReturn(mockCreateFunction).Return(nil)

	//

	cacheStore := cache.NewCacheManagerStore("memcached", environment)
	cacheManager := cache.NewCacheManager(cacheStore)

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)
	err := cacheRepository.Create(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_CachedUserRepository_Create_Err(t *testing.T) {

	var args []string
	environment := config.InitConfig(&args)
	defer config.StopConfig()

	cacheStore := cache.NewCacheManagerStore("memcached", environment)
	cacheManager := cache.NewCacheManager(cacheStore)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	user := &model.User{
		Id:    -1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockCreateFunction := func(ctx context.Context, user *model.User) error {
		user.Id = 1
		return nil
	}
	delegateRepository.EXPECT().Create(ctx, user).DoAndReturn(mockCreateFunction).Return(nil)

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)
	err := cacheRepository.Create(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}
