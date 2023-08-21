package handlers

import (
	grpc "best-goph-keeper/internal/api/proto"
	"best-goph-keeper/internal/model"
	"best-goph-keeper/internal/service/validator"
	"best-goph-keeper/internal/storage/errors"
	"context"
)

// HandleRegistration - registration new user
func (h *Handler) HandleRegistration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	if correctPassword := validator.VerifyPassword(req.Password); correctPassword != true {
		err := errors.ErrBadPassword
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, err
	}

	UserData := &model.UserRequest{}
	UserData.Username = req.Username
	UserData.Password = req.Password

	exists, err := h.user.UserExists(UserData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, err
	}
	if exists == true {
		err = errors.ErrUserAlreadyExists
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, err
	}
	registeredUser, err := h.user.Registration(UserData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.RegistrationResponse{}, err
	}
	h.logger.Debug(registeredUser)
	return &grpc.RegistrationResponse{UserId: registeredUser.ID, Username: registeredUser.Username}, nil
}
