package config

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	"jukebox-app/pkg/datasource"
	"jukebox-app/pkg/environment"
)

const (
	DATASOURCE_DRIVER   = "DATASOURCE_DRIVER"
	DATASOURCE_USERNAME = "DATASOURCE_USERNAME"
	DATASOURCE_PASSWORD = "DATASOURCE_PASSWORD"
	DATASOURCE_URL      = "DATASOURCE_URL"
)

var singletonDataSource datasource.RelationalDataSource

func InitDB(environment environment.Environment) datasource.RelationalDataSource {

	zap.L().Info("server starting up - setting up DB connection")

	driver := environment.GetValue(DATASOURCE_DRIVER).AsString()
	if driver != datasource.POSTGRES_DRIVER_NAME && driver != datasource.MYSQL_DRIVER_NAME {
		zap.L().Fatal("server starting up - error setting up DB connection: invalid driver name")
	}

	username := environment.GetValue(DATASOURCE_USERNAME).AsString()
	password := environment.GetValue(DATASOURCE_PASSWORD).AsString()
	url := environment.GetValue(DATASOURCE_URL).AsString()

	singletonDataSource = datasource.NewRelationalDataSource(driver, username, password, url, sql.Open)
	return singletonDataSource
}

func StopDB() error {

	var err error
	var database *sql.DB

	zap.L().Info("server shutting down - closing DB")

	if database, err = singletonDataSource.GetDatabase(); err != nil {
		zap.L().Error(fmt.Sprintf("server shutting down - error closing DB: %s", err.Error()))
		return err
	}

	if err = database.Close(); err != nil {
		zap.L().Error(fmt.Sprintf("server shutting down - error closing DB: %s", err.Error()))
		return err
	}

	zap.L().Info("server shutting down - DB closed")
	return nil
}
