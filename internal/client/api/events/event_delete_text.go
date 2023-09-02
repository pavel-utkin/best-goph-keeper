package events

import (
	"best-goph-keeper/internal/client/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
)

func (c Event) EventDeleteText(text []string, token model.Token) error {
	c.logger.Info("Delete text")

	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	deletedText, err := c.grpc.HandleDeleteText(context.Background(),
		&grpc.DeleteTextRequest{Name: text[0], AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(text)
	c.logger.Debug(deletedText)
	return nil
}
