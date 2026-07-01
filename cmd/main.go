package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_logger "github.com/Vlad-6894/test_task/internal/core/logger"
	core_postgres_pool "github.com/Vlad-6894/test_task/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/Vlad-6894/test_task/internal/core/transport/http/middleware"
	core_http_server "github.com/Vlad-6894/test_task/internal/core/transport/http/server"
	subscriptions_postgres_repository "github.com/Vlad-6894/test_task/internal/features/subscriptions/repository/postgres"
	subscriptions_service "github.com/Vlad-6894/test_task/internal/features/subscriptions/service"
	subscriptions_transport_http "github.com/Vlad-6894/test_task/internal/features/subscriptions/transport/http"
	"go.uber.org/zap"

	_ "github.com/Vlad-6894/test_task/docs"
)

var (
	timeZone = time.UTC
)

// @title Test task API
// @version 1.0
// @description API scheme
// @host 127.0.0.1:5050
// @BasePath /api/v1
func main() {
	time.Local = timeZone

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewLoggerConfigMust())
	if err != nil {
		fmt.Println("init application logger error!")
		os.Exit(1)
	}
	defer logger.Close()

	logger.Info("golang application time zone:", zap.Any("zone", timeZone))

	logger.Info("Init connection pool!")
	pool, err := core_postgres_pool.NewConnectionPool(ctx, core_postgres_pool.NewConnectionPoolConfigMust())
	if err != nil {
		logger.Fatal("Failed to init connection pool!", zap.Error(err))
	}
	defer pool.Close()

	logger.Info("Init feature subscriptions!")
	subscriptionsRepository := subscriptions_postgres_repository.NewSubscriptionRepository(pool)
	subscriptionsService := subscriptions_service.NewSubscriptionService(subscriptionsRepository)
	subscriptionTransportHTTP := subscriptions_transport_http.NewSubscriptionsHTTPHandler(subscriptionsService)

	logger.Info("Init HTTP server")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigHTTPServerMust(),
		logger,
		core_http_middleware.CORS(),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisteRoutes(subscriptionTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error:", zap.Error(err))
	}
}
