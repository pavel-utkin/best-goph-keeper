package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
)

func (c Event) EventUpdateText(name, passwordSecure, text string, token model.Token) error {
	c.logger.Info("Update text")

	secretKey := encryption.AesKeySecureRandom([]byte(passwordSecure))
	encryptText, err := encryption.Encrypt(text, secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	updateText, err := c.grpc.HandleUpdateText(context.Background(), &grpc.UpdateTextRequest{Name: name, Data: []byte(encryptText),
		AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(updateText)
	return nil
}
