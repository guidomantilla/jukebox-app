package repository

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"jukebox-app/internal/model"
	"jukebox-app/pkg/cachemanager"
)

func Test_CachedUserRepository_Create_Ok(t *testing.T) {

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
	mockSetFunction := func(_ context.Context, _ string, _ any, _ any) error {
		return nil
	}
	cacheManager.EXPECT().Set(ctx, "users", int64(1), user).DoAndReturn(mockSetFunction).Return(nil)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	err := cacheRepository.Create(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_CachedUserRepository_Create_Delegate_Err(t *testing.T) {

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

	cacheManager := cachemanager.NewMockCacheManager(ctrl)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
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
	mockSetFunction := func(_ context.Context, _ string, _ any, _ any) error {
		return errors.New("some_error")
	}
	cacheManager.EXPECT().Set(ctx, "users", int64(1), user).DoAndReturn(mockSetFunction).Return(errors.New("some_error"))

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	err := cacheRepository.Create(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

//

func Test_CachedUserRepository_Update_Ok(t *testing.T) {

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
	mockSetFunction := func(_ context.Context, _ string, _ any, _ any) error {
		return nil
	}
	cacheManager.EXPECT().Set(ctx, "users", int64(1), user).DoAndReturn(mockSetFunction).Return(nil)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	err := cacheRepository.Update(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(100), user.Code)
}

func Test_CachedUserRepository_Update_Delegate_Err(t *testing.T) {

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

	cacheManager := cachemanager.NewMockCacheManager(ctrl)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
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
	mockSetFunction := func(_ context.Context, _ string, _ any, _ any) error {
		return errors.New("some_error")
	}
	cacheManager.EXPECT().Set(ctx, "users", int64(1), user).DoAndReturn(mockSetFunction).Return(errors.New("some_error"))

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	err := cacheRepository.Update(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, int64(100), user.Code)
}

//

func Test_CachedUserRepository_Delete_Ok(t *testing.T) {

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
	mockDeleteFunction2 := func(_ context.Context, _ string, _ int64) error {
		return nil
	}
	cacheManager.EXPECT().Delete(ctx, "users", int64(1)).DoAndReturn(mockDeleteFunction2).Return(nil)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	err := cacheRepository.DeleteById(ctx, user.Id)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_CachedUserRepository_Delete_Delegate_Err(t *testing.T) {

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

	cacheManager := cachemanager.NewMockCacheManager(ctrl)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
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
	mockSetFunction := func(_ context.Context, _ string, _ any) error {
		return errors.New("some_error")
	}
	cacheManager.EXPECT().Delete(ctx, "users", user.Id).DoAndReturn(mockSetFunction).Return(errors.New("some_error"))

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	err := cacheRepository.DeleteById(ctx, user.Id)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

//

func Test_CachedUserRepository_findByIdAndSet_Ok(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockFindByIdFunction := func(ctx context.Context, id int64) (*model.User, error) {
		return user, nil
	}
	delegateRepository.EXPECT().FindById(ctx, user.Id).DoAndReturn(mockFindByIdFunction).Return(user, nil)

	cacheManager := cachemanager.NewMockCacheManager(ctrl)
	mockSetFunction := func(_ context.Context, _ string, _ any, _ any) error {
		return nil
	}
	cacheManager.EXPECT().Set(ctx, "users", user.Id, user).DoAndReturn(mockSetFunction).Return(nil)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	user2, err := cacheRepository.findByIdAndSet(ctx, user.Id)

	assert.Nil(t, err)
	assert.Equal(t, user.Id, user2.Id)
	assert.Equal(t, user.Code, user2.Code)
	assert.Equal(t, user.Name, user2.Name)
	assert.Equal(t, user.Email, user2.Email)
}

func Test_CachedUserRepository_findByIdAndSet_Delegate_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockFindByIdFunction := func(ctx context.Context, id int64) (*model.User, error) {
		return nil, errors.New("some_error")
	}
	delegateRepository.EXPECT().FindById(ctx, user.Id).DoAndReturn(mockFindByIdFunction).Return(nil, errors.New("some_error"))

	cacheManager := cachemanager.NewMockCacheManager(ctrl)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	user2, err := cacheRepository.findByIdAndSet(ctx, user.Id)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, user2)
	assert.Equal(t, "some_error", err.Error())
}

func Test_CachedUserRepository_findByIdAndSet_Set_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockFindByIdFunction := func(ctx context.Context, id int64) (*model.User, error) {
		return user, nil
	}
	delegateRepository.EXPECT().FindById(ctx, user.Id).DoAndReturn(mockFindByIdFunction).Return(user, nil)

	cacheManager := cachemanager.NewMockCacheManager(ctrl)
	mockSetFunction := func(_ context.Context, _ string, _ any, _ any) error {
		return errors.New("some_error")
	}
	cacheManager.EXPECT().Set(ctx, "users", user.Id, user).DoAndReturn(mockSetFunction).Return(errors.New("some_error"))

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	user2, err := cacheRepository.findByIdAndSet(ctx, user.Id)

	assert.Nil(t, err)
	assert.Equal(t, user.Id, user2.Id)
	assert.Equal(t, user.Code, user2.Code)
	assert.Equal(t, user.Name, user2.Name)
	assert.Equal(t, user.Email, user2.Email)
}

//

func Test_CachedUserRepository_findAllAndSet_Ok(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	users := &[]model.User{
		{Id: 1, Code: 2, Name: "3", Email: "4"},
		{Id: 5, Code: 6, Name: "7", Email: "8"},
		{Id: 9, Code: 10, Name: "11", Email: "12"},
		{Id: 13, Code: 14, Name: "15", Email: "16"},
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockFindAllFunction := func(ctx context.Context) (*[]model.User, error) {
		return users, nil
	}
	delegateRepository.EXPECT().FindAll(ctx).DoAndReturn(mockFindAllFunction).Return(users, nil)

	cacheManager := cachemanager.NewMockCacheManager(ctrl)
	mockSetFunction := func(_ context.Context, _ string, _ any, _ any) error {
		return nil
	}
	cacheManager.EXPECT().Set(ctx, "users", "all", users).DoAndReturn(mockSetFunction).Return(nil)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	_, err := cacheRepository.findAllAndSet(ctx)

	assert.Nil(t, err)
}

func Test_CachedUserRepository_findAllAndSet_Delegate_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	delegateRepository := NewMockUserRepository(ctrl)
	mockFindAllFunction := func(ctx context.Context) (*[]model.User, error) {
		return nil, errors.New("some_error")
	}
	delegateRepository.EXPECT().FindAll(ctx).DoAndReturn(mockFindAllFunction).Return(nil, errors.New("some_error"))

	cacheManager := cachemanager.NewMockCacheManager(ctrl)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	_, err := cacheRepository.findAllAndSet(ctx)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_CachedUserRepository_findAllAndSet_Set_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	users := &[]model.User{
		{Id: 1, Code: 2, Name: "3", Email: "4"},
		{Id: 5, Code: 6, Name: "7", Email: "8"},
		{Id: 9, Code: 10, Name: "11", Email: "12"},
		{Id: 13, Code: 14, Name: "15", Email: "16"},
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockFindAllFunction := func(ctx context.Context) (*[]model.User, error) {
		return users, nil
	}
	delegateRepository.EXPECT().FindAll(ctx).DoAndReturn(mockFindAllFunction).Return(users, nil)

	cacheManager := cachemanager.NewMockCacheManager(ctrl)
	mockSetFunction := func(_ context.Context, _ string, _ any, _ any) error {
		return errors.New("some_error")
	}
	cacheManager.EXPECT().Set(ctx, "users", "all", users).DoAndReturn(mockSetFunction).Return(errors.New("some_error"))

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	_, err := cacheRepository.findAllAndSet(ctx)

	assert.Nil(t, err)
}

//

func Test_CachedUserRepository_FindById_Ok(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	cacheManager := cachemanager.NewMockCacheManager(ctrl)

	mockGetFunction := func(_ context.Context, _ string, _ any, something any) error {
		marshal, _ := json.Marshal(user)
		_ = json.Unmarshal(marshal, something)
		return nil
	}
	cacheManager.EXPECT().Get(ctx, "users", user.Id, gomock.Any()).DoAndReturn(mockGetFunction).Return(nil)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	user2, err := cacheRepository.FindById(ctx, user.Id)

	assert.Nil(t, err)
	assert.Equal(t, user.Id, user2.Id)
	assert.Equal(t, user.Code, user2.Code)
	assert.Equal(t, user.Name, user2.Name)
	assert.Equal(t, user.Email, user2.Email)
}

func Test_CachedUserRepository_Get_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	delegateRepository := NewMockUserRepository(ctrl)
	mockFindByIdFunction := func(ctx context.Context, id int64) (*model.User, error) {
		return user, nil
	}
	delegateRepository.EXPECT().FindById(ctx, user.Id).DoAndReturn(mockFindByIdFunction).Return(user, nil)

	cacheManager := cachemanager.NewMockCacheManager(ctrl)
	mockGetFunction := func(_ context.Context, _ string, _ any, something any) error {
		return errors.New("some_error")
	}
	cacheManager.EXPECT().Get(ctx, "users", user.Id, gomock.Any()).DoAndReturn(mockGetFunction).Return(errors.New("some_error"))

	mockSetFunction := func(_ context.Context, _ string, _ any, _ any) error {
		return nil
	}
	cacheManager.EXPECT().Set(ctx, "users", user.Id, user).DoAndReturn(mockSetFunction).Return(nil)

	//

	cacheRepository := NewCachedUserRepository(delegateRepository, cacheManager, json.Marshal, json.Unmarshal)
	user2, err := cacheRepository.FindById(ctx, user.Id)

	assert.Nil(t, err)
	assert.Equal(t, user.Id, user2.Id)
	assert.Equal(t, user.Code, user2.Code)
	assert.Equal(t, user.Name, user2.Name)
	assert.Equal(t, user.Email, user2.Email)
}
