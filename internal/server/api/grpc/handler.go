package grpchandler

import (
	"best-goph-keeper/internal/server/config"
	"best-goph-keeper/internal/server/database"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage"
	"best-goph-keeper/internal/server/storage/repositories/entity"
	"best-goph-keeper/internal/server/storage/repositories/file"
	"best-goph-keeper/internal/server/storage/repositories/token"
	"best-goph-keeper/internal/server/storage/repositories/user"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	database *database.DB
	config   *config.Config
	user     *user.User
	file     *file.File
	storage  *storage.Storage
	entity   *entity.Entity
	token    *token.Token
	logger   *logrus.Logger
	grpc.UnimplementedGophkeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *database.DB, config *config.Config, userRepository *user.User,
	binaryRepository *file.File, storage *storage.Storage, entityRepository *entity.Entity, tokenRepository *token.Token, log *logrus.Logger) *Handler {
	return &Handler{database: db, config: config, user: userRepository, file: binaryRepository, storage: storage,
		entity: entityRepository, token: tokenRepository, logger: log}
}
