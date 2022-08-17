package repository

import (
	"context"
	"database/sql"
	"fmt"
	"jukebox-app/internal/core/model"
	repositoryUtils "jukebox-app/pkg/repository"

	"go.uber.org/zap"
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
	err = repositoryUtils.RelationalContext(ctx, repository.statementCreate, func(statement *sql.Stmt) error {

		var result sql.Result
		if result, err = statement.Exec(song.Code, song.Name, song.ArtistId); err != nil {
			return err
		}

		if song.Id, err = result.LastInsertId(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *RelationalSongRepository) Update(ctx context.Context, song *model.Song) error {

	var err error
	err = repositoryUtils.RelationalContext(ctx, repository.statementUpdate, func(statement *sql.Stmt) error {

		if _, err = statement.Exec(song.Code, song.Name, song.ArtistId, song.Id); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *RelationalSongRepository) DeleteById(ctx context.Context, id int64) error {

	var err error
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		if _, err = statement.Exec(id); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *RelationalSongRepository) FindById(ctx context.Context, id int64) (*model.Song, error) {

	var err error
	var song model.Song
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(id)
		if err = row.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("song with id %d not found", id)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (repository *RelationalSongRepository) FindAll(ctx context.Context) (*[]model.Song, error) {

	var err error
	songs := make([]model.Song, 0)
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		var rows *sql.Rows
		if rows, err = statement.Query(); err != nil {
			return err
		}
		defer func(rows *sql.Rows) {
			err = rows.Close()
			if err != nil {
				zap.L().Error("Error closing the result set")
			}
		}(rows)

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
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(code)
		if err = row.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("song with code %d not found", code)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (repository *RelationalSongRepository) FindByName(ctx context.Context, name string) (*model.Song, error) {

	var err error
	var song model.Song
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(name)
		if err = row.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("song with name %s not found", name)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &song, nil

}

func (repository *RelationalSongRepository) FindByArtistId(ctx context.Context, id string) (*[]model.Song, error) {

	var err error
	songs := make([]model.Song, 0)
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		var rows *sql.Rows
		if rows, err = statement.Query(id); err != nil {
			return err
		}
		defer func(rows *sql.Rows) {
			err = rows.Close()
			if err != nil {
				zap.L().Error("Error closing the result set")
			}
		}(rows)

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
