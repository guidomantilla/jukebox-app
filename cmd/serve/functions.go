package serve

import (
	"syscall"

	"github.com/qmdx00/lifecycle"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"jukebox-app/internal/config"
	"jukebox-app/internal/repository"
	"jukebox-app/pkg/transaction"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	//

	environment := config.InitConfig(&args)
	dataSource := config.InitDB(environment)
	cacheManager := config.InitCache(environment)

	//

	_ = transaction.NewRelationalTransactionHandler(dataSource)

	userRepository := repository.NewRelationalUserRepository()
	_ = repository.NewCachedUserRepository(userRepository, cacheManager)
	_ = repository.NewRelationalArtistRepository()
	_ = repository.NewRelationalSongRepository()

	//

	ginServer := config.InitGinServer(environment)

	//

	jukeboxApp := lifecycle.NewApp(
		lifecycle.WithName("jukebox-app"),
		lifecycle.WithVersion("v1.0"),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)
	jukeboxApp.Attach("ginServer", ginServer)
	jukeboxApp.Cleanup(config.StopDB, config.StopConfig)

	if err := jukeboxApp.Run(); err != nil {
		zap.L().Fatal(err.Error())
	}
}
