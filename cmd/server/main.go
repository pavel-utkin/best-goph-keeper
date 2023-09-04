package main

import (
	"best-goph-keeper/internal/server/api"
	grpchandler "best-goph-keeper/internal/server/api/grpc"
	resthandler "best-goph-keeper/internal/server/api/rest"
	"best-goph-keeper/internal/server/api/router"
	"best-goph-keeper/internal/server/config"
	"best-goph-keeper/internal/server/database"
	"best-goph-keeper/internal/server/storage"
	"best-goph-keeper/internal/server/storage/repositories/entity"
	"best-goph-keeper/internal/server/storage/repositories/file"
	"best-goph-keeper/internal/server/storage/repositories/token"
	"best-goph-keeper/internal/server/storage/repositories/user"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

func main() {
	logger := logrus.New()
	serverConfig := configserver.NewConfigServer(logger)
	logger.SetLevel(serverConfig.DebugLevel)

	db, err := database.New(serverConfig, logger)
	if err != nil {
		logger.Fatal(err)
	} else {
		defer db.Close()
		db.CreateTablesMigration("file://../migrations")
	}

	userRepository := user.New(db)
	binaryRepository := file.New(db)
	storage := storage.New("/tmp")
	entityRepository := entity.New(db)
	tokenRepository := token.New(db)

	handlerRest := resthandler.NewHandler(db, config, userRepository, tokenRepository, logger)
	routerService := router.Route(handlerRest)
	rs := chi.NewRouter()
	rs.Mount("/", routerService)

	handlerGrpc := grpchandler.NewHandler(db, config, userRepository, binaryRepository,
		&storage, entityRepository, tokenRepository, logger)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go api.StartGRPCService(handlerGrpc, config, logger)
	go api.StartRESTService(rs, config, logger)

	<-ctx.Done()
	logger.Info("server shutdown on signal with:", ctx.Err())
}
