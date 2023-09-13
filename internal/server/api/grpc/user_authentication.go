package grpchandler

import (
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authentication - user authentication, create access token
func (h *Handler) Authentication(ctx context.Context, req *grpc.AuthenticationRequest) (*grpc.AuthenticationResponse, error) {
	h.logger.Info("authentication")
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

	token, err := h.token.Create(user.UserId, h.config.AccessTokenLifetime)
	if err != nil {
		h.logger.Error(err)
		return &grpc.AuthenticationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)

	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)

	h.logger.Debug(authenticatedUser)
	return &grpc.AuthenticationResponse{AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}}, nil
}
