package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_DRIVER_NAME = "mysql"
)

var _ DBDataSource = (*MysqlDataSource)(nil)

type DBDataSource interface {
	GetDatabase() (*sql.DB, error)
}

func GetDBDataSourceFromDriverName(driverName string, username string, password string, url string) (DBDataSource, error) {

	if driverName == MYSQL_DRIVER_NAME {
		return NewMysqlDataSource(username, password, url), nil
	}

	return nil, fmt.Errorf("wrong driver name")
}
