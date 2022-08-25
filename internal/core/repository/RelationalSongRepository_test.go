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

func Test_RelationalSongRepository_Create_Ok(t *testing.T) {

	song := &model.Song{
		Id:       -1,
		Code:     1,
		Name:     "2",
		ArtistId: 3,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositorySaveFunction(t, repository.statementCreate, song, repository.Create, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), song.Id)
}

func Test_RelationalSongRepository_Create_Err(t *testing.T) {

	song := &model.Song{
		Id:       -1,
		Code:     1,
		Name:     "2",
		ArtistId: 3,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositorySaveFunction(t, repository.statementCreate, song, repository.Create, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(-1), song.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalSongRepository_Update_Ok(t *testing.T) {

	song := &model.Song{
		Id:       1,
		Code:     2,
		Name:     "3",
		ArtistId: 4,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositorySaveFunction(t, repository.statementUpdate, song, repository.Update, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), song.Id)
}

func Test_RelationalSongRepository_Update_Err(t *testing.T) {

	song := &model.Song{
		Id:       1,
		Code:     2,
		Name:     "3",
		ArtistId: 4,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositorySaveFunction(t, repository.statementUpdate, song, repository.Update, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), song.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalSongRepository_DeleteById_Ok(t *testing.T) {

	user := &model.User{
		Id: 1,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositoryDeleteFunction(t, repository.statementDelete, user.Id, repository.DeleteById, false)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.Id)
}

func Test_RelationalSongRepository_DeleteById_Err(t *testing.T) {

	user := &model.User{
		Id: 1,
	}

	repository := NewRelationalSongRepository()
	err := CallRelationalSongRepositoryDeleteFunction(t, repository.statementDelete, user.Id, repository.DeleteById, true)

	assert.NotNil(t, err)
	assert.Equal(t, int64(1), user.Id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalSongRepository_FindAll_Ok(t *testing.T) {

	repository := NewRelationalSongRepository()
	users, err := CallRelationalSongRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, users)
}

func Test_RelationalSongRepository_FindAll_Query_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	users, err := CallRelationalSongRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalSongRepository_FindAll_Scan_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	users, err := CallRelationalSongRepositoryFindAllFunction(t, repository.statementFind, repository.FindAll, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, users)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalSongRepository_FindById_Ok(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalSongRepository_FindById_Query_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalSongRepository_FindById_Scan_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindById, 1, repository.FindById, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalSongRepository_FindByCode_Ok(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalSongRepository_FindByCode_Query_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalSongRepository_FindByCode_Scan_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByInt64Function(t, repository.statementFindByCode, 1, repository.FindByCode, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

func Test_RelationalSongRepository_FindByName_Ok(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
}

func Test_RelationalSongRepository_FindByName_Query_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, true, false)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalSongRepository_FindByName_Scan_Err(t *testing.T) {

	repository := NewRelationalSongRepository()
	user, err := CallRelationalSongRepositoryFindByStringFnFunction(t, repository.statementFindByName, "some_name", repository.FindByName, false, true)

	assert.NotNil(t, err)
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.True(t, strings.Index(err.Error(), "sql: Scan ") == 0)
}

//

type RelationalSongRepository_FindAllFn func(ctx context.Context) (*[]model.Song, error)
type RelationalSongRepository_FindByInt64Fn func(ctx context.Context, n int64) (*model.Song, error)
type RelationalSongRepository_FindByStringFn func(ctx context.Context, s string) (*model.Song, error)

func CallRelationalSongRepositoryFindAllFunction(t *testing.T, statementFind string, findAllFn RelationalSongRepository_FindAllFn, withQueryError bool, withScanError bool) (*[]model.Song, error) {

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
				sqlmock.NewRows([]string{"id", "code", "name", "artistId"}).
					AddRow("1", "101", "test01", "201").
					AddRow("1", "a", "test02", "202"), // makes the rows.scan(...) fail
			)
		} else {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name", "artistId"}).
					AddRow("1", "101", "test01", "201").
					AddRow("1", "102", "test02", "202"),
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

func CallRelationalSongRepositoryFindByInt64Function(t *testing.T, statementFind string, n int64, findByInt64Fn RelationalSongRepository_FindByInt64Fn, withQueryError bool, withScanError bool) (*model.Song, error) {

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
				sqlmock.NewRows([]string{"id", "code", "name", "artistId"}).
					AddRow("1", "a", "test01", "201"), // makes the rows.scan(...) fail
			)
		} else {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name", "artistId"}).
					AddRow("1", "101", "test01", "201"),
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

func CallRelationalSongRepositoryFindByStringFnFunction(t *testing.T, statementFind string, s string, findByStringFn RelationalSongRepository_FindByStringFn, withQueryError bool, withScanError bool) (*model.Song, error) {

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
				sqlmock.NewRows([]string{"id", "code", "name", "artistId"}).
					AddRow("1", "a", "test01", "201"), // makes the rows.scan(...) fail
			)
		} else {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name", "artistId"}).
					AddRow("1", "101", "test01", "201"),
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

type RelationalSongRepository_SaveFn func(_ context.Context, _ *model.Song) error
type RelationalSongRepository_DeleteFn func(ctx context.Context, id int64) error

func CallRelationalSongRepositorySaveFunction(t *testing.T, statementCreate string, song *model.Song, saveFn RelationalSongRepository_SaveFn, withExecError bool) error {

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

	return saveFn(txCtx, song)
}

func CallRelationalSongRepositoryDeleteFunction(t *testing.T, statementCreate string, id int64, deleteFn RelationalSongRepository_DeleteFn, withExecError bool) error {

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
