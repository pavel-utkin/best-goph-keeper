package grpchandler

import (
	grpc "best-goph-keeper/internal/server/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Ping - checks the database connection
func (h *Handler) Ping(ctx context.Context, req *grpc.PingRequest) (*grpc.PingResponse, error) {
	h.logger.Info("ping")
	var msg string
	err := h.database.Ping()
	if err != nil {
		msg = "unsuccessful database connection"
		h.logger.Error(err)
		return &grpc.PingResponse{Message: msg}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	msg = "successful database connection"
	h.logger.Info(msg)
	return &grpc.PingResponse{Message: msg}, nil
}
