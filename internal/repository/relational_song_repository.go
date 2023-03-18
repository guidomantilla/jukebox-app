package repository

import (
	"context"
	"database/sql"

	feather_relational_dao "github.com/guidomantilla/go-feather-sql/pkg/feather-relational-dao"

	"jukebox-app/internal/model"
)

type RelationalSongRepository struct {
	dao                     feather_relational_dao.CrudDao
	statementFindByCode     string
	statementFindByName     string
	statementFindByArtistId string
}

/* TYPES CONSTRUCTOR */

func NewRelationalSongRepository() *RelationalSongRepository {
	return &RelationalSongRepository{
		statementFindByCode:     "select id, code, name, artistId from songs where code = ?",
		statementFindByName:     "select id, code, name, artistId from songs where name = ?",
		statementFindByArtistId: "select id, code, name, artistId from songs where artistId = ?",
	}
}

func (repository *RelationalSongRepository) Create(ctx context.Context, song *model.Song) error {

	var err error
	var id *int64
	if id, err = repository.dao.Save(ctx, song.Code, song.Name, song.ArtistId); err != nil {
		return err
	}

	song.Id = *id

	return nil
}

func (repository *RelationalSongRepository) Update(ctx context.Context, song *model.Song) error {

	var err error
	if err = repository.dao.Update(ctx, song.Code, song.Name, song.ArtistId, song.Id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalSongRepository) DeleteById(ctx context.Context, id int64) error {

	var err error
	if err = repository.dao.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalSongRepository) FindById(ctx context.Context, id int64) (*model.Song, error) {

	var err error
	var song model.Song
	if err = repository.dao.FindById(ctx, id, &song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
		return nil, err
	}

	return &song, nil
}

func (repository *RelationalSongRepository) FindAll(ctx context.Context) (*[]model.Song, error) {

	var err error
	songs := make([]model.Song, 0)
	err = repository.dao.FindAll(ctx, func(rows *sql.Rows) error {

		for rows.Next() {
			var song model.Song
			if err = rows.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
				return err
			}
			songs = append(songs, song)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &songs, nil
}

func (repository *RelationalSongRepository) FindByCode(ctx context.Context, code int64) (*model.Song, error) {

	var err error
	var song model.Song
	if err = feather_relational_dao.ReadRowContext(ctx, repository.statementFindByCode, code, &song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
		return nil, err
	}

	return &song, nil
}

func (repository *RelationalSongRepository) FindByName(ctx context.Context, name string) (*model.Song, error) {

	var err error
	var song model.Song
	if err = feather_relational_dao.ReadRowContext(ctx, repository.statementFindByName, name, &song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
		return nil, err
	}

	return &song, nil
}

func (repository *RelationalSongRepository) FindByArtistId(ctx context.Context, id int64) (*[]model.Song, error) {

	return repository.internalFindAllBy(ctx, repository.statementFindByArtistId)
}

func (repository *RelationalSongRepository) internalFindAllBy(ctx context.Context, sqlStatement string) (*[]model.Song, error) {
	var err error
	songs := make([]model.Song, 0)
	err = repository.dao.FindAll(ctx, func(rows *sql.Rows) error {

		for rows.Next() {
			var song model.Song
			if err = rows.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
				return err
			}
			songs = append(songs, song)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &songs, nil
}
