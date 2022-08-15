package datasource

import (
	"database/sql"
	"strings"

	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresDataSource struct {
	driverName string
	username   string
	password   string
	url        string
	database   *sql.DB
	openFunc   func(driverName, dataSourceName string) (*sql.DB, error)
}

func NewPostgresDataSource(username string, password string, url string) *PostgresDataSource {
	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)

	return &PostgresDataSource{
		driverName: POSTGRES_DRIVER_NAME,
		username:   username,
		password:   password,
		url:        url,
		database:   nil,
		openFunc:   sql.Open,
	}
}

func (dataSource *PostgresDataSource) GetDriverName() string {
	return dataSource.driverName
}

func (dataSource *PostgresDataSource) GetDatabase() (*sql.DB, error) {

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
