package repository

import (
	"context"
	"database/sql"
	"jukebox-app/pkg/datasource"
	"jukebox-app/pkg/transaction"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_RelationalContext_Ok(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectExec(sqlStatement).WillReturnResult(sqlmock.NewResult(1, 1))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalContext(txCtx, sqlStatement, func(statement *sql.Stmt) error {

		if _, err = statement.Exec(nil); err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func Test_RelationalContext_Prepare_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement).WillReturnError(errors.New("some_error"))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalContext(txCtx, sqlStatement, func(statement *sql.Stmt) error {

		if _, err = statement.Exec(nil); err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalContext_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectExec(sqlStatement).WillReturnResult(sqlmock.NewResult(1, 1))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalContext(txCtx, sqlStatement, func(statement *sql.Stmt) error {

		if _, err = statement.Exec(nil); err != nil {
			return err
		}

		return errors.New("some_error")
	})

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}
