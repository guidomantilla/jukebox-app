package datasource

import (
	"database/sql"
	"fmt"
)

const (
	MYSQL_DRIVER_NAME    = "mysql"
	POSTGRES_DRIVER_NAME = "pgx"
)

var _ DBDataSource = (*MysqlDataSource)(nil)
var _ DBDataSource = (*PostgresDataSource)(nil)

type DBDataSource interface {
	GetDriverName() string
	GetDatabase() (*sql.DB, error)
}

func GetDBDataSourceFromDriverName(driverName string, username string, password string, url string) (DBDataSource, error) {

	if driverName == MYSQL_DRIVER_NAME {
		return NewMysqlDataSource(username, password, url), nil
	}

	if driverName == POSTGRES_DRIVER_NAME {
		return NewPostgresDataSource(username, password, url), nil
	}

	return nil, fmt.Errorf("wrong driver name")
}
