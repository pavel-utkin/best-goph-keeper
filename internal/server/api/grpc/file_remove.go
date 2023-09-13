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

// FileRemove - checks the validity of the token, delete record, remove file on server
func (h *Handler) FileRemove(ctx context.Context, req *grpc.DeleteBinaryRequest) (*grpc.DeleteBinaryResponse, error) {
	h.logger.Info("file remove")

	endDateToken, err := h.token.GetEndDateToken(req.AccessToken.Token)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteBinaryResponse{}, err
	}
	valid := h.token.Validate(endDateToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	FileData := &model.FileRequest{}
	FileData.UserID = req.AccessToken.UserId
	FileData.Name = req.Name

	BinaryId, err := h.file.DeleteFile(FileData)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	err = service.RemoveFile(h.config.FileFolder, req.AccessToken.UserId, req.Name)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteBinaryResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	return &grpc.DeleteBinaryResponse{Id: BinaryId}, nil
}
