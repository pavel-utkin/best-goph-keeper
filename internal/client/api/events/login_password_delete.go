package events

import (
	"best-goph-keeper/internal/client/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"best-goph-keeper/internal/server/storage/vars"
	"context"
)

// LoginPasswordDelete - delete login-password
func (c Event) LoginPasswordDelete(loginPassword []string, token model.Token) error {
	c.logger.Info("login password delete")

	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)

	deletedLoginPasswordEntityID, err := c.grpc.EntityDelete(context.Background(),
		&grpc.DeleteEntityRequest{Name: loginPassword[0], Type: vars.LoginPassword.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(deletedLoginPasswordEntityID)
	return nil
}
