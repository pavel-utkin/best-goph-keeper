package grpchandler

import (
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EntityCreate - check the validity of the token and save record (text, bank card or login password)
func (h *Handler) EntityCreate(ctx context.Context, req *grpc.CreateEntityRequest) (*grpc.CreateEntityResponse, error) {
	h.logger.Info("entity create")

	endDateToken, err := h.token.GetEndDateToken(req.AccessToken.Token)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	var metadata model.MetadataEntity
	err = json.Unmarshal([]byte(req.Metadata), &metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, err
	}

	EntityData := &model.CreateEntityRequest{}
	EntityData.UserID = req.AccessToken.UserId
	EntityData.Data = req.Data
	EntityData.Metadata = metadata
	if metadata.Name == "" {
		err := errors.ErrNoMetadataSet
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}

	exists, err := h.entity.Exists(EntityData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists == true {
		err = errors.ErrNameAlreadyExists
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	CreatedEntityID, err := h.entity.Create(EntityData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(CreatedEntityID)
	return &grpc.CreateEntityResponse{Id: CreatedEntityID}, nil
}
