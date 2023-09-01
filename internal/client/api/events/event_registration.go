package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
)

func (c Event) EventRegistration(username, password string) (model.Token, error) {
	c.logger.Info("Registration")
	token := model.Token{}
	password, err := encryption.HashPassword(password)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	registeredUser, err := c.grpc.HandleRegistration(c.context, &grpc.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	created, _ := service.ConvertTimestampToTime(registeredUser.AccessToken.CreatedAt)
	endDate, _ := service.ConvertTimestampToTime(registeredUser.AccessToken.EndDateAt)
	token = model.Token{AccessToken: registeredUser.AccessToken.Token, UserID: registeredUser.AccessToken.UserId,
		CreatedAt: created, EndDateAt: endDate}
	return token, nil
}
