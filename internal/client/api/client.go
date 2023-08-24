package api

import (
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/service/table"
	"best-goph-keeper/internal/server/model"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Client struct {
	grpc    grpc.GophkeeperClient
	logger  *logrus.Logger
	context context.Context
	grpc.UnimplementedGophkeeperServer
}

// NewClient - creates a new grpc client instance
func NewClient(ctx context.Context, log *logrus.Logger, client grpc.GophkeeperClient) *Client {
	return &Client{context: ctx, logger: log, grpc: client}
}

func (c Client) Ping() (string, error) {
	c.logger.Info("ping")
	msg, err := c.grpc.HandlePing(c.context, &grpc.PingRequest{})
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	return msg.Message, nil
}

func (c Client) UserExist(username string) (bool, error) {
	c.logger.Info("user exist")
	user, err := c.grpc.HandleUserExist(c.context, &grpc.UserExistRequest{Username: username})
	if err != nil {
		c.logger.Error(err)
		return user.Exist, err
	}
	return user.Exist, nil
}

func (c Client) Registration(username, password string) (model.Token, error) {
	c.logger.Info("registration")
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

func (c Client) Authentication(username, password string) (model.Token, error) {
	c.logger.Info("authentication")
	token := model.Token{}
	password, err := encryption.HashPassword(password)
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	authenticatedUser, err := c.grpc.HandleAuthentication(c.context, &grpc.AuthenticationRequest{Username: username, Password: password})
	if err != nil {
		c.logger.Error(err)
		return token, err
	}
	created, _ := service.ConvertTimestampToTime(authenticatedUser.AccessToken.CreatedAt)
	endDate, _ := service.ConvertTimestampToTime(authenticatedUser.AccessToken.EndDateAt)
	token = model.Token{AccessToken: authenticatedUser.AccessToken.Token, UserID: authenticatedUser.AccessToken.UserId,
		CreatedAt: created, EndDateAt: endDate}
	return token, nil
}

func (c Client) CreateText(name, description, password, plaintext string, token model.Token) error {
	c.logger.Info("create text")
	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptText, err := encryption.Encrypt(plaintext, secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	createdText, err := c.grpc.HandleCreateText(context.Background(),
		&grpc.CreateTextRequest{Name: name, Description: description, Text: []byte(encryptText),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return err
	}
	c.logger.Debug(createdText.Text)
	return nil
}

func (c Client) Synchronization(password string, token model.Token) ([][]string, [][]string, error) {
	dataTblText := [][]string{}
	dataTblCart := [][]string{}
	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	nodes, err := c.grpc.HandleGetListText(c.context,
		&grpc.GetListTextRequest{AccessToken: &grpc.Token{Token: token.AccessToken,
			UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCart, err
	}

	var plaintext string

	titleText := []string{"ID", "NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"}
	titleCart := []string{"NAME", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC",
		"END DATE", "CREATED_AT", "UPDATED_AT"}
	dataTblText = append(dataTblText, titleText)
	dataTblCart = append(dataTblCart, titleCart)
	dataTblTextPointer := &dataTblText

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	for _, node := range nodes.Node {
		plaintext, err = encryption.Decrypt(string(node.Text), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCart, err
		}
		index := table.GetIndexText(dataTblText, table.ColId, strconv.Itoa(int(node.Id)))
		if index == 0 { // entity_id does not exist, add record
			table.AppendText(node, dataTblTextPointer, plaintext)
		} else { // entity_id exists, update tags
			table.UpdateText(node, dataTblTextPointer, index)
		}
	}
	table.DeleteTextColId(dataTblTextPointer)
	logrus.Debug(dataTblText)

	return dataTblText, dataTblCart, nil
}
