package handlers

import (
	"best-goph-keeper/internal/model"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc "best-goph-keeper/internal/api/proto"
)

// HandleGetNodeText - get node text
func (h *Handler) HandleGetNodeText(ctx context.Context, req *grpc.GetNodeTextRequest) (*grpc.GetNodeTextResponse, error) {
	h.logger.Info("Get node text")

	valid, accessToken, err := h.token.Validate(req.AccessToken)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetNodeTextResponse{}, status.Errorf(
			codes.Unauthenticated, err.Error(),
		)
	}
	if !valid {
		h.logger.Error("Not validate token")
		return &grpc.GetNodeTextResponse{}, status.Errorf(
			codes.Unauthenticated, err.Error(),
		)
	}

	TextData := &model.GetNodeTextRequest{}
	TextData.UserID = accessToken.UserID
	TextData.Key = req.Key
	TextData.Value = req.Value
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
