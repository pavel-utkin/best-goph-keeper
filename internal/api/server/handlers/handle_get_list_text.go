package handlers

import (
	"best-goph-keeper/internal/model"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc "best-goph-keeper/internal/api/proto"
)

// HandleGetListText - get list text
func (h *Handler) HandleGetListText(ctx context.Context, req *grpc.GetListTextRequest) (*grpc.GetListTextResponse, error) {
	h.logger.Info("Get list text")

	valid, accessToken, err := h.token.Validate(req.AccessToken)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListTextResponse{}, status.Errorf(
			codes.Unauthenticated, err.Error(),
		)
	}
	if !valid {
		h.logger.Error("Not validate token")
		return &grpc.GetListTextResponse{}, status.Errorf(
			codes.Unauthenticated, err.Error(),
		)
	}

	ListText, err := h.text.GetListText(accessToken.UserID)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListTextResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list := model.GetListData(ListText)

	h.logger.Debug(ListText)
	return &grpc.GetListTextResponse{Node: list}, nil
}
