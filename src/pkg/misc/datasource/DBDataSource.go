package datasource

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type DBDataSource interface {
	GetDatabase() *sql.DB
}

type MysqlDataSource struct {
	username string
	password string
	url      string
	database *sql.DB
}

func NewMysqlDataSource(username string, password string, url string) *MysqlDataSource {

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)

	return &MysqlDataSource{
		username: username,
		password: password,
		url:      url,
		database: nil,
	}
}

func (mysqlDataSource *MysqlDataSource) GetDatabase() *sql.DB {

	var err error

	if mysqlDataSource.database == nil {
		if mysqlDataSource.database, err = open(mysqlDataSource.url); err != nil {
			zap.L().Error(err.Error())
		}
	}

	if err = mysqlDataSource.database.Ping(); err != nil {
		if mysqlDataSource.database, err = open(mysqlDataSource.url); err != nil {
			zap.L().Error(err.Error())
		}

		if err = mysqlDataSource.database.Ping(); err != nil {
			zap.L().Error(err.Error())
		}
	}

	return mysqlDataSource.database
}

func open(url string) (*sql.DB, error) {

	var err error
	var database *sql.DB

	if database, err = sql.Open("mysql", url); err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
