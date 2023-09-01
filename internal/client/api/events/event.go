package events

import (
	grpc "best-goph-keeper/internal/server/proto"
	"context"
	"github.com/sirupsen/logrus"
)

type Event struct {
	grpc    grpc.GophkeeperClient
	logger  *logrus.Logger
	context context.Context
	grpc.UnimplementedGophkeeperServer
}

// NewEvent - creates a new grpc client instance
func NewEvent(ctx context.Context, log *logrus.Logger, client grpc.GophkeeperClient) *Event {
	return &Event{context: ctx, logger: log, grpc: client}
}
