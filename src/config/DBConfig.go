package config

import (
	"jukebox-app/src/misc/datasource"
	"jukebox-app/src/misc/environment"

	"go.uber.org/zap"
)

const (
	DATASOURCE_USERNAME = "DATASOURCE_USERNAME"
	DATASOURCE_PASSWORD = "DATASOURCE_PASSWORD"
	DATASOURCE_URL      = "DATASOURCE_URL"
)

var singletonDataSource datasource.DBDataSource

func StopDB() {

	if err := singletonDataSource.GetDatabase().Close(); err != nil {
		zap.L().Error("Error closing DB")
		return
	}
}

func InitDB(environment environment.Environment) datasource.DBDataSource {
	username := environment.GetValue(DATASOURCE_USERNAME).AsString()
	password := environment.GetValue(DATASOURCE_PASSWORD).AsString()
	url := environment.GetValue(DATASOURCE_URL).AsString()
	singletonDataSource = datasource.NewMysqlDataSource(username, password, url)
	return singletonDataSource
}
