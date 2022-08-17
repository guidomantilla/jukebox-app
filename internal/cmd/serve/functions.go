package serve

import (
	"jukebox-app/internal/config"
	"jukebox-app/internal/core/repository"
	"jukebox-app/pkg/transaction"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	environment := config.InitConfig(&args)
	defer config.StopConfig()

	dataSource := config.InitDB(environment)
	defer config.StopDB()

	_ = transaction.NewRelationalTransactionHandler(dataSource)

	_ = repository.NewRelationalUserRepository()
	_ = repository.NewRelationalArtistRepository()
	_ = repository.NewRelationalSongRepository()

	if err := config.InitWebServer(environment); err != nil {
		zap.L().Fatal("error starting the server.")
	}
	defer config.StopWebServer()
}
