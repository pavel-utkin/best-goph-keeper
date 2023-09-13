package resthandler

import (
	"best-goph-keeper/internal/server/config"
	"best-goph-keeper/internal/server/database"
	"best-goph-keeper/internal/server/storage/repositories/token"
	"best-goph-keeper/internal/server/storage/repositories/user"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	database *database.DB
	config   *config.Config
	user     *user.User
	token    *token.Token
	log      *logrus.Logger
}

// NewHandler - creates a new server instance
func NewHandler(db *database.DB, config *config.Config, userRepository *user.User, tokenRepository *token.Token,
	log *logrus.Logger) *Handler {
	return &Handler{database: db, config: config, user: userRepository, token: tokenRepository, log: log}
}
