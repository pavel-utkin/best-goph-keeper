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

// @Title Password Manager best-goph-keeper
// @Description GophKeeper is a client-server system that allows the user to safely and securely store logins, passwords, binary data and other private information.
// @Version 1.0

// @Contact.email pavel@utkin-pro.ru

func main() {
	logger := logrus.New()
	serverConfig := config.NewConfig(logger)
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

	handlerRest := resthandler.NewHandler(db, serverConfig, userRepository, tokenRepository, logger)
	routerService := router.Route(handlerRest)
	rs := chi.NewRouter()
	rs.Mount("/", routerService)

	handlerGrpc := grpchandler.NewHandler(db, serverConfig, userRepository, binaryRepository,
		&storage, entityRepository, tokenRepository, logger)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go api.StartGRPCService(handlerGrpc, serverConfig, logger)
	go api.StartRESTService(rs, serverConfig, logger)

	<-ctx.Done()
	logger.Info("server shutdown on signal with:", ctx.Err())
}
