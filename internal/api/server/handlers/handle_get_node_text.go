package handlers

import (
	"best-goph-keeper/internal/model"
	"context"

	grpc "best-goph-keeper/internal/api/proto"
)

// HandleGetNodeText - get node text
func (h *Handler) HandleGetNodeText(ctx context.Context, req *grpc.GetNodeTextRequest) (*grpc.GetNodeTextResponse, error) {
	TextData := &model.GetNodeTextRequest{}
	TextData.TextId = req.TextId
	GetNodeText, err := h.text.GetNodeText(TextData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetNodeTextResponse{}, err
	}
	h.logger.Debug(GetNodeText)
	return &grpc.GetNodeTextResponse{TextId: GetNodeText.ID, Text: GetNodeText.Text}, nil
}
