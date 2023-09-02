package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
	"encoding/json"
)

func (c Event) EventUpdateLoginPassword(name, passwordSecure, login, password string, token model.Token) error {
	c.logger.Info("Update login password")

	loginPassword := model.LoginPassword{Login: login, Password: password}
	jsonLoginPassword, err := json.Marshal(loginPassword)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	secretKey := encryption.AesKeySecureRandom([]byte(passwordSecure))
	encryptLoginPassword, err := encryption.Encrypt(string(jsonLoginPassword), secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	updateLoginPassword, err := c.grpc.HandleUpdateLoginPassword(context.Background(), &grpc.UpdateLoginPasswordRequest{Name: name, Data: []byte(encryptLoginPassword),
		AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(updateLoginPassword)
	return nil
}
