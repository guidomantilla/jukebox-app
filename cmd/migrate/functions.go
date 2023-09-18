package migrate

import (
	"context"
	"database/sql"
	feather_boot "github.com/guidomantilla/go-feather-boot/pkg/boot"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/guidomantilla/go-feather-sql/pkg/datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
	"github.com/spf13/cobra"
)

type MigrationFunction func(migration *migrate.Migrate) error

func createMigrateDriver(datasource datasource.Datasource, datasourceContext datasource.DatasourceContext) (database.Driver, error) {

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

	ctx := context.Background()
	appName, version := "jukebox-app", "1.0"
	builder := feather_boot.NewBeanBuilder(ctx)

	err := feather_boot.Init(appName, version, args, builder, func(appCtx feather_boot.ApplicationContext) error {

		var err error
		var driver database.Driver
		if driver, err = createMigrateDriver(appCtx.Datasource, appCtx.DatasourceContext); err != nil {
			return err
		}

		workingDirectory, _ := os.Getwd()
		migrationsDirectory := filepath.Join(workingDirectory, "db/migrations/"+appCtx.DatasourceContext.GetDriverName().String())

		var migration *migrate.Migrate
		if migration, err = migrate.NewWithDatabaseInstance("file:///"+migrationsDirectory, appCtx.DatasourceContext.GetDriverName().String(), driver); err != nil {
			return err
		}

		if err = fn(migration); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
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
