package grpchandler

import (
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EntityDelete - check the validity of the token and delete record (text, bank card or login password)
func (h *Handler) EntityDelete(ctx context.Context, req *grpc.DeleteEntityRequest) (*grpc.DeleteEntityResponse, error) {
	h.logger.Info("delete entity")

	endDateToken, err := h.token.GetEndDateToken(req.AccessToken.Token)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteEntityResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	DeletedEntityID, err := h.entity.Delete(req.AccessToken.UserId, req.Name, req.Type)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(DeletedEntityID)
	return &grpc.DeleteEntityResponse{Id: DeletedEntityID}, nil
}
