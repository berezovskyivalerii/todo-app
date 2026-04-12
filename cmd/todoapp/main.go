package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
	corepgxpool "github.com/berezovskyivalerii/todo-app/internal/core/repository/postgres/pool/pgx"
	coremiddleware "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/middleware"
	coreserver "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/server"
	usersrepository "github.com/berezovskyivalerii/todo-app/internal/features/users/repository/postgres"
	usersservice "github.com/berezovskyivalerii/todo-app/internal/features/users/service"
	userhttp "github.com/berezovskyivalerii/todo-app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := corelogger.NewLogger(corelogger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger: %w", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing postgres connection pool")

	pool, err := corepgxpool.NewPool(ctx, corepgxpool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool: %w", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := usersrepository.NewUsersRepository(pool)
	usersService := usersservice.NewUsersService(usersRepository)
	usersTransportHTTP := userhttp.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing HTTP server")

	httpServer := coreserver.NewHTTPServer(
		coreserver.NewConfigMust(),
		logger,
		coremiddleware.RequestID(),
		coremiddleware.Logger(logger),
		coremiddleware.Trace(),
		coremiddleware.Panic(),
	)
	apiVersionRouterV1 := coreserver.NewAPIVersionRouter(coreserver.APIVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTransportHTTP.Routes()...)

	// apiVersionRouterV2 := coreserver.NewAPIVersionRouter(
	// 	coreserver.APIVersion2,
	// 	coremiddleware.Dummy("api v2 middleware"),
	// )
	// apiVersionRouterV2.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(
		apiVersionRouterV1,
	)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error: %w", zap.Error(err))
	}
}
