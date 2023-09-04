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

// FileDownload - checks the validity of the token, save record, upload file on client
func (h *Handler) FileDownload(ctx context.Context, req *grpc.DownloadBinaryRequest) (*grpc.DownloadBinaryResponse, error) {
	h.logger.Info("file download")

	endDateToken, err := h.token.GetEndDateToken(req.AccessToken.Token)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DownloadBinaryResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	FileData := &model.FileRequest{}
	FileData.UserID = req.AccessToken.UserId
	FileData.Name = req.Name

	exists, err := h.file.FileExists(FileData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}
	if exists != true {
		err = errors.ErrFileNotExists
		h.logger.Error(err)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.AlreadyExists, err.Error(),
		)
	}

	data, err := service.DownloadFile(h.config.FileFolder, req.AccessToken.UserId, req.Name)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DownloadBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DownloadBinaryResponse{Data: data}, nil
}
