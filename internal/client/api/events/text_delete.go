package events

import (
	"best-goph-keeper/internal/client/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"best-goph-keeper/internal/server/storage/vars"
	"context"
)

func (c Event) TextDelete(text []string, token model.Token) error {
	c.logger.Info("text delete")

	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)

	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)

	deletedTextEntityID, err := c.grpc.EntityDelete(context.Background(),
		&grpc.DeleteEntityRequest{Name: text[0], Type: vars.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(deletedTextEntityID)
	return nil
}
