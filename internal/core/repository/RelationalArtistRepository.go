package repository

import (
	"context"
	"database/sql"
	"fmt"
	"jukebox-app/internal/core/model"
	repositoryUtils "jukebox-app/pkg/repository"

	"go.uber.org/zap"
)

type RelationalArtistRepository struct {
	statementCreate     string
	statementUpdate     string
	statementDelete     string
	statementFindById   string
	statementFind       string
	statementFindByCode string
	statementFindByName string
}

func (repository *RelationalArtistRepository) Create(ctx context.Context, artist *model.Artist) error {

	var err error
	err = repositoryUtils.RelationalContext(ctx, repository.statementCreate, func(statement *sql.Stmt) error {

		var result sql.Result
		if result, err = statement.Exec(artist.Code, artist.Name); err != nil {
			return err
		}

		if artist.Id, err = result.LastInsertId(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *RelationalArtistRepository) Update(ctx context.Context, artist *model.Artist) error {

	var err error
	err = repositoryUtils.RelationalContext(ctx, repository.statementUpdate, func(statement *sql.Stmt) error {

		if _, err = statement.Exec(artist.Code, artist.Name, artist.Id); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *RelationalArtistRepository) DeleteById(ctx context.Context, id int64) error {

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

func (repository *RelationalArtistRepository) FindById(ctx context.Context, id int64) (*model.Artist, error) {

	var err error
	var artist model.Artist
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(id)
		if err = row.Scan(&artist.Id, &artist.Code, &artist.Name); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("artist with id %d not found", id)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (repository *RelationalArtistRepository) FindAll(ctx context.Context) (*[]model.Artist, error) {

	var err error
	artists := make([]model.Artist, 0)
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
			var artist model.Artist
			if err = rows.Scan(&artist.Id, &artist.Code, &artist.Name); err != nil {
				return err
			}
			artists = append(artists, artist)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &artists, nil
}

func (repository *RelationalArtistRepository) FindByCode(ctx context.Context, code int64) (*model.Artist, error) {

	var err error
	var artist model.Artist
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(code)
		if err = row.Scan(&artist.Id, &artist.Code, &artist.Name); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("artist with code %d not found", code)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (repository *RelationalArtistRepository) FindByName(ctx context.Context, name string) (*model.Artist, error) {

	var err error
	var artist model.Artist
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(name)
		if err = row.Scan(&artist.Id, &artist.Code, &artist.Name); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("artist with name %s not found", name)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

/* TYPES CONSTRUCTOR */

func NewRelationalArtistRepository() *RelationalArtistRepository {
	return &RelationalArtistRepository{
		statementCreate:   "insert into artists (code, name) values (?, ?)",
		statementUpdate:   "update artists set code = ?, name = ? where id = ?",
		statementDelete:   "delete from artists where id = ?",
		statementFindById: "select id, code, name from artists where id = ?",
		statementFind:     "select id, code, name from artists",

		statementFindByCode: "select id, code, name from artists where code = ?",
		statementFindByName: "select id, code, name from artists where name = ?",
	}
}
