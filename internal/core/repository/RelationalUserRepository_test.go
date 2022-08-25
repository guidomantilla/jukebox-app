package repository

import (
	"context"
	"database/sql"
	"jukebox-app/internal/core/model"
	"jukebox-app/pkg/datasource"
	"jukebox-app/pkg/transaction"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_RelationalUserRepository_Create_Ok(t *testing.T) {

	user := &model.User{
		Id:    -1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositorySaveFunction(t, repository.statementCreate, user, repository.Create, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_RelationalUserRepository_Create_Err(t *testing.T) {

	user := &model.User{
		Id:    -1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositorySaveFunction(t, repository.statementCreate, user, repository.Create, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(-1), user.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalUserRepository_Update_Ok(t *testing.T) {

	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositorySaveFunction(t, repository.statementUpdate, user, repository.Update, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_RelationalUserRepository_Update_Err(t *testing.T) {

	user := &model.User{
		Id:    1,
		Code:  2,
		Name:  "3",
		Email: "4",
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositorySaveFunction(t, repository.statementUpdate, user, repository.Update, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), user.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalUserRepository_DeleteById_Ok(t *testing.T) {

	user := &model.User{
		Id: 1,
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositoryDeleteFunction(t, repository.statementDelete, user.Id, repository.DeleteById, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_RelationalUserRepository_DeleteById_Err(t *testing.T) {

	user := &model.User{
		Id: 1,
	}

	repository := NewRelationalUserRepository()
	err := CallRelationalUserRepositoryDeleteFunction(t, repository.statementDelete, user.Id, repository.DeleteById, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), user.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalUserRepository_FindAll_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	users, err := CallRelationalUserRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, users)
}

func Test_RelationalUserRepository_FindAll_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	users, err := CallRelationalUserRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindAll_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	users, err := CallRelationalUserRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalUserRepository_FindById_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalUserRepository_FindById_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindById_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalUserRepository_FindByCode_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalUserRepository_FindByCode_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindByCode_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalUserRepository_FindByName_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalUserRepository_FindByName_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindByName_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalUserRepository_FindByEmail_Ok(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByEmail, "some_email", repository.FindByEmail, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalUserRepository_FindByEmail_Query_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByEmail, "some_email", repository.FindByEmail, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalUserRepository_FindByEmail_Scan_Err(t *testing.T) {

	repository := NewRelationalUserRepository()
	user, err := CallRelationalUserRepositoryFindByStringFnFunction(t, repository.statementFindByEmail, "some_email", repository.FindByEmail, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

type RelationalUserRepository_FindAllFn func(ctx context.Context) (*[]model.User, error)
type RelationalUserRepository_FindByInt64Fn func(ctx context.Context, n int64) (*model.User, error)
type RelationalUserRepository_FindByStringFn func(ctx context.Context, s string) (*model.User, error)

func CallRelationalUserRepositoryFindAllFunction(t *testing.T, statementFind string, findAllFn RelationalUserRepository_FindAllFn, withQueryError bool, withScanError bool) (*[]model.User, error) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectPrepare(statementFind)
	expectQuery := mock.ExpectQuery(statementFind)

	if withQueryError {
		expectQuery.WillReturnError(errors.New("some_error"))
	} else {
		if withScanError {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name", "email"}).
					AddRow("1", "101", "test01", "test01@test.com").
					AddRow("1", "a", "test02", "test02@test.com"), // makes the rows.scan(...) fail
			)
		} else {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name", "email"}).
					AddRow("1", "101", "test01", "test01@test.com").
					AddRow("1", "102", "test02", "test02@test.com"),
			)
		}
	}

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	return findAllFn(txCtx)
}

func CallRelationalUserRepositoryFindByInt64Function(t *testing.T, statementFind string, n int64, findByInt64Fn RelationalUserRepository_FindByInt64Fn, withQueryError bool, withScanError bool) (*model.User, error) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectPrepare(statementFind)
	expectQuery := mock.ExpectQuery(statementFind)

	if withQueryError {
		expectQuery.WillReturnError(errors.New("some_error"))
	} else {
		if withScanError {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name", "email"}).
					AddRow("1", "a", "test01", "test01@test.com"), // makes the rows.scan(...) fail
			)
		} else {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name", "email"}).
					AddRow("1", "101", "test01", "test01@test.com"),
			)
		}
	}

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	return findByInt64Fn(txCtx, n)
}

func CallRelationalUserRepositoryFindByStringFnFunction(t *testing.T, statementFind string, s string, findByStringFn RelationalUserRepository_FindByStringFn, withQueryError bool, withScanError bool) (*model.User, error) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectPrepare(statementFind)
	expectQuery := mock.ExpectQuery(statementFind)

	if withQueryError {
		expectQuery.WillReturnError(errors.New("some_error"))
	} else {
		if withScanError {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name", "email"}).
					AddRow("1", "a", "test01", "test01@test.com"), // makes the rows.scan(...) fail
			)
		} else {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name", "email"}).
					AddRow("1", "101", "test01", "test01@test.com"),
			)
		}
	}

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	return findByStringFn(txCtx, s)
}

//

type RelationalUserRepository_SaveFn func(_ context.Context, _ *model.User) error
type RelationalUserRepository_DeleteFn func(ctx context.Context, id int64) error

func CallRelationalUserRepositorySaveFunction(t *testing.T, statementCreate string, user *model.User, saveFn RelationalUserRepository_SaveFn, withExecError bool) error {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectPrepare(statementCreate)
	expectExec := mock.ExpectExec(statementCreate)

	if withExecError {
		expectExec.WillReturnError(errors.New("some_error"))
	} else {
		expectExec.WillReturnResult(sqlmock.NewResult(1, 1))
	}

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	return saveFn(txCtx, user)
}

func CallRelationalUserRepositoryDeleteFunction(t *testing.T, statementCreate string, id int64, deleteFn RelationalUserRepository_DeleteFn, withExecError bool) error {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectPrepare(statementCreate)
	expectExec := mock.ExpectExec(statementCreate)

	if withExecError {
		expectExec.WillReturnError(errors.New("some_error"))
	} else {
		expectExec.WillReturnResult(sqlmock.NewResult(1, 1))
	}

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	return deleteFn(txCtx, id)
}
