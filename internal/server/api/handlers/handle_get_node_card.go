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

// HandleGetNodeCard - get node card
func (h *Handler) HandleGetNodeCard(ctx context.Context, req *grpc.GetNodeCardRequest) (*grpc.GetNodeCardResponse, error) {
	h.logger.Info("Get node card")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetNodeCardResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	CardData := &model.GetNodeCardRequest{}
	CardData.UserID = req.AccessToken.UserId
	CardData.Key = string(vars.Name)
	CardData.Value = req.Name
	GetNodeCard, err := h.card.GetNodeCard(CardData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetNodeCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	card := model.GetCardData(GetNodeCard)

	h.logger.Debug(GetNodeCard)
	return &grpc.GetNodeCardResponse{Data: card}, nil
}
