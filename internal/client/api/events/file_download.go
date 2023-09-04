package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
)

func (c Event) FileDownload(name string, password string, token model.Token) error {
	c.logger.Info("file download")

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	createdToken, err := service.ConvertTimeToTimestamp(token.CreatedAt)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	endDateToken, err := service.ConvertTimeToTimestamp(token.EndDateAt)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	downloadFile, err := c.grpc.FileDownload(context.Background(),
		&grpc.DownloadBinaryRequest{Name: name, AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
			CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	file, err := encryption.Decrypt(string(downloadFile.Data), secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	err = service.UploadFile(c.config.FileFolder, token.UserID, name, []byte(file))
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(name)
	return nil
}
