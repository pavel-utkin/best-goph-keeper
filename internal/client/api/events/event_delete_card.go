package events

import (
	"best-goph-keeper/internal/client/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
)

func (c Event) EventDeleteCard(card []string, token model.Token) error {
	c.logger.Info("Delete card")

	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	deletedCard, err := c.grpc.HandleDeleteCard(context.Background(),
		&grpc.DeleteCardRequest{Name: card[0], AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(card)
	c.logger.Debug(deletedCard)
	return nil
}
