package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
)

func (c Event) Authentication(username, password string) (model.Token, error) {
	c.logger.Info("authentication")

	token := model.Token{}
	password, err := encryption.HashPassword(password)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	authenticatedUser, err := c.grpc.Authentication(c.context, &grpc.AuthenticationRequest{Username: username, Password: password})
	if err != nil {
		c.logger.Error(err)
		return token, err
	}

	createdToken, err := service.ConvertTimestampToTime(authenticatedUser.AccessToken.CreatedAt)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	endDateToken, err := service.ConvertTimestampToTime(authenticatedUser.AccessToken.EndDateAt)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	token = model.Token{AccessToken: authenticatedUser.AccessToken.Token, UserID: authenticatedUser.AccessToken.UserId,
		CreatedAt: createdToken, EndDateAt: endDateToken}

	err = service.CreateStorageNotExistsUser(c.config.FileFolder, token.UserID)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}

	return token, nil
}
