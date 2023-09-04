package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
)

func (c Event) FileUpload(name string, password string, file []byte, token model.Token) (string, error) {
	c.logger.Info("file upload")

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptFile, err := encryption.Encrypt(string(file), secretKey)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	createdToken, err := service.ConvertTimeToTimestamp(token.CreatedAt)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	endDateToken, err := service.ConvertTimeToTimestamp(token.EndDateAt)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	uploadFile, err := c.grpc.FileUpload(context.Background(),
		&grpc.UploadBinaryRequest{Name: name, Data: []byte(encryptFile),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	c.logger.Debug(uploadFile.Name)
	return uploadFile.Name, nil
}
