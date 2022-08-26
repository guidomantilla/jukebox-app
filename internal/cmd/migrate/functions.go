package migrate

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"

	"jukebox-app/internal/config"
	"jukebox-app/pkg/datasource"
)

type MigrationFunction func(migration *migrate.Migrate) error

func createMigrateDriver(dataSource datasource.RelationalDataSource) (database.Driver, error) {

	var err error
	var db *sql.DB
	if db, err = dataSource.GetDatabase(); err != nil {
		return nil, err
	}

	var driver database.Driver

	if dataSource.GetDriverName() == datasource.MYSQL_DRIVER_NAME {
		if driver, err = mysql.WithInstance(db, &mysql.Config{}); err != nil {
			return nil, err
		}
	}

	if dataSource.GetDriverName() == datasource.POSTGRES_DRIVER_NAME {
		if driver, err = pgx.WithInstance(db, &pgx.Config{}); err != nil {
			return nil, err
		}
	}

	return driver, nil
}

func handleMigration(args []string, fn MigrationFunction) error {

	var err error

	env := config.InitConfig(&args)
	defer config.StopConfig()

	dataSource := config.InitDB(env)
	defer config.StopDB()

	var driver database.Driver
	if driver, err = createMigrateDriver(dataSource); err != nil {
		return err
	}

	workingDirectory, _ := os.Getwd()
	migrationsDirectory := filepath.Join(workingDirectory, "db/migrations/"+dataSource.GetDriverName())

	var migration *migrate.Migrate
	if migration, err = migrate.NewWithDatabaseInstance("file:///"+migrationsDirectory, dataSource.GetDriverName(), driver); err != nil {
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
