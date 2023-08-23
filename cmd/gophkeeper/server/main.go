package main

import (
	"best-goph-keeper/internal/api/server"
	configserver "best-goph-keeper/internal/api/server/config"
	grpcHandler "best-goph-keeper/internal/api/server/handlers"
	"best-goph-keeper/internal/api/server/storage/repositories/metadata"
	"best-goph-keeper/internal/api/server/storage/repositories/text"
	"best-goph-keeper/internal/api/server/storage/repositories/token"
	"best-goph-keeper/internal/api/server/storage/repositories/user"
	"best-goph-keeper/internal/database"
	"context"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

func main() {
	logger := logrus.New()
	config := configserver.NewConfigServer(logger)
	logger.SetLevel(config.DebugLevel)

	db, err := database.New(config, logger)
	if err != nil {
		logger.Fatal(err)
	} else {
		defer db.Close()
	}

	userRepository := user.New(db)
	textRepository := text.New(db)
	metadataRepository := metadata.New(db)
	tokenRepository := token.New(db)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	handlerGrpc := grpcHandler.NewHandler(db, userRepository, textRepository, metadataRepository, tokenRepository, logger)
	go server.StartService(handlerGrpc, config, logger)

	<-ctx.Done()
	logger.Info("server shutdown on signal with:", ctx.Err())
}
