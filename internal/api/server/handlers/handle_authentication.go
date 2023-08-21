package handlers

import (
	grpc "best-goph-keeper/internal/api/proto"
	"best-goph-keeper/internal/model"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleAuthentication - authentication user
func (h *Handler) HandleAuthentication(ctx context.Context, req *grpc.AuthenticationRequest) (*grpc.AuthenticationResponse, error) {
	h.logger.Info("Authentication")
	UserData := &model.UserRequest{
		Username: req.Username,
		Password: req.Password,
	}

	authenticatedUser, err := h.user.Authentication(UserData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.AuthenticationResponse{}, status.Errorf(
			codes.Unauthenticated, err.Error(),
		)
	}
	user := model.GetUserData(authenticatedUser)

	token, err := h.token.Create(user.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.AuthenticationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(authenticatedUser)
	return &grpc.AuthenticationResponse{User: user, AccessToken: token}, nil
}
