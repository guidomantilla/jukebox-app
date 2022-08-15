package repository

import (
	"context"
	"database/sql"
	"fmt"
	"jukebox-app/internal/core/model"
	"jukebox-app/pkg/transaction"

	"go.uber.org/zap"
)

var _ SongRepository = (*DefaultSongRepository)(nil)

/* TYPES DEFINITION */

type SongRepository interface {
	Create(_ context.Context, _ *model.Song) error
	Update(_ context.Context, _ *model.Song) error
	DeleteById(_ context.Context, id int64) error
	FindById(_ context.Context, id int64) (*model.Song, error)
	FindAll(_ context.Context) (*[]model.Song, error)

	//Custom Finders

	FindByCode(_ context.Context, code int64) (*model.Song, error)
	FindByName(_ context.Context, name string) (*model.Song, error)
	FindByArtistId(_ context.Context, id string) (*[]model.Song, error)
}

type DefaultSongRepository struct {
	statementCreate         string
	statementUpdate         string
	statementDelete         string
	statementFindById       string
	statementFind           string
	statementFindByCode     string
	statementFindByName     string
	statementFindByArtistId string
}

/* DefaultUserRepository METHODS */

func (repository *DefaultSongRepository) Create(ctx context.Context, song *model.Song) error {

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
	if result, err = statement.Exec(song.Code, song.Name, song.ArtistId); err != nil {
		return err
	}

	if song.Id, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (repository *DefaultSongRepository) Update(ctx context.Context, song *model.Song) error {

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

	if _, err = statement.Exec(song.Code, song.Name, song.ArtistId, song.Id); err != nil {
		return err
	}

	return nil
}

func (repository *DefaultSongRepository) DeleteById(ctx context.Context, id int64) error {

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

func (repository *DefaultSongRepository) FindById(ctx context.Context, id int64) (*model.Song, error) {

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

	var song model.Song
	if err = row.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("song with id %d not found", id)
		}
		return nil, err
	}

	return &song, nil
}

func (repository *DefaultSongRepository) FindAll(ctx context.Context) (*[]model.Song, error) {

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

	songs := make([]model.Song, 0)
	for rows.Next() {

		var song model.Song
		if err = rows.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	return &songs, nil
}

func (repository *DefaultSongRepository) FindByCode(ctx context.Context, code int64) (*model.Song, error) {

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

	var song model.Song
	if err = row.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("song with code %d not found", code)
		}
		return nil, err
	}

	return &song, nil
}

func (repository *DefaultSongRepository) FindByName(ctx context.Context, name string) (*model.Song, error) {

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

	var song model.Song
	if err = row.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("song with name %s not found", name)
		}
		return nil, err
	}

	return &song, nil
}

func (repository *DefaultSongRepository) FindByArtistId(ctx context.Context, id string) (*[]model.Song, error) {

	var tx = ctx.Value(transaction.RelationalTransactionContext{}).(*sql.Tx)

	var err error
	var statement *sql.Stmt

	if statement, err = tx.Prepare(repository.statementFindByArtistId); err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			zap.L().Error("Error closing the statement")
		}
	}(statement)

	var rows *sql.Rows
	if rows, err = statement.Query(id); err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			zap.L().Error("Error closing the result set")
		}
	}(rows)

	songs := make([]model.Song, 0)
	for rows.Next() {

		var song model.Song
		if err = rows.Scan(&song.Id, &song.Code, &song.Name, &song.ArtistId); err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	return &songs, nil
}

/* TYPES CONSTRUCTOR */

func NewDefaultSongRepository() *DefaultSongRepository {
	return &DefaultSongRepository{
		statementCreate:   "insert song account (code, name, artistId) values (?, ?, ?)",
		statementUpdate:   "update song set code = ?, name = ?, artistId = ? where id = ?",
		statementDelete:   "delete from song where id = ?",
		statementFindById: "select id, code, name, artistId from song where id = ?",
		statementFind:     "select id, code, name, artistId from song",

		statementFindByCode:     "select id, code, name, artistId from song where code = ?",
		statementFindByName:     "select id, code, name, artistId from song where name = ?",
		statementFindByArtistId: "select id, code, name, artistId from song where artistId = ?",
	}
}
