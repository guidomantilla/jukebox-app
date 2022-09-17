package config

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"jukebox-app/pkg/environment"
)

const (
	HOST_POST               = "HOST_POST"
	HOST_POST_DEFAULT_VALUE = ":8080"
)

var _singletonServer *http.Server

func InitWebServer(environment environment.Environment) *http.Server {

	zap.L().Info("server starting up - starting http server")

	router := gin.Default()
	routerGroup := router.Group("/api")
	routerGroup.GET("/songs", nil)
	routerGroup.GET("/songs/:id", nil)
	routerGroup.POST("/songs", nil)
	routerGroup.PUT("/songs/:id", nil)
	routerGroup.PATCH("/songs/:id", nil)
	routerGroup.DELETE("/songs/:id", nil)

	hostAddress := environment.GetValueOrDefault(HOST_POST, HOST_POST_DEFAULT_VALUE).AsString()
	_singletonServer = &http.Server{
		Addr:    hostAddress,
		Handler: router,
	}

	go func() {
		if err := _singletonServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal(fmt.Sprintf("server starting up - error: %s", err.Error()))
		}
	}()

	return _singletonServer
}

func StopWebServer() {

	zap.L().Info("server shutting down - stopping http server")

	// The context is used to inform the server it has 5 seconds to finish the request it is currently handling
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := _singletonServer.Shutdown(timeoutCtx); err != nil {
		zap.L().Fatal(fmt.Sprintf("server shutting down - forced shutdown: %s", err.Error()))
	}

	zap.L().Info("server shutting down - http server stopped")
}
