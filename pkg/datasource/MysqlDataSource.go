package datasource

import (
	"database/sql"
	"strings"

	"go.uber.org/zap"
)

var _ DBDataSource = (*MysqlDataSource)(nil)

type MysqlDataSource struct {
	driverName string
	username   string
	password   string
	url        string
	database   *sql.DB
	openFunc   func(driverName, dataSourceName string) (*sql.DB, error)
}

func NewMysqlDataSource(username string, password string, url string) *MysqlDataSource {

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)

	return &MysqlDataSource{
		driverName: "mysql",
		username:   username,
		password:   password,
		url:        url,
		database:   nil,
		openFunc:   sql.Open,
	}
}

func (mysqlDataSource *MysqlDataSource) GetDatabase() (*sql.DB, error) {

	var err error

	if mysqlDataSource.database == nil {
		if mysqlDataSource.database, err = mysqlDataSource.openFunc(mysqlDataSource.driverName, mysqlDataSource.url); err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
	}

	if err = mysqlDataSource.database.Ping(); err != nil {
		if mysqlDataSource.database, err = mysqlDataSource.openFunc(mysqlDataSource.driverName, mysqlDataSource.url); err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
	}

	return mysqlDataSource.database, nil
}
