package handlers

import (
	"best-goph-keeper/internal/server/model"
	"best-goph-keeper/internal/server/storage/errors"
	"best-goph-keeper/internal/server/storage/vars"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc "best-goph-keeper/internal/server/proto"
)

// HandleGetNodeText - get node text
func (h *Handler) HandleGetNodeText(ctx context.Context, req *grpc.GetNodeTextRequest) (*grpc.GetNodeTextResponse, error) {
	h.logger.Info("Get node text")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetNodeTextResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	TextData := &model.GetNodeTextRequest{}
	TextData.UserID = req.AccessToken.UserId
	TextData.Key = string(vars.Name)
	TextData.Value = req.Name
	GetNodeText, err := h.text.GetNodeText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetNodeTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	text := model.GetTextData(GetNodeText)

	h.logger.Debug(GetNodeText)
	return &grpc.GetNodeTextResponse{Text: text}, nil
}
