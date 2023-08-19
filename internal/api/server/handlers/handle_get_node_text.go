package handlers

import (
	"context"

	grpc "best-goph-keeper/internal/api/proto"
)

// HandleGetNodeText - get node text
func (h *Handler) HandleGetNodeText(ctx context.Context, req *grpc.GetNodeTextRequest) (*grpc.GetNodeTextResponse, error) {
	var resp string
	return &grpc.GetNodeTextResponse{Resp: resp}, nil
}
