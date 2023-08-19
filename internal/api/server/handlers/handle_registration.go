package handlers

import (
	grpc "best-goph-keeper/internal/api/proto"
	"best-goph-keeper/internal/model"
	"context"
)

// HandleRegistration - registration user
func (h *Handler) HandleRegistration(ctx context.Context, req *grpc.RegistrationRequest) (*grpc.RegistrationResponse, error) {
	var resp string
	RegistrationData := &model.RegistrationRequest{}

	RegistrationData.Username = req.Username
	RegistrationData.Password = req.Password

	_, err := h.user.Registration(RegistrationData)
	if err != nil {
		resp = "unsuccessful registration user"
		h.logger.Error(err)
		return &grpc.RegistrationResponse{Resp: resp}, err
	}
	resp = "successful registration user"
	h.logger.Info(resp)

	return &grpc.RegistrationResponse{Resp: resp}, err
}
