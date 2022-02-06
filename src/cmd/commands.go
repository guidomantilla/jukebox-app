package cmd

import (
	"jukebox-app/src/cmd/migrate"
	"jukebox-app/src/cmd/serve"
	"jukebox-app/src/cmd/test"
	"log"

	"github.com/spf13/cobra"
)

func ExecuteAppCmd() {

	appCmd := &cobra.Command{}
	appCmd.AddCommand(createServeCmd(), createMigrateCmd(), createTestCmd())

	if err := appCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}

func createServeCmd() *cobra.Command {

	serveCmd := &cobra.Command{
		Use: "serve",
		Run: serve.ExecuteCmdFn,
	}

	return serveCmd
}

func createMigrateCmd() *cobra.Command {

	migrateUpCmd := &cobra.Command{
		Use: "up",
		Run: migrate.UpCmdFn,
	}

	migrateDownCmd := &cobra.Command{
		Use: "down",
		Run: migrate.DownCmdFn,
	}

	migrateDropCmd := &cobra.Command{
		Use: "drop",
		Run: migrate.DropCmdFn,
	}

	migrateCmd := &cobra.Command{
		Use: "migrate",
	}

	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd, migrateDropCmd)

	return migrateCmd
}

func createTestCmd() *cobra.Command {

	testCmd := &cobra.Command{
		Use: "test",
		Run: test.ExecuteCmdFn,
	}

	return testCmd
}
