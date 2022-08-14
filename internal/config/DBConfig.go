package config

import (
	"database/sql"
	"jukebox-app/pkg/datasource"
	"jukebox-app/pkg/environment"

	"go.uber.org/zap"
)

const (
	DATASOURCE_USERNAME = "DATASOURCE_USERNAME"
	DATASOURCE_PASSWORD = "DATASOURCE_PASSWORD"
	DATASOURCE_URL      = "DATASOURCE_URL"
)

var singletonDataSource datasource.DBDataSource

func InitDB(environment environment.Environment) datasource.DBDataSource {
	username := environment.GetValue(DATASOURCE_USERNAME).AsString()
	password := environment.GetValue(DATASOURCE_PASSWORD).AsString()
	url := environment.GetValue(DATASOURCE_URL).AsString()

	var err error
	if singletonDataSource, err = datasource.GetDBDataSourceFromDriverName("mysql", username, password, url); err != nil {
		zap.L().Fatal(err.Error())
	}
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
