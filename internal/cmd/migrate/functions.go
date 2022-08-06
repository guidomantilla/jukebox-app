package migrate

import (
	"database/sql"
	"jukebox-app/internal/config"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

type MigrationFunction func(migration *migrate.Migrate) error

func handleMigration(args []string, fn MigrationFunction) error {

	env := config.InitConfig(&args)
	defer config.StopConfig()

	dataSource := config.InitDB(env)
	defer config.StopDB()

	var err error
	var db *sql.DB
	var driver database.Driver

	db, err = dataSource.GetDatabase()

	if driver, err = mysql.WithInstance(db, &mysql.Config{}); err != nil {
		return err
	}

	workingDirectory, _ := os.Getwd()
	migrationsDirectory := filepath.Join(workingDirectory, "db/migrations")

	var migration *migrate.Migrate
	if migration, err = migrate.NewWithDatabaseInstance("file:///"+migrationsDirectory, "mysql", driver); err != nil {
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
