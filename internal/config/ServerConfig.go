package config

import (
	"github.com/gin-gonic/gin"

	"jukebox-app/pkg/environment"
)

const (
	HOST_POST               = "HOST_POST"
	HOST_POST_DEFAULT_VALUE = ":8080"
)

var _singletonEngine *gin.Engine

func StopWebServer() {
	//Nothing to do here yet
}

func InitWebServer(environment environment.Environment) error {

	_singletonEngine = gin.Default()

	loadApiRoutes(nil)

	hostAddress := environment.GetValueOrDefault(HOST_POST, HOST_POST_DEFAULT_VALUE).AsString()
	return _singletonEngine.Run(hostAddress)
}

func loadApiRoutes(something any) {

	group := _singletonEngine.Group("/api")
	group.GET("/songs", nil)
	group.GET("/songs/:id", nil)
	group.POST("/songs", nil)
	group.PUT("/songs/:id", nil)
	group.PATCH("/songs/:id", nil)
	group.DELETE("/songs/:id", nil)
}
