package serve

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	feather_boot "github.com/guidomantilla/go-feather-boot/pkg/boot"
	feather_commons_config "github.com/guidomantilla/go-feather-commons/pkg/config"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
	"github.com/spf13/cobra"
	"log/slog"
	"net/http"
	"os"

	"jukebox-app/internal/config"
	"jukebox-app/internal/repository"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	ctx := context.Background()
	appName, version := "jukebox-app", "1.0"

	builder := feather_boot.NewBeanBuilder(ctx)
	builder.Config = func(appCtx *feather_boot.ApplicationContext) {
		var cfg config.Config
		if err := feather_commons_config.Process(ctx, appCtx.Environment, &cfg); err != nil {
			slog.Error("starting up - error setting up configuration.", "message", err.Error())
			os.Exit(1)
		}

		appCtx.HttpConfig = &feather_boot.HttpConfig{
			Host: cfg.Host,
			Port: cfg.Port,
		}

		appCtx.SecurityConfig = &feather_boot.SecurityConfig{
			TokenSignatureKey:       cfg.TokenSignatureKey,
			PasswordMinSpecialChars: cfg.PasswordMinSpecialChars,
			PasswordMinNumber:       cfg.PasswordMinNumber,
			PasswordMinUpperCase:    cfg.PasswordMinUpperCase,
			PasswordLength:          cfg.PasswordLength,
		}

		appCtx.DatabaseConfig = &feather_boot.DatabaseConfig{
			ParamHolder:        feather_sql.UndefinedParamHolder.ValueFromName(*cfg.ParamHolder),
			Driver:             feather_sql.UndefinedDriverName.ValueFromName(*cfg.DatasourceDriver),
			DatasourceUrl:      cfg.DatasourceUrl,
			DatasourceServer:   cfg.DatasourceServer,
			DatasourceService:  cfg.DatasourceService,
			DatasourceUsername: cfg.DatasourceUsername,
			DatasourcePassword: cfg.DatasourcePassword,
		}
	}

	/*
		jukeboxApp.Attach("RabbitMQDispatcher", config.InitRabbitMQDispatcher(environment))
		jukeboxApp.Attach("NatsDispatcher", config.InitNatsDispatcher(environment))
	*/

	err := feather_boot.Init(appName, version, args, builder, func(appCtx feather_boot.ApplicationContext) error {

		cacheManager := config.InitCache(appCtx.Environment)
		userRepository := repository.NewRelationalUserRepository()
		_ = repository.NewCachedUserRepository(userRepository, cacheManager, json.Marshal, json.Unmarshal)
		_ = repository.NewRelationalArtistRepository()
		_ = repository.NewRelationalSongRepository()

		appCtx.PrivateRouter.GET("/info", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"appName": appName})
		})

		return nil
	})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
