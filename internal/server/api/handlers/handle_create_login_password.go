package handlers

import (
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage/errors"
	"best-goph-keeper/internal/server/storage/vars"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleCreateLoginPassword - create login password
func (h *Handler) HandleCreateLoginPassword(ctx context.Context, req *grpc.CreateLoginPasswordRequest) (*grpc.CreateLoginPasswordResponse, error) {
	h.logger.Info("Create login password")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.CreateLoginPasswordResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	LoginPasswordData := &model.CreateLoginPasswordRequest{}
	LoginPasswordData.UserID = req.AccessToken.UserId
	LoginPasswordData.Name = req.Name
	LoginPasswordData.Data = req.Data
	LoginPasswordData.Description = req.Description
	if LoginPasswordData.Name == "" {
		err := errors.ErrNoMetadataSet
		h.logger.Error(err)
		return &grpc.CreateLoginPasswordResponse{}, status.Errorf(
			codes.InvalidArgument, err.Error(),
		)
	}
	exists, err := h.loginPassword.KeyExists(LoginPasswordData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists == true {
		err = errors.ErrKeyAlreadyExists
		h.logger.Error(err)
		return &grpc.CreateLoginPasswordResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	CreatedLoginPassword, err := h.loginPassword.CreateLoginPassword(LoginPasswordData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	loginPassword := model.GetLoginPassword(CreatedLoginPassword)

	Metadata := &model.CreateMetadataRequest{}
	Metadata.EntityId = CreatedLoginPassword.ID
	Metadata.Key = string(vars.Name)
	Metadata.Value = LoginPasswordData.Name
	Metadata.Type = string(vars.LoginPassword)
	CreatedMetadataName, err := h.metadata.CreateMetadata(Metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	Metadata = &model.CreateMetadataRequest{}
	Metadata.EntityId = CreatedLoginPassword.ID
	Metadata.Key = string(vars.Description)
	Metadata.Value = LoginPasswordData.Description
	Metadata.Type = string(vars.LoginPassword)
	CreatedMetadataDescription, err := h.metadata.CreateMetadata(Metadata)
	if err != nil {
		h.logger.Error(err)
		return &grpc.CreateLoginPasswordResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(CreatedLoginPassword)
	h.logger.Debug(CreatedMetadataName)
	h.logger.Debug(CreatedMetadataDescription)
	return &grpc.CreateLoginPasswordResponse{Data: loginPassword}, nil
}
