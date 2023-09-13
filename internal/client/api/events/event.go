package events

import (
	"best-goph-keeper/internal/client/config"
	grpc "best-goph-keeper/internal/server/proto"
	"context"
	"github.com/sirupsen/logrus"
)

type Event struct {
	grpc    grpc.GophkeeperClient
	config  *config.ConfigClient
	logger  *logrus.Logger
	context context.Context
	grpc.UnimplementedGophkeeperServer
}

// NewEvent - creates a new grpc client instance
func NewEvent(ctx context.Context, config *config.ConfigClient, log *logrus.Logger, client grpc.GophkeeperClient) *Event {
	return &Event{context: ctx, config: config, logger: log, grpc: client}
}

// GetConfig - returns config in string format
func (s Event) GetConfig() *config.ConfigClient {
	return s.config
}
