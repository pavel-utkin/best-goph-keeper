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

// HandleDeleteCard - delete card
func (h *Handler) HandleDeleteCard(ctx context.Context, req *grpc.DeleteCardRequest) (*grpc.DeleteCardResponse, error) {
	h.logger.Info("Delete card")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteCardResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	cardID, err := h.card.GetIdCard(req.Name, req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	metadataRequest := model.DeleteMetadataRequest{cardID, string(vars.Name), req.Name, string(vars.Card)}
	err = h.metadata.DeleteMetadata(metadataRequest)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = h.card.DeleteCard(cardID)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteCardResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DeleteCardResponse{Id: cardID}, nil
}
