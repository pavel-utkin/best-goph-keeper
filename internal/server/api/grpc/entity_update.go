package grpchandler

import (
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EntityUpdate - checks the validity of the token and update record (text, bank card or login password)
func (h *Handler) EntityUpdate(ctx context.Context, req *grpc.UpdateEntityRequest) (*grpc.UpdateEntityResponse, error) {
	h.logger.Info("entity update")

	endDateToken, err := h.token.GetEndDateToken(req.AccessToken.Token)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UpdateEntityResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.UpdateEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	UpdatedEntityID, err := h.entity.Update(req.AccessToken.UserId, req.Name, req.Type, req.Data)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UpdateEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(UpdatedEntityID)
	return &grpc.UpdateEntityResponse{Id: UpdatedEntityID}, nil
}
