package repository

import (
	"context"
	"database/sql"
	"jukebox-app/internal/model"

	repositoryUtils "jukebox-app/pkg/repository"
)

type RelationalSongRepository struct {
	statementCreate         string
	statementUpdate         string
	statementDelete         string
	statementFindById       string
	statementFind           string
	statementFindByCode     string
	statementFindByName     string
	statementFindByArtistId string
}

func (repository *RelationalSongRepository) Create(ctx context.Context, song *model.Song) error {

	var err error
	var id *int64
	if id, err = repositoryUtils.RelationalWriteContext(ctx, repository.statementCreate, song.Code, song.Name, song.ArtistId); err != nil {
		return err
	}

	song.Id = *id

	return nil
}

func (repository *RelationalSongRepository) Update(ctx context.Context, song *model.Song) error {

	var err error
	if _, err = repositoryUtils.RelationalWriteContext(ctx, repository.statementUpdate, song.Code, song.Name, song.ArtistId, song.Id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalSongRepository) DeleteById(ctx context.Context, id int64) error {

	var err error
	if _, err = repositoryUtils.RelationalWriteContext(ctx, repository.statementDelete, id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalSongRepository) FindById(ctx context.Context, id int64) (*model.Song, error) {

	var err error
	var song model.Song
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindById, id, &song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
		return nil, err
	}

	return &song, nil
}

func (repository *RelationalSongRepository) FindAll(ctx context.Context) (*[]model.Song, error) {

	return internalFindAllBy(ctx, repository.statementFind)
}

func (repository *RelationalSongRepository) FindByCode(ctx context.Context, code int64) (*model.Song, error) {

	var err error
	var song model.Song
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindByCode, code, &song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
		return nil, err
	}

	return &song, nil
}

func (repository *RelationalSongRepository) FindByName(ctx context.Context, name string) (*model.Song, error) {

	var err error
	var song model.Song
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindByName, name, &song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
		return nil, err
	}

	return &song, nil
}

func (repository *RelationalSongRepository) FindByArtistId(ctx context.Context, id int64) (*[]model.Song, error) {

	return internalFindAllBy(ctx, repository.statementFindByArtistId)
}

/* TYPES CONSTRUCTOR */

func NewRelationalSongRepository() *RelationalSongRepository {
	return &RelationalSongRepository{
		statementCreate:   "insert into songs (code, name, artistId) values (?, ?, ?)",
		statementUpdate:   "update songs set code = ?, name = ?, artistId = ? where id = ?",
		statementDelete:   "delete from songs where id = ?",
		statementFindById: "select id, code, name, artistId from songs where id = ?",
		statementFind:     "select id, code, name, artistId from songs",

		statementFindByCode:     "select id, code, name, artistId from songs where code = ?",
		statementFindByName:     "select id, code, name, artistId from songs where name = ?",
		statementFindByArtistId: "select id, code, name, artistId from songs where artistId = ?",
	}
}

func internalFindAllBy(ctx context.Context, sqlStatement string) (*[]model.Song, error) {
	var err error
	songs := make([]model.Song, 0)
	err = repositoryUtils.RelationalQueryContext(ctx, sqlStatement, func(rows *sql.Rows) error {

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
