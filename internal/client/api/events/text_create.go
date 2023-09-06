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

func (c Event) TextCreate(name, description, password, plaintext string, token model.Token) error {
	c.logger.Info("text create")

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptText, err := encryption.Encrypt(plaintext, secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)

	metadata := model.MetadataEntity{Name: name, Description: description, Type: vars.Text.ToString()}
	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	createdEntityID, err := c.grpc.EntityCreate(context.Background(),
		&grpc.CreateEntityRequest{Data: []byte(encryptText), Metadata: string(jsonMetadata),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(createdEntityID)
	return nil
}
