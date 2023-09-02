package handlers

import (
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleUpdateText - Update text
func (h *Handler) HandleUpdateText(ctx context.Context, req *grpc.UpdateTextRequest) (*grpc.UpdateTextResponse, error) {
	h.logger.Info("Update text")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.UpdateTextResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	textID, err := h.text.GetIdText(req.Name, req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UpdateTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = h.text.UpdateText(textID, req.Data)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UpdateTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.UpdateTextResponse{Id: textID}, nil
}
