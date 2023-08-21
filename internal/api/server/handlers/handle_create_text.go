package handlers

import (
	"best-goph-keeper/internal/model"
	"best-goph-keeper/internal/service/validator"
	"best-goph-keeper/internal/storage/errors"
	"context"

	grpc "best-goph-keeper/internal/api/proto"
)

// HandleCreateText - create text
func (h *Handler) HandleCreateText(ctx context.Context, req *grpc.CreateTextRequest) (*grpc.CreateTextResponse, error) {
	if correctText := validator.VerifyText(req.Text); correctText != true {
		err := errors.ErrBadText
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}

	TextData := &model.CreateTextRequest{}
	TextData.UserID = req.UserId
	TextData.Text = req.Text
	CreatedText, err := h.text.CreateText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, err
	}
	h.logger.Debug(CreatedText)

	return &grpc.CreateTextResponse{TextId: CreatedText.ID, Text: CreatedText.Text}, nil
}
