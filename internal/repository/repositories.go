package repository

import (
	"context"
	"jukebox-app/internal/model"
)

var (
	_ ArtistRepository = (*RelationalArtistRepository)(nil)
	_ SongRepository   = (*RelationalSongRepository)(nil)
	_ UserRepository   = (*RelationalUserRepository)(nil)
	_ UserRepository   = (*CachedUserRepository)(nil)
)

type ArtistRepository interface {
	Create(ctx context.Context, artist *model.Artist) error
	Update(ctx context.Context, artist *model.Artist) error
	DeleteById(ctx context.Context, id int64) error
	FindById(ctx context.Context, id int64) (*model.Artist, error)
	FindAll(ctx context.Context) (*[]model.Artist, error)

	//Custom Finders

	FindByCode(ctx context.Context, code int64) (*model.Artist, error)
	FindByName(ctx context.Context, name string) (*model.Artist, error)
}

type SongRepository interface {
	Create(ctx context.Context, song *model.Song) error
	Update(ctx context.Context, song *model.Song) error
	DeleteById(ctx context.Context, id int64) error
	FindById(ctx context.Context, id int64) (*model.Song, error)
	FindAll(ctx context.Context) (*[]model.Song, error)

	//Custom Finders

	FindByCode(ctx context.Context, code int64) (*model.Song, error)
	FindByName(ctx context.Context, name string) (*model.Song, error)
	FindByArtistId(ctx context.Context, id int64) (*[]model.Song, error)
}
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	DeleteById(ctx context.Context, id int64) error
	FindById(ctx context.Context, id int64) (*model.User, error)
	FindAll(ctx context.Context) (*[]model.User, error)

	//Custom Finders

	FindByCode(ctx context.Context, code int64) (*model.User, error)
	FindByName(ctx context.Context, name string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
