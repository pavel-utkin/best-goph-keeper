package handlers

import (
	"context"

	grpc "best-goph-keeper/internal/api/proto"
)

// HandleCreateText - create text
func (h *Handler) HandleCreateText(ctx context.Context, req *grpc.CreateTextRequest) (*grpc.CreateTextResponse, error) {
	var resp string
	return &grpc.CreateTextResponse{Resp: resp}, nil
}
