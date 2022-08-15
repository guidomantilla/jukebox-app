package datasource

import (
	"database/sql"
	"strings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

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
		driverName: MYSQL_DRIVER_NAME,
		username:   username,
		password:   password,
		url:        url,
		database:   nil,
		openFunc:   sql.Open,
	}
}

func (dataSource *MysqlDataSource) GetDriverName() string {
	return dataSource.driverName
}

func (dataSource *MysqlDataSource) GetDatabase() (*sql.DB, error) {

	var err error

	if dataSource.database == nil {
		if dataSource.database, err = dataSource.openFunc(dataSource.driverName, dataSource.url); err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
	}

	if err = dataSource.database.Ping(); err != nil {
		if dataSource.database, err = dataSource.openFunc(dataSource.driverName, dataSource.url); err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
	}

	return dataSource.database, nil
}
