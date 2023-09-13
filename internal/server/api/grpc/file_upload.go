package grpchandler

import (
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"best-goph-keeper/internal/server/storage/errors"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FileUpload - checks the validity of the token, upload file on client
func (h *Handler) FileUpload(ctx context.Context, req *grpc.UploadBinaryRequest) (*grpc.UploadBinaryResponse, error) {
	h.logger.Info("file upload")

	endDateToken, err := h.token.GetEndDateToken(req.AccessToken.Token)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	FileData := &model.FileRequest{}
	FileData.UserID = req.AccessToken.UserId
	FileData.Name = req.Name

	exists, err := h.file.FileExists(FileData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists {
		err = errors.ErrNameAlreadyExists
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	UploadFile, err := h.file.UploadFile(FileData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = service.UploadFile(h.config.FileFolder, req.AccessToken.UserId, req.Name, req.Data)
	if err != nil {
		h.logger.Error(err)
		return &grpc.UploadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(UploadFile.Name)
	return &grpc.UploadBinaryResponse{Name: UploadFile.Name}, nil
}
