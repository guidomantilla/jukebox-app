package repository

import (
	"context"
	"testing"

	"github.com/eko/gocache/store"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"jukebox-app/internal/config"
	"jukebox-app/internal/core/model"
	"jukebox-app/pkg/cache-manager"
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

	cacheManager := cachemanager.NewDefaultCacheManager(store.GoCacheType, environment)
	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)
	err := cacheRepository.Create(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_CachedUserRepository_Create_Delegate_Err(t *testing.T) {

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
		return errors.New("some_error")
	}
	delegateRepository.EXPECT().Create(ctx, user).DoAndReturn(mockCreateFunction).Return(errors.New("some_error"))

	//

	cacheManager := cachemanager.NewDefaultCacheManager(store.GoCacheType, environment)
	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)
	err := cacheRepository.Create(ctx, user)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_CachedUserRepository_Create_Set_Err(t *testing.T) {

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

	cacheManager := cachemanager.NewMockCacheManager(ctrl)
	mockSetFunction := func(_ string, _ any, _ any) error {
		return errors.New("some_error")
	}
	cacheManager.EXPECT().Set("users", int64(1), user).DoAndReturn(mockSetFunction).Return(errors.New("some_error"))

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)

	err := cacheRepository.Create(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

//

func Test_CachedUserRepository_Update_Ok(t *testing.T) {

	var args []string
	environment := config.InitConfig(&args)
	defer config.StopConfig()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockUpdateFunction := func(ctx context.Context, user *model.User) error {
		user.Code = 100
		return nil
	}
	delegateRepository.EXPECT().Update(ctx, user).DoAndReturn(mockUpdateFunction).Return(nil)

	//

	cacheManager := cachemanager.NewDefaultCacheManager(store.GoCacheType, environment)
	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)
	err := cacheRepository.Update(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(100), user.Code)
}

func Test_CachedUserRepository_Update_Delegate_Err(t *testing.T) {

	var args []string
	environment := config.InitConfig(&args)
	defer config.StopConfig()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockUpdateFunction := func(ctx context.Context, user *model.User) error {
		return errors.New("some_error")
	}
	delegateRepository.EXPECT().Update(ctx, user).DoAndReturn(mockUpdateFunction).Return(errors.New("some_error"))

	//

	cacheManager := cachemanager.NewDefaultCacheManager(store.GoCacheType, environment)
	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)
	err := cacheRepository.Update(ctx, user)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_CachedUserRepository_Update_Set_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockUpdateFunction := func(ctx context.Context, user *model.User) error {
		user.Code = 100
		return nil
	}
	delegateRepository.EXPECT().Update(ctx, user).DoAndReturn(mockUpdateFunction).Return(nil)

	cacheManager := cachemanager.NewMockCacheManager(ctrl)
	mockSetFunction := func(_ string, _ any, _ any) error {
		return errors.New("some_error")
	}
	cacheManager.EXPECT().Set("users", int64(1), user).DoAndReturn(mockSetFunction).Return(errors.New("some_error"))

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)

	err := cacheRepository.Update(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(100), user.Code)
}

//

func Test_CachedUserRepository_Delete_Ok(t *testing.T) {

	var args []string
	environment := config.InitConfig(&args)
	defer config.StopConfig()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockDeleteFunction := func(ctx context.Context, id int64) error {
		return nil
	}
	delegateRepository.EXPECT().DeleteById(ctx, user.Id).DoAndReturn(mockDeleteFunction).Return(nil)

	//

	cacheManager := cachemanager.NewDefaultCacheManager(store.GoCacheType, environment)
	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)

	_ = cacheRepository.cacheManager.Set(cacheRepository.cacheName, user.Id, user)
	err := cacheRepository.DeleteById(ctx, user.Id)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_CachedUserRepository_Delete_Delegate_Err(t *testing.T) {

	var args []string
	environment := config.InitConfig(&args)
	defer config.StopConfig()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockDeleteFunction := func(ctx context.Context, id int64) error {
		return errors.New("some_error")
	}
	delegateRepository.EXPECT().DeleteById(ctx, user.Id).DoAndReturn(mockDeleteFunction).Return(errors.New("some_error"))

	//

	cacheManager := cachemanager.NewDefaultCacheManager(store.GoCacheType, environment)
	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)
	err := cacheRepository.DeleteById(ctx, user.Id)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_CachedUserRepository_Delete_Set_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockDeleteFunction := func(ctx context.Context, id int64) error {
		return nil
	}
	delegateRepository.EXPECT().DeleteById(ctx, user.Id).DoAndReturn(mockDeleteFunction).Return(nil)

	cacheManager := cachemanager.NewMockCacheManager(ctrl)
	mockSetFunction := func(_ string, _ any) error {
		return errors.New("some_error")
	}
	cacheManager.EXPECT().Delete("users", user.Id).DoAndReturn(mockSetFunction).Return(errors.New("some_error"))

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager)

	err := cacheRepository.DeleteById(ctx, user.Id)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}
