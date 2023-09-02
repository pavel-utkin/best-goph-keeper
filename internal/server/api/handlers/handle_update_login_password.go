package handlers

import (
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleUpdateLoginPassword - Update login password
func (h *Handler) HandleUpdateLoginPassword(ctx context.Context, req *grpc.UpdateLoginPasswordRequest) (*grpc.UpdateLoginPasswordResponse, error) {
	h.logger.Info("Update login password")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.UpdateLoginPasswordResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	loginPasswordID, err := h.loginPassword.GetIdLoginPassword(req.Name, req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UpdateLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = h.loginPassword.UpdateLoginPassword(loginPasswordID, req.Data)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UpdateLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.UpdateLoginPasswordResponse{Id: loginPasswordID}, nil
}
