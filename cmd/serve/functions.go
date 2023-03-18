package serve

import (
	"encoding/json"
	"syscall"

	feather_config "github.com/guidomantilla/go-feather-sql/pkg/config"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
	"github.com/guidomantilla/go-feather-sql/pkg/transaction"
	"github.com/qmdx00/lifecycle"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"jukebox-app/internal/config"
	"jukebox-app/internal/repository"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	jukeboxApp := lifecycle.NewApp(
		lifecycle.WithName("jukebox-app"),
		lifecycle.WithVersion("1.0"),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)
	jukeboxApp.Cleanup(feather_config.Stop, config.StopConfig)

	//

	environment := config.InitConfig(&args)
	datasource, _ := feather_config.Init("", environment, feather_sql.QuestionedParamHolder)
	cacheManager := config.InitCache(environment)

	//

	_ = transaction.BuildRelationalTransactionHandler(datasource)

	userRepository := repository.NewRelationalUserRepository()
	_ = repository.NewCachedUserRepository(userRepository, cacheManager, json.Marshal, json.Unmarshal)
	_ = repository.NewRelationalArtistRepository()
	_ = repository.NewRelationalSongRepository()

	//

	jukeboxApp.Attach("RabbitMQDispatcher", config.InitRabbitMQDispatcher(environment))
	jukeboxApp.Attach("NatsDispatcher", config.InitNatsDispatcher(environment))
	jukeboxApp.Attach("GinServer", config.InitGinServer(environment))

	//

	if err := jukeboxApp.Run(); err != nil {
		zap.L().Fatal(err.Error())
	}
}
