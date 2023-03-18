package config

import (
	"database/sql"
	"fmt"

	feather_relational_datasource "github.com/guidomantilla/go-feather-sql/pkg/feather-relational-datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/feather-sql"
	"go.uber.org/zap"

	"jukebox-app/pkg/environment"
)

var _driverNameMap = map[string]feather_sql.DriverName{
	feather_sql.PostgresDriverName.String(): feather_sql.PostgresDriverName,
	feather_sql.MysqlDriverName.String():    feather_sql.MysqlDriverName,
	feather_sql.OracleDriverName.String():   feather_sql.OracleDriverName,
}

var _singletonDatasource feather_relational_datasource.RelationalDatasource
var _datasourceContext feather_relational_datasource.RelationalDatasourceContext

func InitDB(environment environment.Environment) (feather_relational_datasource.RelationalDatasource, feather_relational_datasource.RelationalDatasourceContext) {

	zap.L().Info("server starting up - setting up DB connection")

	driver := environment.GetValue(DATASOURCE_DRIVER).AsString()
	if driver != feather_sql.PostgresDriverName.String() && driver != feather_sql.MysqlDriverName.String() && driver != feather_sql.OracleDriverName.String() {
		zap.L().Fatal("server starting up - error setting up DB connection: invalid driver name")
	}

	url := environment.GetValue(DATASOURCE_URL).AsString()
	if url == "" {
		zap.L().Fatal("server starting up - error setting up DB connection: invalid url")
	}

	username := environment.GetValue(DATASOURCE_USERNAME).AsString()
	password := environment.GetValue(DATASOURCE_PASSWORD).AsString()
	server := environment.GetValue(DATASOURCE_SERVER).AsString()
	service := environment.GetValue(DATASOURCE_SERVICE).AsString()

	_datasourceContext = feather_relational_datasource.BuildRelationalDatasourceContext(feather_sql.UnknownDriverName.ValueOf(driver), feather_sql.QuestionedParamHolder,
		url, username, password, server, service)

	_singletonDatasource = feather_relational_datasource.BuildRelationalDatasource(_datasourceContext, sql.Open)
	return _singletonDatasource, _datasourceContext
}

func StopDB() error {

	var err error
	var database *sql.DB

	zap.L().Info("server shutting down - closing DB")

	if database, err = _singletonDatasource.GetDatabase(); err != nil {
		zap.L().Error(fmt.Sprintf("server shutting down - error closing DB: %s", err.Error()))
		return err
	}

	if err = database.Close(); err != nil {
		zap.L().Error(fmt.Sprintf("server shutting down - error closing DB: %s", err.Error()))
		return err
	}

	zap.L().Info("server shutting down - DB closed")
	return nil
}
