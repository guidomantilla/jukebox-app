package serve

import (
	"jukebox-app/src/config"
	"jukebox-app/src/misc/transaction"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	environment := config.InitConfig(&args)
	defer config.StopConfig()

	dataSource := config.InitDB(environment)
	defer config.StopDB()

	_ = transaction.NewDefaultDBTransactionHandler(dataSource)

	if err := config.InitWebServer(environment); err != nil {
		zap.L().Fatal("error starting the server.")
	}
	defer config.StopWebServer()
}
