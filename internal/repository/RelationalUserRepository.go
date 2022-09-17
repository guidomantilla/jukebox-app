package repository

import (
	"context"
	"database/sql"

	"jukebox-app/internal/model"
	repositoryUtils "jukebox-app/pkg/repository"
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
	var id *int64
	if id, err = repositoryUtils.RelationalWriteContext(ctx, repository.statementCreate, user.Code, user.Name, user.Email); err != nil {
		return err
	}

	user.Id = *id

	return nil
}

func (repository *RelationalUserRepository) Update(ctx context.Context, user *model.User) error {

	var err error
	if _, err = repositoryUtils.RelationalWriteContext(ctx, repository.statementUpdate, user.Code, user.Name, user.Email, user.Id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalUserRepository) DeleteById(ctx context.Context, id int64) error {

	var err error
	if _, err = repositoryUtils.RelationalWriteContext(ctx, repository.statementDelete, id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalUserRepository) FindById(ctx context.Context, id int64) (*model.User, error) {

	var err error
	var user model.User
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindById, id, &user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *RelationalUserRepository) FindAll(ctx context.Context) (*[]model.User, error) {

	var err error
	users := make([]model.User, 0)
	err = repositoryUtils.RelationalQueryContext(ctx, repository.statementFind, func(rows *sql.Rows) error {

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
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindByCode, code, &user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *RelationalUserRepository) FindByName(ctx context.Context, name string) (*model.User, error) {

	var err error
	var user model.User
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindByName, name, &user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *RelationalUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {

	var err error
	var user model.User
	if err = repositoryUtils.RelationalQueryRowContext(ctx, repository.statementFindByEmail, email, &user.Id, &user.Code, &user.Name, &user.Email); err != nil {
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
