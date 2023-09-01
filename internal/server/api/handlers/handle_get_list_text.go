package handlers

import (
	"best-goph-keeper/internal/server/model"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc "best-goph-keeper/internal/server/proto"
)

// HandleGetListText - get list text
func (h *Handler) HandleGetListText(ctx context.Context, req *grpc.GetListTextRequest) (*grpc.GetListTextResponse, error) {
	h.logger.Info("Get list text")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetListTextResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	ListText, err := h.text.GetListText(req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list := model.GetListText(ListText)

	h.logger.Debug(ListText)
	return &grpc.GetListTextResponse{Node: list}, nil
}
