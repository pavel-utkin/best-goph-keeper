package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/storage/layouts"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"best-goph-keeper/internal/server/storage/vars"
	"context"
	"encoding/json"
	"strconv"
	"time"
)

// CardUpdate - update card
func (c Event) CardUpdate(name, passwordSecure, paymentSystem, number, holder, cvc, endDateCard string, token model.Token) error {
	c.logger.Info("card update")

	intCvc, err := strconv.Atoi(cvc)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	timeEndDate, err := time.Parse(layouts.LayoutDate.ToString(), endDateCard)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	card := model.Card{PaymentSystem: paymentSystem, Number: number, Holder: holder, CVC: intCvc, EndDate: timeEndDate}
	jsonCard, err := json.Marshal(card)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	secretKey := encryption.AesKeySecureRandom([]byte(passwordSecure))
	encryptCard, err := encryption.Encrypt(string(jsonCard), secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	createdToken := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken := service.ConvertTimeToTimestamp(token.EndDateAt)

	updatedCardEntityID, err := c.grpc.EntityUpdate(context.Background(),
		&grpc.UpdateEntityRequest{Name: name, Data: []byte(encryptCard), Type: vars.Card.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(updatedCardEntityID)
	return nil
}
