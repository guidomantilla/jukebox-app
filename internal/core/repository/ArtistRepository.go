package repository

import (
	"context"
	"database/sql"
	"fmt"
	"jukebox-app/internal/core/model"
	"jukebox-app/pkg/transaction"

	"go.uber.org/zap"
)

var _ ArtistRepository = (*DefaultArtistRepository)(nil)

/* TYPES DEFINITION */

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

//

type DefaultArtistRepository struct {
	statementCreate     string
	statementUpdate     string
	statementDelete     string
	statementFindById   string
	statementFind       string
	statementFindByCode string
	statementFindByName string
}

/* DefaultUserRepository METHODS */

func (repository *DefaultArtistRepository) Create(ctx context.Context, artist *model.Artist) error {

	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)

	var err error

	var statement *sql.Stmt
	if statement, err = tx.Prepare(repository.statementCreate); err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error("Error closing the statement")
		}
	}(statement)

	var result sql.Result
	if result, err = statement.Exec(artist.Code, artist.Name); err != nil {
		return err
	}

	if artist.Id, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (repository *DefaultArtistRepository) Update(ctx context.Context, artist *model.Artist) error {

	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)

	var err error

	var statement *sql.Stmt
	if statement, err = tx.Prepare(repository.statementUpdate); err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error("Error closing the statement")
		}
	}(statement)

	if _, err = statement.Exec(artist.Code, artist.Name, artist.Id); err != nil {
		return err
	}

	return nil
}

func (repository *DefaultArtistRepository) DeleteById(ctx context.Context, id int64) error {

	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)

	var err error

	var statement *sql.Stmt
	if statement, err = tx.Prepare(repository.statementDelete); err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error("Error closing the statement")
		}
	}(statement)

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository *DefaultArtistRepository) FindById(ctx context.Context, id int64) (*model.Artist, error) {

	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)

	var err error
	var statement *sql.Stmt

	if statement, err = tx.Prepare(repository.statementFindById); err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error("Error closing the statement")
		}
	}(statement)

	row := statement.QueryRow(id)

	var artist model.Artist
	if err = row.Scan(&artist.Id, &artist.Code, &artist.Name); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("artist with id %d not found", id)
		}
		return nil, err
	}

	return &artist, nil
}

func (repository *DefaultArtistRepository) FindAll(ctx context.Context) (*[]model.Artist, error) {

	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)

	var err error
	var statement *sql.Stmt

	if statement, err = tx.Prepare(repository.statementFind); err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error("Error closing the statement")
		}
	}(statement)

	var rows *sql.Rows
	if rows, err = statement.Query(); err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			zap.L().Error("Error closing the result set")
		}
	}(rows)

	artists := make([]model.Artist, 0)
	for rows.Next() {

		var artist model.Artist
		if err = rows.Scan(&artist.Id, &artist.Code, &artist.Name); err != nil {
			return nil, err
		}

		artists = append(artists, artist)
	}

	return &artists, nil
}

func (repository *DefaultArtistRepository) FindByCode(ctx context.Context, code int64) (*model.Artist, error) {

	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)

	var err error
	var statement *sql.Stmt

	if statement, err = tx.Prepare(repository.statementFindByCode); err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error("Error closing the statement")
		}
	}(statement)

	row := statement.QueryRow(code)

	var artist model.Artist
	if err = row.Scan(&artist.Id, &artist.Code, &artist.Name); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("artist with code %d not found", code)
		}
		return nil, err
	}

	return &artist, nil
}

func (repository *DefaultArtistRepository) FindByName(ctx context.Context, name string) (*model.Artist, error) {

	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)

	var err error
	var statement *sql.Stmt

	if statement, err = tx.Prepare(repository.statementFindByName); err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error("Error closing the statement")
		}
	}(statement)

	row := statement.QueryRow(name)

	var artist model.Artist
	if err = row.Scan(&artist.Id, &artist.Code, &artist.Name); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("artist with name %s not found", name)
		}
		return nil, err
	}

	return &artist, nil
}

/* TYPES CONSTRUCTOR */

func NewDefaultArtistRepository() *DefaultArtistRepository {
	return &DefaultArtistRepository{
		statementCreate:   "insert artist account (code, name) values (?, ?)",
		statementUpdate:   "update artist set code = ?, name = ? where id = ?",
		statementDelete:   "delete from artist where id = ?",
		statementFindById: "select id, code, name from artist where id = ?",
		statementFind:     "select id, code, name from artist",

		statementFindByCode: "select id, code, name from artist where code = ?",
		statementFindByName: "select id, code, name from artist where name = ?",
	}
}
