package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_http_middleware "github.com/Vlad-6894/test_task/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux        *http.ServeMux
	config     ConfigHTTPServer
	logger     *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	config ConfigHTTPServer,
	logger *core_logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		logger:     logger,
		middleware: middleware,
	}
}

func (h *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(prefix+"/", http.StripPrefix(prefix, router))
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middleware...)

	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}

	channelErrors := make(chan error, 1)

	go func() {
		defer close(channelErrors)

		h.logger.Warn("http server started!", zap.String("addr", h.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			channelErrors <- err
		}
	}()

	select {
	case <-ctx.Done():
		h.logger.Info("HTTP server shutdown!")

		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), h.config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctxWithTimeout); err != nil {
			server.Close()

			return fmt.Errorf("Shutdown HTTP server error: %w", err)
		}

		h.logger.Info("HTTP server stopped!")
	case err := <-channelErrors:
		if err != nil {
			return fmt.Errorf("Server error: %w", err)
		}
	}
	return nil
}
