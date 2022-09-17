package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qmdx00/lifecycle"
	"go.uber.org/zap"
)

type AppServer struct {
	srv *http.Server
}

func NewAppServer() lifecycle.Server {
	router := gin.Default()
	router.GET("/sample", func(c *gin.Context) {
		log.Println("something")
	})

	return &AppServer{
		srv: &http.Server{
			Addr:    ":3000",
			Handler: router,
		},
	}
}

func (e *AppServer) Run(ctx context.Context) error {
	log.Printf("server starting up - starting http server")

	info, _ := lifecycle.FromContext(ctx)
	log.Printf(fmt.Sprintf("server %s, v.%s", info.Name(), info.Version()))

	if err := e.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf(fmt.Sprintf("server starting up - error: %s", err.Error()))
		return err
	}

	return nil
}

func (e *AppServer) Stop(ctx context.Context) error {
	log.Printf("server shutting down - stopping http server")

	info, _ := lifecycle.FromContext(ctx)
	log.Printf(fmt.Sprintf("server %s, v.%s", info.Name(), info.Version()))

	if err := e.srv.Shutdown(ctx); err != nil {
		log.Printf(fmt.Sprintf("server shutting down - forced shutdown: %s", err.Error()))
		return err
	}

	zap.L().Info("server shutting down - http server stopped")
	return nil
}

//

func main() {
	fmt.Println("Hello World")

	app := lifecycle.NewApp(
		lifecycle.WithName("test"),
		lifecycle.WithVersion("v1.0"),
	)

	app.Attach("app-server", NewAppServer())
	app.Cleanup(func() error {
		log.Println("do cleanup")
		return nil
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
