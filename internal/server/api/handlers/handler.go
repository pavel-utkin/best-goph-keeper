package handlers

import (
	"best-goph-keeper/internal/database"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/repositories/metadata"
	"best-goph-keeper/internal/server/storage/repositories/text"
	"best-goph-keeper/internal/server/storage/repositories/token"
	"best-goph-keeper/internal/server/storage/repositories/user"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	database *database.DB
	user     *user.User
	text     *text.Text
	metadata *metadata.Metadata
	token    *token.Token
	logger   *logrus.Logger
	grpc.UnimplementedGophkeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *database.DB, userRepository *user.User, textRepository *text.Text, metadataRepository *metadata.Metadata, tokenRepository *token.Token, log *logrus.Logger) *Handler {
	return &Handler{database: db, user: userRepository, text: textRepository, metadata: metadataRepository, token: tokenRepository, logger: log}
}
