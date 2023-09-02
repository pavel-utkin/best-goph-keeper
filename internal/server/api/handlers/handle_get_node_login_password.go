package handlers

import (
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"best-goph-keeper/internal/server/storage/vars"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGetNodeLoginPassword - get node login password
func (h *Handler) HandleGetNodeLoginPassword(ctx context.Context, req *grpc.GetNodeLoginPasswordRequest) (*grpc.GetNodeLoginPasswordResponse, error) {
	h.logger.Info("Get node login password")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetNodeLoginPasswordResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	LoginPasswordData := &model.GetNodeLoginPasswordRequest{}
	LoginPasswordData.UserID = req.AccessToken.UserId
	LoginPasswordData.Key = string(vars.Name)
	LoginPasswordData.Value = req.Name
	GetNodeLoginPassword, err := h.loginPassword.GetNodeLoginPassword(LoginPasswordData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetNodeLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	loginPassword := model.GetLoginPassword(GetNodeLoginPassword)

	h.logger.Debug(GetNodeLoginPassword)
	return &grpc.GetNodeLoginPasswordResponse{Data: loginPassword}, nil
}
