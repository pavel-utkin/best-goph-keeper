package grpchandler

import (
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EntityGetList - checks the validity of token and get list records (text, bank card or login-password)
func (h *Handler) EntityGetList(ctx context.Context, req *grpc.GetListEntityRequest) (*grpc.GetListEntityResponse, error) {
	h.logger.Info("Get list entity")

	endDateToken, err := h.token.GetEndDateToken(req.AccessToken.Token)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListEntityResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.GetListEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	ListEntity, err := h.entity.GetList(req.AccessToken.UserId, req.Type)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	list, err := model.GetListEntity(ListEntity)
	if err != nil {
		h.logger.Error(err)
		return &grpc.GetListEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(list)
	return &grpc.GetListEntityResponse{Node: list}, nil
}
