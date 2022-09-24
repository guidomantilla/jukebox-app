package serve

import (
	"encoding/json"
	"syscall"
	"time"

	"github.com/qmdx00/lifecycle"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"jukebox-app/internal/config"
	"jukebox-app/internal/repository"
	appserver "jukebox-app/pkg/application-server"
	"jukebox-app/pkg/application-server/messaging"
	"jukebox-app/pkg/transaction"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	jukeboxApp := lifecycle.NewApp(
		lifecycle.WithName("jukebox-app"),
		lifecycle.WithVersion("v1.0"),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)
	jukeboxApp.Cleanup(config.StopDB, config.StopConfig)

	//

	environment := config.InitConfig(&args)
	dataSource := config.InitDB(environment)
	cacheManager := config.InitCache(environment)

	//

	_ = transaction.NewRelationalTransactionHandler(dataSource)

	userRepository := repository.NewRelationalUserRepository()
	_ = repository.NewCachedUserRepository(userRepository, cacheManager, json.Marshal, json.Unmarshal)
	_ = repository.NewRelationalArtistRepository()
	_ = repository.NewRelationalSongRepository()

	//

	client := messaging.NewDefaultRabbitMQQueueConnection("my-queue", "amqp://raven-dev:raven-dev*+@192.168.0.101:5672/", amqp.Dial)
	jukeboxApp.Attach("rabbitMQListener", appserver.BuildRabbitMQMessageDispatcher(client))
	<-time.After(time.Second)

	jukeboxApp.Attach("ginServer", config.InitGinServer(environment))

	//

	if err := jukeboxApp.Run(); err != nil {
		zap.L().Fatal(err.Error())
	}
}
