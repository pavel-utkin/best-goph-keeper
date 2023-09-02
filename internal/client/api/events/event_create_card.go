package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/storage/layouts"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
	"encoding/json"
	"strconv"
	"time"
)

func (c Event) EventCreateCard(name, description, password, paymentSystem, number, holder, cvc, endDate string, token model.Token) error {
	c.logger.Info("Create card")

	intCvc, err := strconv.Atoi(cvc)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	timeEndDate, err := time.Parse(layouts.LayoutDate.ToString(), endDate)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	card := model.Card{Name: name, Description: description, PaymentSystem: paymentSystem, Number: number, Holder: holder, EndDate: timeEndDate, CVC: intCvc}
	jsonCard, err := json.Marshal(card)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptCard, err := encryption.Encrypt(string(jsonCard), secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	createdCard, err := c.grpc.HandleCreateCard(context.Background(),
		&grpc.CreateCardRequest{Name: name, Description: description, Data: []byte(encryptCard),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}
	c.logger.Debug(createdCard)
	return nil
}
