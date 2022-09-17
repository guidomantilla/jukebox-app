package config

import (
	appserver "jukebox-app/pkg/application-server"
	"jukebox-app/pkg/environment"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qmdx00/lifecycle"
)

const (
	HOST_POST               = "HOST_POST"
	HOST_POST_DEFAULT_VALUE = ":8080"
)

func InitGinServer(environment environment.Environment) lifecycle.Server {

	//

	router := gin.Default()
	routerGroup := router.Group("/api")
	routerGroup.GET("/songs", nil)
	routerGroup.GET("/songs/:id", nil)
	routerGroup.POST("/songs", nil)
	routerGroup.PUT("/songs/:id", nil)
	routerGroup.PATCH("/songs/:id", nil)
	routerGroup.DELETE("/songs/:id", nil)

	//

	hostAddress := environment.GetValueOrDefault(HOST_POST, HOST_POST_DEFAULT_VALUE).AsString()
	httpServer := &http.Server{
		Addr:              hostAddress,
		Handler:           router,
		ReadHeaderTimeout: 60,
	}

	//

	ginServer := appserver.BuildHttpServer(httpServer)
	return ginServer
}
