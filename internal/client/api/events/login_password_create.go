package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"best-goph-keeper/internal/server/storage/vars"
	"context"
	"encoding/json"
)

// LoginPasswordCreate - add login-password
func (c Event) LoginPasswordCreate(name, description, passwordSecure, login, password string, token model.Token) error {
	c.logger.Info("login password create")

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
	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)

	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)

	metadata := model.MetadataEntity{Name: name, Description: description, Type: vars.LoginPassword.ToString()}
	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	createdEntityID, err := c.grpc.EntityCreate(context.Background(),
		&grpc.CreateEntityRequest{Data: []byte(encryptLoginPassword), Metadata: string(jsonMetadata),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(createdEntityID)
	return nil
}
