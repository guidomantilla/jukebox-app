package repository

import (
	"context"
	"database/sql"
	"errors"
	"jukebox-app/internal/core/model"
	"jukebox-app/pkg/datasource"
	"jukebox-app/pkg/transaction"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

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

//

type RelationalArtistRepository_FindAllFn func(ctx context.Context) (*[]model.Artist, error)
type RelationalArtistRepository_FindByInt64Fn func(ctx context.Context, n int64) (*model.Artist, error)
type RelationalArtistRepository_FindByStringFn func(ctx context.Context, s string) (*model.Artist, error)

func CallRelationalArtistRepositoryFindAllFunction(t *testing.T, statementFind string, findAllFn RelationalArtistRepository_FindAllFn, withQueryError bool, withScanError bool) (*[]model.Artist, error) {

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
				sqlmock.NewRows([]string{"id", "code", "name"}).
					AddRow("1", "101", "test01").
					AddRow("1", "a", "test02"), // makes the rows.scan(...) fail
			)
		} else {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name"}).
					AddRow("1", "101", "test01").
					AddRow("1", "102", "test02"),
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

func CallRelationalArtistRepositoryFindByInt64Function(t *testing.T, statementFind string, n int64, findByInt64Fn RelationalArtistRepository_FindByInt64Fn, withQueryError bool, withScanError bool) (*model.Artist, error) {

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
				sqlmock.NewRows([]string{"id", "code", "name"}).
					AddRow("1", "a", "test01"), // makes the rows.scan(...) fail
			)
		} else {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name"}).
					AddRow("1", "101", "test01"),
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

func CallRelationalArtistRepositoryFindByStringFnFunction(t *testing.T, statementFind string, s string, findByStringFn RelationalArtistRepository_FindByStringFn, withQueryError bool, withScanError bool) (*model.Artist, error) {

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
				sqlmock.NewRows([]string{"id", "code", "name"}).
					AddRow("1", "a", "test01"), // makes the rows.scan(...) fail
			)
		} else {
			expectQuery.WillReturnRows(
				sqlmock.NewRows([]string{"id", "code", "name"}).
					AddRow("1", "101", "test01"),
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

type RelationalArtistRepository_SaveFn func(_ context.Context, _ *model.Artist) error
type RelationalArtistRepository_DeleteFn func(ctx context.Context, id int64) error

func CallRelationalArtistRepositorySaveFunction(t *testing.T, statementCreate string, song *model.Artist, saveFn RelationalArtistRepository_SaveFn, withExecError bool) error {

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

func CallRelationalArtistRepositoryDeleteFunction(t *testing.T, statementCreate string, id int64, deleteFn RelationalArtistRepository_DeleteFn, withExecError bool) error {

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
