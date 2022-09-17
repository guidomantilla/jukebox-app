package serve

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/eko/gocache/v2/store"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"jukebox-app/internal/config"
	"jukebox-app/internal/repository"
	cachemanager "jukebox-app/pkg/cache-manager"
	"jukebox-app/pkg/transaction"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	// Create context that listens for the interrupt signal from the OS.
	ctx := context.Background()
	notifyCtx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer stop()

	//

	environment := config.InitConfig(&args)
	defer config.StopConfig()

	dataSource := config.InitDB(environment)
	defer config.StopDB()

	cacheInterface, _ := cachemanager.NewCache(store.GoCacheType, environment)
	cacheManager := cachemanager.NewDefaultCacheManager(cacheInterface)

	_ = transaction.NewRelationalTransactionHandler(dataSource)

	userRepository := repository.NewRelationalUserRepository()
	_ = repository.NewCachedUserRepository(userRepository, cacheManager)
	_ = repository.NewRelationalArtistRepository()
	_ = repository.NewRelationalSongRepository()

	//

	config.InitWebServer(environment)
	defer config.StopWebServer()

	<-notifyCtx.Done() // Listen for the interrupt signal.
	stop()             // Restore default behavior on the interrupt signal and notify user of shutdown.
	zap.L().Info("server shutting down gracefully, press Ctrl+C again to force")
}
