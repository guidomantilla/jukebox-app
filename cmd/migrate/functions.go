package migrate

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	feather_relational_datasource "github.com/guidomantilla/go-feather-sql/pkg/feather-relational-datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/feather-sql"
	"github.com/spf13/cobra"

	"jukebox-app/internal/config"
)

type MigrationFunction func(migration *migrate.Migrate) error

func createMigrateDriver(datasource feather_relational_datasource.RelationalDatasource, datasourceContext feather_relational_datasource.RelationalDatasourceContext) (database.Driver, error) {

	var err error
	var db *sql.DB
	if db, err = datasource.GetDatabase(); err != nil {
		return nil, err
	}

	var driver database.Driver

	if datasourceContext.GetDriverName() == feather_sql.MysqlDriverName {
		if driver, err = mysql.WithInstance(db, &mysql.Config{}); err != nil {
			return nil, err
		}
	}

	if datasourceContext.GetDriverName() == feather_sql.PostgresDriverName {
		if driver, err = pgx.WithInstance(db, &pgx.Config{}); err != nil {
			return nil, err
		}
	}

	return driver, nil
}

func handleMigration(args []string, fn MigrationFunction) error {

	var err error

	env := config.InitConfig(&args)
	defer func() {
		_ = config.StopConfig()
	}()

	datasource, datasourceContext := config.InitDB(env)
	defer func() {
		_ = config.StopDB()
	}()

	var driver database.Driver
	if driver, err = createMigrateDriver(datasource, datasourceContext); err != nil {
		return err
	}

	workingDirectory, _ := os.Getwd()
	migrationsDirectory := filepath.Join(workingDirectory, "db/migrations/"+datasourceContext.GetDriverName().String())

	var migration *migrate.Migrate
	if migration, err = migrate.NewWithDatabaseInstance("file:///"+migrationsDirectory, datasourceContext.GetDriverName().String(), driver); err != nil {
		return err
	}

	if err = fn(migration); err != nil {
		return err
	}

	return nil
}

func UpCmdFn(_ *cobra.Command, args []string) {

	var err error
	err = handleMigration(args, func(migration *migrate.Migrate) error {

		if err = migration.Up(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err.Error())
	}
}

func DownCmdFn(_ *cobra.Command, args []string) {

	var err error
	err = handleMigration(args, func(migration *migrate.Migrate) error {

		if err = migration.Down(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err.Error())
	}
}

func DropCmdFn(_ *cobra.Command, args []string) {

	var err error
	err = handleMigration(args, func(migration *migrate.Migrate) error {

		if err = migration.Drop(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err.Error())
	}
}
