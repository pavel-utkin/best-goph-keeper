package events

import (
	"best-goph-keeper/internal/client/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"best-goph-keeper/internal/server/storage/vars"
	"context"
)

// CardDelete -  delete card
func (c Event) CardDelete(card []string, token model.Token) error {
	c.logger.Info("card delete")

	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)

	deletedCardEntityID, err := c.grpc.EntityDelete(context.Background(),
		&grpc.DeleteEntityRequest{Name: card[0], Type: vars.Card.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(deletedCardEntityID)
	return nil
}
