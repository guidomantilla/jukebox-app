package datasource

import (
	"database/sql"
	"strings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	MYSQL_DRIVER_NAME    = "mysql"
	POSTGRES_DRIVER_NAME = "pgx"
)

type RelationalDataSource interface {
	GetDriverName() string
	GetDatabase() (*sql.DB, error)
}

type DefaultRelationalDataSource struct {
	driverName string
	username   string
	password   string
	url        string
	database   *sql.DB
	openFunc   func(driverName, dataSourceUrl string) (*sql.DB, error)
}

func NewRelationalDataSource(driverName string, username string, password string, url string) *DefaultRelationalDataSource {
	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)

	return &DefaultRelationalDataSource{
		driverName: driverName,
		username:   username,
		password:   password,
		url:        url,
		database:   nil,
		openFunc:   sql.Open,
	}
}

func (dataSource *DefaultRelationalDataSource) GetDriverName() string {
	return dataSource.driverName
}

func (dataSource *DefaultRelationalDataSource) GetDatabase() (*sql.DB, error) {

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
