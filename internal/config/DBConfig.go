package config

import (
	"database/sql"
	"jukebox-app/pkg/datasource"
	"jukebox-app/pkg/environment"

	"go.uber.org/zap"
)

const (
	DATASOURCE_DRIVER   = "DATASOURCE_DRIVER"
	DATASOURCE_USERNAME = "DATASOURCE_USERNAME"
	DATASOURCE_PASSWORD = "DATASOURCE_PASSWORD"
	DATASOURCE_URL      = "DATASOURCE_URL"
)

var singletonDataSource datasource.RelationalDataSource

func InitDB(environment environment.Environment) datasource.RelationalDataSource {

	driver := environment.GetValue(DATASOURCE_DRIVER).AsString()
	if driver != datasource.POSTGRES_DRIVER_NAME && driver != datasource.MYSQL_DRIVER_NAME {
		zap.L().Fatal("invalid driver name")
	}

	username := environment.GetValue(DATASOURCE_USERNAME).AsString()
	password := environment.GetValue(DATASOURCE_PASSWORD).AsString()
	url := environment.GetValue(DATASOURCE_URL).AsString()

	singletonDataSource = datasource.NewRelationalDataSource(driver, username, password, url, sql.Open)
	return singletonDataSource
}

func StopDB() {

	var err error
	var database *sql.DB

	if database, err = singletonDataSource.GetDatabase(); err != nil {
		zap.L().Error("Error closing DB: " + err.Error())
		return
	}

	if err = database.Close(); err != nil {
		zap.L().Error("Error closing DB")
		return
	}
}
