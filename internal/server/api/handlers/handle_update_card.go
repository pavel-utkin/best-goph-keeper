package handlers

import (
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleUpdateCard - update card
func (h *Handler) HandleUpdateCard(ctx context.Context, req *grpc.UpdateCardRequest) (*grpc.UpdateCardResponse, error) {
	h.logger.Info("Update card")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.UpdateCardResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	cardID, err := h.card.GetIdCard(req.Name, req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UpdateCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = h.card.UpdateCard(cardID, req.Data)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UpdateCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.UpdateCardResponse{Id: cardID}, nil
}
