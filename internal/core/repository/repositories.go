package repository

import (
	"context"
	"jukebox-app/internal/core/model"
)

var (
	_ ArtistRepository = (*RelationalArtistRepository)(nil)
	_ SongRepository   = (*RelationalSongRepository)(nil)
	_ UserRepository   = (*RelationalUserRepository)(nil)
	_ UserRepository   = (*CachedUserRepository)(nil)
)

type ArtistRepository interface {
	Create(_ context.Context, _ *model.Artist) error
	Update(_ context.Context, _ *model.Artist) error
	DeleteById(_ context.Context, id int64) error
	FindById(_ context.Context, id int64) (*model.Artist, error)
	FindAll(_ context.Context) (*[]model.Artist, error)

	//Custom Finders

	FindByCode(_ context.Context, code int64) (*model.Artist, error)
	FindByName(_ context.Context, name string) (*model.Artist, error)
}

type SongRepository interface {
	Create(_ context.Context, _ *model.Song) error
	Update(_ context.Context, _ *model.Song) error
	DeleteById(_ context.Context, id int64) error
	FindById(_ context.Context, id int64) (*model.Song, error)
	FindAll(_ context.Context) (*[]model.Song, error)

	//Custom Finders

	FindByCode(_ context.Context, code int64) (*model.Song, error)
	FindByName(_ context.Context, name string) (*model.Song, error)
	FindByArtistId(_ context.Context, id int64) (*[]model.Song, error)
}
type UserRepository interface {
	Create(_ context.Context, _ *model.User) error
	Update(_ context.Context, _ *model.User) error
	DeleteById(_ context.Context, id int64) error
	FindById(_ context.Context, id int64) (*model.User, error)
	FindAll(_ context.Context) (*[]model.User, error)

	//Custom Finders

	FindByCode(_ context.Context, code int64) (*model.User, error)
	FindByName(_ context.Context, name string) (*model.User, error)
	FindByEmail(_ context.Context, email string) (*model.User, error)
}
