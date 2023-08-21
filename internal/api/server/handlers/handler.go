package handlers

import (
	grpc "best-goph-keeper/internal/api/proto"
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/storage/repositories/text"
	"best-goph-keeper/internal/storage/repositories/user"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	database *database.DB
	user     *user.User
	text     *text.Text
	logger   *logrus.Logger
	grpc.UnimplementedGophkeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *database.DB, userRepository *user.User, textRepository *text.Text, log *logrus.Logger) *Handler {
	return &Handler{database: db, user: userRepository, text: textRepository, logger: log}
}
