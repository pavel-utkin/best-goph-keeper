package handlers

import (
	"best-goph-keeper/internal/model"
	"best-goph-keeper/internal/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc "best-goph-keeper/internal/api/proto"
)

// HandleCreateText - create text
func (h *Handler) HandleCreateText(ctx context.Context, req *grpc.CreateTextRequest) (*grpc.CreateTextResponse, error) {
	h.logger.Info("Create text")

	valid, accessToken, err := h.token.Validate(req.AccessToken)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Unauthenticated, err.Error(),
		)
	}
	if !valid {
		h.logger.Error("Not validate token")
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Unauthenticated, err.Error(),
		)
	}

	TextData := &model.CreateTextRequest{}
	TextData.UserID = accessToken.UserID
	TextData.Key = req.Key
	TextData.Value = req.Value
	TextData.Text = req.Text

	if TextData.Key == "" || TextData.Value == "" {
		err := errors.ErrNoMetadataSet
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}
	exists, err := h.text.KeyExists(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists == true {
		err = errors.ErrKeyAlreadyExists
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	CreatedText, err := h.text.CreateText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	text := model.GetTextData(CreatedText)

	Metadata := &model.CreateMetadataRequest{}
	Metadata.EntityId = CreatedText.ID
	Metadata.Key = TextData.Key
	Metadata.Value = TextData.Value
	Metadata.Type = TextData.Type
	CreatedMetadata, err := h.metadata.CreateMetadata(Metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(CreatedText)
	h.logger.Debug(CreatedMetadata)
	return &grpc.CreateTextResponse{Text: text}, nil
}
