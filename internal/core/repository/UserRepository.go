package repository

import (
	"context"
	"database/sql"
	"fmt"
	"jukebox-app/internal/core/model"
	"jukebox-app/pkg/transaction"

	"go.uber.org/zap"
)

var _ UserRepository = (*DefaultUserRepository)(nil)

/* TYPES DEFINITION */

type UserRepository interface {
	Create(_ context.Context, _ *model.User) error
	Update(_ context.Context, _ *model.User) error
	DeleteById(_ context.Context, id int64) error
	FindById(_ context.Context, id int64) (*model.User, error)
	FindAll(_ context.Context) (*[]model.User, error)

	//Custom Finders

	FindByCode(_ context.Context, code int64) (*model.User, error)
	FindByName(_ context.Context, name string) (*model.User, error)
	FindByEmail(_ context.Context, email string) (*model.User, error)
}

type DefaultUserRepository struct {
	statementCreate      string
	statementUpdate      string
	statementDelete      string
	statementFindById    string
	statementFind        string
	statementFindByCode  string
	statementFindByName  string
	statementFindByEmail string
}

/* DefaultUserRepository METHODS */

func (repository *DefaultUserRepository) Create(ctx context.Context, user *model.User) error {

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
	if result, err = statement.Exec(user.Code, user.Name, user.Email); err != nil {
		return err
	}

	if user.Id, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (repository *DefaultUserRepository) Update(ctx context.Context, user *model.User) error {

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

	if _, err = statement.Exec(user.Code, user.Name, user.Email, user.Id); err != nil {
		return err
	}

	return nil
}

func (repository *DefaultUserRepository) DeleteById(ctx context.Context, id int64) error {

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

func (repository *DefaultUserRepository) FindById(ctx context.Context, id int64) (*model.User, error) {

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

	var user model.User
	if err = row.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	return &user, nil
}

func (repository *DefaultUserRepository) FindAll(ctx context.Context) (*[]model.User, error) {

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

	users := make([]model.User, 0)
	for rows.Next() {

		var user model.User
		if err = rows.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func (repository *DefaultUserRepository) FindByCode(ctx context.Context, code int64) (*model.User, error) {

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

	var user model.User
	if err = row.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("user with code %d not found", code)
		}
		return nil, err
	}

	return &user, nil
}

func (repository *DefaultUserRepository) FindByName(ctx context.Context, name string) (*model.User, error) {

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

	var user model.User
	if err = row.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("user with name %s not found", name)
		}
		return nil, err
	}

	return &user, nil
}

func (repository *DefaultUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {

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

	row := statement.QueryRow(email)

	var user model.User
	if err = row.Scan(&user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}

	return &user, nil
}

/* TYPES CONSTRUCTOR */

func NewDefaultUserRepository() *DefaultUserRepository {
	return &DefaultUserRepository{
		statementCreate:   "insert user account (code, name, email) values (?, ?, ?)",
		statementUpdate:   "update user set code = ?, name = ?, email = ? where id = ?",
		statementDelete:   "delete from user where id = ?",
		statementFindById: "select id, code, name, email from user where id = ?",
		statementFind:     "select id, code, name, email from user",

		statementFindByCode:  "select id, code, name, email from user where code = ?",
		statementFindByName:  "select id, code, name, email from user where name = ?",
		statementFindByEmail: "select id, code, name, email from user where email = ?",
	}
}
