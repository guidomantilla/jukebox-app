package repository

import (
	"context"
	"database/sql"
	"jukebox-app/internal/core/model"
	repositoryUtils "jukebox-app/pkg/repository"
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
	var id *int64
	if id, err = repositoryUtils.RelationalWriteContext(ctx, repository.statementCreate, artist.Code, artist.Name); err != nil {
		return err
	}

	artist.Id = *id

	return nil
}

func (repository *RelationalArtistRepository) Update(ctx context.Context, artist *model.Artist) error {

	var err error
	if _, err = repositoryUtils.RelationalWriteContext(ctx, repository.statementUpdate, artist.Code, artist.Name, artist.Id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalArtistRepository) DeleteById(ctx context.Context, id int64) error {

	var err error
	if _, err = repositoryUtils.RelationalWriteContext(ctx, repository.statementDelete, id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalArtistRepository) FindById(ctx context.Context, id int64) (*model.Artist, error) {

	var err error
	var artist model.Artist
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindById, id, &artist.Id, &artist.Code, &artist.Name); err != nil {
		return nil, err
	}

	return &artist, nil
}

func (repository *RelationalArtistRepository) FindAll(ctx context.Context) (*[]model.Artist, error) {

	var err error
	artists := make([]model.Artist, 0)
	err = repositoryUtils.RelationalQueryContext(ctx, repository.statementFind, func(rows *sql.Rows) error {

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
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindByCode, code, &artist.Id, &artist.Code, &artist.Name); err != nil {
		return nil, err
	}

	return &artist, nil
}

func (repository *RelationalArtistRepository) FindByName(ctx context.Context, name string) (*model.Artist, error) {

	var err error
	var artist model.Artist
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindByName, name, &artist.Id, &artist.Code, &artist.Name); err != nil {
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
