package repository

import (
	"context"
	"database/sql"
	"fmt"
	"jukebox-app/internal/core/model"
	repositoryUtils "jukebox-app/pkg/repository"

	"go.uber.org/zap"
)

type RelationalUserRepository struct {
	statementCreate      string
	statementUpdate      string
	statementDelete      string
	statementFindById    string
	statementFind        string
	statementFindByCode  string
	statementFindByName  string
	statementFindByEmail string
}

func (repository *RelationalUserRepository) Create(ctx context.Context, user *model.User) error {

	var err error
	err = repositoryUtils.RelationalContext(ctx, repository.statementCreate, func(statement *sql.Stmt) error {

		var result sql.Result
		if result, err = statement.Exec(user.Code, user.Name, user.Email); err != nil {
			return err
		}

		if user.Id, err = result.LastInsertId(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *RelationalUserRepository) Update(ctx context.Context, user *model.User) error {

	var err error
	err = repositoryUtils.RelationalContext(ctx, repository.statementUpdate, func(statement *sql.Stmt) error {

		if _, err = statement.Exec(user.Code, user.Name, user.Email, user.Id); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *RelationalUserRepository) DeleteById(ctx context.Context, id int64) error {

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

func (repository *RelationalUserRepository) FindById(ctx context.Context, id int64) (*model.User, error) {

	var err error
	var user model.User
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(id)
		if err = row.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("user with id %d not found", id)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *RelationalUserRepository) FindAll(ctx context.Context) (*[]model.User, error) {

	var err error
	users := make([]model.User, 0)
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
			var user model.User
			if err = rows.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
				return err
			}
			users = append(users, user)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (repository *RelationalUserRepository) FindByCode(ctx context.Context, code int64) (*model.User, error) {

	var err error
	var user model.User
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(code)
		if err = row.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("user with code %d not found", code)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *RelationalUserRepository) FindByName(ctx context.Context, name string) (*model.User, error) {

	var err error
	var user model.User
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(name)
		if err = row.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("user with name %s not found", name)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *RelationalUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {

	var err error
	var user model.User
	err = repositoryUtils.RelationalContext(ctx, repository.statementDelete, func(statement *sql.Stmt) error {

		row := statement.QueryRow(email)
		if err = row.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("user with email %s not found", email)
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/* TYPES CONSTRUCTOR */

func NewRelationalUserRepository() *RelationalUserRepository {
	return &RelationalUserRepository{
		statementCreate:   "insert into users (code, name, email) values (?, ?, ?)",
		statementUpdate:   "update users set code = ?, name = ?, email = ? where id = ?",
		statementDelete:   "delete from users where id = ?",
		statementFindById: "select id, code, name, email from users where id = ?",
		statementFind:     "select id, code, name, email from users",

		statementFindByCode:  "select id, code, name, email from users where code = ?",
		statementFindByName:  "select id, code, name, email from users where name = ?",
		statementFindByEmail: "select id, code, name, email from users where email = ?",
	}
}
