package handlers

import (
	grpc "best-goph-keeper/internal/api/proto"
	"best-goph-keeper/internal/model"
	"best-goph-keeper/internal/service/validator"
	"context"
)

// HandleRegistration - registration user
func (h *Handler) HandleRegistration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	var resp string
	UserData := &model.UserRequest{}
	UserData.Username = req.Username
	UserData.Password = req.Password

	if correctPassword := validator.VerifyPassword(req.Password); correctPassword != true {
		resp = "password rules: at least 7 letters, 1 number, 1 upper case, 1 special character"
		h.logger.Error(resp)
		return &grpc.RegistrationResponse{Resp: resp}, nil
	}

	exists, err := h.user.UserExists(UserData)
	if err != nil {
		resp = "server error"
		h.logger.Error(err)
		return &grpc.RegistrationResponse{Resp: resp}, err
	}
	if exists == true {
		resp = "user with this name already exists"
		h.logger.Error(err)
		return &grpc.RegistrationResponse{Resp: resp}, err
	}

	_, err = h.user.Registration(UserData)
	if err != nil {
		resp = "unsuccessful registration user"
		h.logger.Error(err)
		return &grpc.RegistrationResponse{Resp: resp}, err
	}
	resp = "successful registration user"
	h.logger.Info(resp)

	return &grpc.RegistrationResponse{Resp: resp}, nil
}
