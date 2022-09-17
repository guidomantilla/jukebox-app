package serve

import (
	"context"
	"jukebox-app/internal/config"
	"jukebox-app/internal/repository"
	cachemanager "jukebox-app/pkg/cache-manager"
	"jukebox-app/pkg/transaction"
	"os"
	"os/signal"
	"syscall"

	"github.com/eko/gocache/v2/store"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	// Create context that listens for the interrupt signal from the OS.
	ctx := context.Background()
	notifyCtx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer stop()

	//

	environment := config.InitConfig(&args)
	dataSource := config.InitDB(environment)

	cacheInterface, _ := cachemanager.NewCache(store.GoCacheType, environment)
	cacheManager := cachemanager.NewDefaultCacheManager(cacheInterface)

	_ = transaction.NewRelationalTransactionHandler(dataSource)

	userRepository := repository.NewRelationalUserRepository()
	_ = repository.NewCachedUserRepository(userRepository, cacheManager)
	_ = repository.NewRelationalArtistRepository()
	_ = repository.NewRelationalSongRepository()

	//

	config.InitWebServer(environment)
	<-notifyCtx.Done() // Listen for the interrupt signal.
	stop()             // Restore default behavior on the interrupt signal and notify user of shutdown.
	zap.L().Info("server shutting down gracefully, press Ctrl+C again to force")

	config.StopDB()
	config.StopConfig()
	config.StopWebServer()

	zap.L().Info("server shutdown")

	//

	os.Exit(0)
}
