package handlers

import (
	"context"

	grpc "best-goph-keeper/internal/api/proto"
)

// HandleGetListText - get list text
func (h *Handler) HandleGetListText(ctx context.Context, req *grpc.GetListTextRequest) (*grpc.GetListTextResponse, error) {
	var resp string
	return &grpc.GetListTextResponse{Resp: resp}, nil
}