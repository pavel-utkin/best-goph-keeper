package handlers

import (
	grpc "best-goph-keeper/internal/server/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleUserExist - check user exist
func (h *Handler) HandleUserExist(ctx context.Context, req *grpc.UserExistRequest) (*grpc.UserExistResponse, error) {
	h.logger.Info("User Exist")
	exist, err := h.user.UserExists(req.Username)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UserExistResponse{Exist: false}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	h.logger.Debug(exist)
	return &grpc.UserExistResponse{Exist: exist}, nil
}
