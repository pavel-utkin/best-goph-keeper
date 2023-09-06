package events

import (
	"best-goph-keeper/internal/client/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
)

func (c Event) FileRemove(binary []string, token model.Token) error {
	c.logger.Info("file remove")

	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)
	deletedCard, err := c.grpc.FileRemove(context.Background(),
		&grpc.DeleteBinaryRequest{Name: binary[0], AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
			CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(deletedCard)
	return nil
}
