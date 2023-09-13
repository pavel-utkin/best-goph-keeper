package grpchandler

import (
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FileGetList - checks the validity of tokens and get list records file
func (h *Handler) FileGetList(ctx context.Context, req *grpc.GetListBinaryRequest) (*grpc.GetListBinaryResponse, error) {
	h.logger.Info("file get list")

	endDateToken, err := h.token.GetEndDateToken(req.AccessToken.Token)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListBinaryResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetListBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	ListFile, err := h.file.GetListFile(req.AccessToken.UserId)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list := model.GetListFile(ListFile)

	h.logger.Debug(ListFile)
	return &grpc.GetListBinaryResponse{Node: list}, nil
}
