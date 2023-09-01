package handlers

import (
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleGetListText - get list card
func (h *Handler) HandleGetListCard(ctx context.Context, req *grpc.GetListCardRequest) (*grpc.GetListCardResponse, error) {
	h.logger.Info("Get list card")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetListCardResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	ListCard, err := h.card.GetListCard(req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list := model.GetListCard(ListCard)

	h.logger.Debug(ListCard)
	return &grpc.GetListCardResponse{Node: list}, nil
}
