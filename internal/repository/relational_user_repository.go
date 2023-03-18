package repository

import (
	"context"
	"database/sql"

	"github.com/guidomantilla/go-feather-sql/pkg/dao"

	"jukebox-app/internal/model"
)

type RelationalUserRepository struct {
	dao                  dao.CrudDao
	statementFindByCode  string
	statementFindByName  string
	statementFindByEmail string
}

/* TYPES CONSTRUCTOR */

func NewRelationalUserRepository() *RelationalUserRepository {
	return &RelationalUserRepository{
		statementFindByCode:  "select id, code, name, email from users where code = ?",
		statementFindByName:  "select id, code, name, email from users where name = ?",
		statementFindByEmail: "select id, code, name, email from users where email = ?",
	}
}

func (repository *RelationalUserRepository) Create(ctx context.Context, user *model.User) error {

	var err error
	var id *int64
	if id, err = repository.dao.Save(ctx, user.Code, user.Name, user.Email); err != nil {
		return err
	}

	user.Id = *id

	return nil
}

func (repository *RelationalUserRepository) Update(ctx context.Context, user *model.User) error {

	var err error
	if err = repository.dao.Update(ctx, user.Code, user.Name, user.Email, user.Id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalUserRepository) DeleteById(ctx context.Context, id int64) error {

	var err error
	if err = repository.dao.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (repository *RelationalUserRepository) FindById(ctx context.Context, id int64) (*model.User, error) {

	var err error
	var user model.User
	if err = repository.dao.FindById(ctx, id, &user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *RelationalUserRepository) FindAll(ctx context.Context) (*[]model.User, error) {

	var err error
	users := make([]model.User, 0)
	err = repository.dao.FindAll(ctx, func(rows *sql.Rows) error {

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
	if err = dao.ReadRowContext(ctx, repository.statementFindByCode, code, &user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *RelationalUserRepository) FindByName(ctx context.Context, name string) (*model.User, error) {

	var err error
	var user model.User
	if err = dao.ReadRowContext(ctx, repository.statementFindByName, name, &user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *RelationalUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {

	var err error
	var user model.User
	if err = dao.ReadRowContext(ctx, repository.statementFindByEmail, email, &user.Id, &user.Code, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
