package repository

import (
	"context"
	"database/sql"

	feather_relational_dao "github.com/guidomantilla/go-feather-sql/pkg/feather-relational-dao"

	"jukebox-app/internal/model"
)

type RelationalArtistRepository struct {
	dao                 feather_relational_dao.CrudDao
	statementFindByCode string
	statementFindByName string
}

/* TYPES CONSTRUCTOR */

func NewRelationalArtistRepository() *RelationalArtistRepository {
	return &RelationalArtistRepository{
		statementFindByCode: "select id, code, name from artists where code = ?",
		statementFindByName: "select id, code, name from artists where name = ?",
	}
}

func (repository *RelationalArtistRepository) Create(ctx context.Context, artist *model.Artist) error {

	var err error
	var id *int64
	if id, err = repository.dao.Save(ctx, artist.Code, artist.Name); err != nil {
		return err
	}

	artist.Id = *id

	return nil
}

func (repository *RelationalArtistRepository) Update(ctx context.Context, artist *model.Artist) error {

	var err error
	if err = repository.dao.Update(ctx, artist.Code, artist.Name, artist.Id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalArtistRepository) DeleteById(ctx context.Context, id int64) error {

	var err error
	if err = repository.dao.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalArtistRepository) FindById(ctx context.Context, id int64) (*model.Artist, error) {

	var err error
	var artist model.Artist
	if err = repository.dao.FindById(ctx, id, &artist.Id, &artist.Code, &artist.Name); err != nil {
		return nil, err
	}

	return &artist, nil
}

func (repository *RelationalArtistRepository) FindAll(ctx context.Context) (*[]model.Artist, error) {

	var err error
	artists := make([]model.Artist, 0)
	err = repository.dao.FindAll(ctx, func(rows *sql.Rows) error {

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
	if err = feather_relational_dao.ReadRowContext(ctx, repository.statementFindByCode, code, &artist.Id, &artist.Code, &artist.Name); err != nil {
		return nil, err
	}

	return &artist, nil
}

func (repository *RelationalArtistRepository) FindByName(ctx context.Context, name string) (*model.Artist, error) {

	var err error
	var artist model.Artist
	if err = feather_relational_dao.ReadRowContext(ctx, repository.statementFindByName, name, &artist.Id, &artist.Code, &artist.Name); err != nil {
		return nil, err
	}

	return &artist, nil
}
