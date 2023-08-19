package handlers

import (
	grpc "best-goph-keeper/internal/api/proto"
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/storage/repositories/user"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	grpc.UnimplementedGophkeeperServer
	database *database.DB
	logger   *logrus.Logger
	user     *user.User
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *database.DB, userRepository *user.User, log *logrus.Logger) *Handler {
	return &Handler{database: db, user: userRepository, logger: log}
}
