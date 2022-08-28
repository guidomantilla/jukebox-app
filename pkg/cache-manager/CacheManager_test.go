package cachemanager

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/eko/gocache/v2/store"
	cachemocks "github.com/eko/gocache/v2/test/mocks/cache"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_NewDefaultCacheManager(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cache := cachemocks.NewMockCacheInterface(ctrl)
	cacheManager := NewDefaultCacheManager(cache)

	assert.NotNil(t, cacheManager)
	assert.NotNil(t, cacheManager.cache)
}

//

func Test_Set_Ok(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockSetFunction := func(ctx context.Context, key, object interface{}, options *store.Options) error {
		return nil
	}
	cache.EXPECT().Set(ctx, "sample-1", gomock.Any(), gomock.Any()).DoAndReturn(mockSetFunction).Return(nil)

	cacheManager := NewDefaultCacheManager(cache)
	err := cacheManager.Set(ctx, "sample", 1, "value")

	assert.Nil(t, err)
}

func Test_Set_Marshal_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	cacheManager := NewDefaultCacheManager(cache)
	cacheManager.marshalFunc = func(v any) ([]byte, error) {
		return nil, errors.New("some_error")
	}

	err := cacheManager.Set(ctx, "sample", 1, "value")

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_Set_Set_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockSetFunction := func(ctx context.Context, key, object interface{}, options *store.Options) error {
		return errors.New("some_error")
	}
	cache.EXPECT().Set(ctx, "sample-1", gomock.Any(), gomock.Any()).DoAndReturn(mockSetFunction).Return(errors.New("some_error"))

	cacheManager := NewDefaultCacheManager(cache)
	err := cacheManager.Set(ctx, "sample", 1, "value")

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_Get_Ok(t *testing.T) {

	type Page struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	mkyong := Page{
		Name: "name",
		Url:  "mkyong.com",
	}
	marshal, _ := json.Marshal(mkyong)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockGetFunction := func(ctx context.Context, key interface{}) (interface{}, error) {
		return marshal, nil
	}
	cache.EXPECT().Get(ctx, "sample-1").DoAndReturn(mockGetFunction).Return(marshal, nil)

	cacheManager := NewDefaultCacheManager(cache)

	var mkyong2 Page
	err := cacheManager.Get(ctx, "sample", 1, &mkyong2)

	assert.Nil(t, err)
	assert.NotEmpty(t, mkyong2)
	assert.Equal(t, "mkyong.com", mkyong2.Url)
}

func Test_Get_Get_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockGetFunction := func(ctx context.Context, key interface{}) (interface{}, error) {
		return nil, errors.New("some_error")
	}
	cache.EXPECT().Get(ctx, "sample-1").DoAndReturn(mockGetFunction).Return(nil, errors.New("some_error"))

	cacheManager := NewDefaultCacheManager(cache)

	err := cacheManager.Get(ctx, "sample", 1, nil)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_Get_Unmarshal_Err(t *testing.T) {

	type Page struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	mkyong := Page{
		Name: "name",
		Url:  "mkyong.com",
	}
	marshal, _ := json.Marshal(mkyong)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockGetFunction := func(ctx context.Context, key interface{}) (interface{}, error) {
		return marshal, nil
	}
	cache.EXPECT().Get(ctx, "sample-1").DoAndReturn(mockGetFunction).Return(marshal, nil)

	cacheManager := NewDefaultCacheManager(cache)
	cacheManager.unmarshalFunc = func(data []byte, v any) error {
		return errors.New("some_error")
	}

	err := cacheManager.Get(ctx, "sample", 1, nil)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_Delete_Ok(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockDeleteFunction := func(ctx context.Context, key interface{}) error {
		return nil
	}
	cache.EXPECT().Delete(ctx, "sample-1").DoAndReturn(mockDeleteFunction).Return(nil)

	cacheManager := NewDefaultCacheManager(cache)

	err := cacheManager.Delete(ctx, "sample", 1)

	assert.Nil(t, err)
}

func Test_Delete_Delete_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockDeleteFunction := func(ctx context.Context, key interface{}) error {
		return errors.New("some_error")
	}
	cache.EXPECT().Delete(ctx, "sample-1").DoAndReturn(mockDeleteFunction).Return(errors.New("some_error"))

	cacheManager := NewDefaultCacheManager(cache)

	err := cacheManager.Delete(ctx, "sample", 1)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_Invalidate_Ok(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockInvalidateFunction := func(ctx context.Context, options store.InvalidateOptions) error {
		return nil
	}
	cache.EXPECT().Invalidate(ctx, store.InvalidateOptions{Tags: []string{"sample"}}).DoAndReturn(mockInvalidateFunction).Return(nil)

	cacheManager := NewDefaultCacheManager(cache)

	err := cacheManager.Invalidate(ctx, "sample")

	assert.Nil(t, err)
}

func Test_Invalidate_Invalidate_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockInvalidateFunction := func(ctx context.Context, options store.InvalidateOptions) error {
		return errors.New("some_error")
	}
	cache.EXPECT().Invalidate(ctx, store.InvalidateOptions{Tags: []string{"sample"}}).DoAndReturn(mockInvalidateFunction).Return(errors.New("some_error"))

	cacheManager := NewDefaultCacheManager(cache)

	err := cacheManager.Invalidate(ctx, "sample")

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_Clear_Ok(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockClearFunction := func(ctx context.Context) error {
		return nil
	}
	cache.EXPECT().Clear(ctx).DoAndReturn(mockClearFunction).Return(nil)

	cacheManager := NewDefaultCacheManager(cache)

	err := cacheManager.Clear(ctx)

	assert.Nil(t, err)
}

func Test_Clear_Clear_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockClearFunction := func(ctx context.Context) error {
		return errors.New("some_error")
	}
	cache.EXPECT().Clear(ctx).DoAndReturn(mockClearFunction).Return(errors.New("some_error"))

	cacheManager := NewDefaultCacheManager(cache)

	err := cacheManager.Clear(ctx)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_GetType_Ok(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cache := cachemocks.NewMockCacheInterface(ctrl)
	mockGetTypeFunction := func() string {
		return "some_type"
	}
	cache.EXPECT().GetType().DoAndReturn(mockGetTypeFunction).Return("some_type")

	cacheManager := NewDefaultCacheManager(cache)

	storeType := cacheManager.GetType()

	assert.NotEmpty(t, storeType)
	assert.Equal(t, "some_type", storeType)
}
