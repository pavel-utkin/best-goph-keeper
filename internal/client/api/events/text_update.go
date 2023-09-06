package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"best-goph-keeper/internal/server/storage/vars"
	"context"
)

func (c Event) TextUpdate(name, passwordSecure, text string, token model.Token) error {
	c.logger.Info("text update")

	secretKey := encryption.AesKeySecureRandom([]byte(passwordSecure))
	encryptText, err := encryption.Encrypt(text, secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)

	updatedTextEntityID, err := c.grpc.EntityUpdate(context.Background(),
		&grpc.UpdateEntityRequest{Name: name, Data: []byte(encryptText), Type: vars.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(updatedTextEntityID)
	return nil
}
