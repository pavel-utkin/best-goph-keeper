package handlers

import (
	grpc "best-goph-keeper/internal/api/proto"
	"best-goph-keeper/internal/model"
	"best-goph-keeper/internal/service/validator"
	"best-goph-keeper/internal/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleRegistration - registration new user
func (h *Handler) HandleRegistration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	h.logger.Info("Registration")
	if correctPassword := validator.VerifyPassword(req.Password); correctPassword != true {
		err := errors.ErrBadPassword
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}
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
	if exists == true {
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

	token, err := h.token.Create(user.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(registeredUser)
	return &grpc.RegistrationResponse{User: user, AccessToken: token}, nil
}
