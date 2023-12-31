package grpchandler

import (
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Registration - registration new user, create access token
func (h *Handler) Registration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	h.logger.Info("registration")

	UserData := &model.UserRequest{}
	UserData.Username = req.Username
	UserData.Password = req.Password

	exists, err := h.user.UserExists(UserData.Username)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists {
		err = errors.ErrUsernameAlreadyExists
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}
	registeredUser, err := h.user.Registration(UserData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	user := model.GetUserData(registeredUser)

	token, err := h.token.Create(user.UserId, h.config.AccessTokenLifetime)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)

	err = service.CreateStorageUser(h.config.FileFolder, token.UserID)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(registeredUser)
	return &grpc.RegistrationResponse{AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
		CreatedAt: createdToken, EndDateAt: endDateToken}}, nil
}
