package applicationserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/qmdx00/lifecycle"
	"go.uber.org/zap"
)

var _ lifecycle.Server = (*HttpServer)(nil)

type HttpServer struct {
	internal *http.Server
}

func BuildHttpServer(server *http.Server) lifecycle.Server {
	return &HttpServer{
		internal: server,
	}
}

func (server *HttpServer) Run(ctx context.Context) error {

	info, _ := lifecycle.FromContext(ctx)
	zap.L().Info(fmt.Sprintf("server starting up - starting http server %s, v.%s", info.Name(), info.Version()))

	if err := server.internal.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Error(fmt.Sprintf("server starting up - error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *HttpServer) Stop(ctx context.Context) error {

	info, _ := lifecycle.FromContext(ctx)
	zap.L().Info(fmt.Sprintf("server shutting down - stopping http server %s, v.%s", info.Name(), info.Version()))

	if err := server.internal.Shutdown(ctx); err != nil {
		zap.L().Error(fmt.Sprintf("server shutting down - forced shutdown: %s", err.Error()))
		return err
	}

	zap.L().Info("server shutting down - http server stopped")
	return nil
}
