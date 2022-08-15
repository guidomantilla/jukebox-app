package datasource

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_NewRelationalDataSource(t *testing.T) {

	mysqlDataSource := NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password")

	assert.NotNil(t, mysqlDataSource)
	assert.Equal(t, "some_driver_name", mysqlDataSource.driverName)
	assert.Equal(t, "some_username", mysqlDataSource.username)
	assert.Equal(t, "some_password", mysqlDataSource.password)
	assert.Equal(t, "some_username_some_password", mysqlDataSource.url)
	assert.Nil(t, mysqlDataSource.database)
}

func Test_GetDriverName(t *testing.T) {

	mysqlDataSource := NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password")

	driver := mysqlDataSource.GetDriverName()
	assert.Equal(t, mysqlDataSource.driverName, driver)
}

func Test_GetDatabase_WhenDBIsNil_Ok(t *testing.T) {

	mysqlDataSource := NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password")
	mysqlDataSource.openFunc = func(driverName, dataSourceName string) (*sql.DB, error) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		return db, nil
	}

	database, err := mysqlDataSource.GetDatabase()

	assert.Nil(t, err)
	assert.NotNil(t, database)
	assert.NotNil(t, mysqlDataSource.database)
	assert.Equal(t, database, mysqlDataSource.database)
}

func Test_GetDatabase_WhenDBIsNil_Error(t *testing.T) {

	mysqlDataSource := NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password")
	mysqlDataSource.openFunc = func(driverName, dataSourceName string) (*sql.DB, error) {
		return nil, errors.New("some error")
	}

	database, err := mysqlDataSource.GetDatabase()
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, database)
	assert.Nil(t, mysqlDataSource.database)
}

func Test_GetDatabase_WhenDBIsNotNil_Ok(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mysqlDataSource := NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password")
	mysqlDataSource.database = db
	mysqlDataSource.openFunc = func(driverName, dataSourceName string) (*sql.DB, error) {
		return mysqlDataSource.database, nil
	}

	mock.ExpectPing()

	database, err := mysqlDataSource.GetDatabase()

	assert.Nil(t, err)
	assert.NotNil(t, database)
	assert.NotNil(t, mysqlDataSource.database)
	assert.Equal(t, database, mysqlDataSource.database)
}

func Test_GetDatabase_WhenDBIsNotNil_Error(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock

	if db, mock, err = sqlmock.New(sqlmock.MonitorPingsOption(true)); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mysqlDataSource := NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password")
	mysqlDataSource.database = db

	mysqlDataSource.openFunc = func(driverName, dataSourceName string) (*sql.DB, error) {

		return nil, errors.New("some error")
	}

	mock.ExpectPing().WillReturnError(errors.New("some error"))

	database, err := mysqlDataSource.GetDatabase()
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, database)
	assert.Nil(t, mysqlDataSource.database)
}
