package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/service/table"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (c Event) EventSynchronization(password string, token model.Token) ([][]string, [][]string, [][]string, error) {
	c.logger.Info("Synchronization")

	dataTblText := [][]string{}
	dataTblCard := [][]string{}
	dataTblLoginPassword := [][]string{}
	created, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate, _ := service.ConvertTimeToTimestamp(token.EndDateAt)

	nodesText, err := c.grpc.HandleGetListText(c.context,
		&grpc.GetListTextRequest{AccessToken: &grpc.Token{Token: token.AccessToken,
			UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, err
	}

	nodesCard, err := c.grpc.HandleGetListCard(c.context,
		&grpc.GetListCardRequest{AccessToken: &grpc.Token{Token: token.AccessToken,
			UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, err
	}

	nodesLoginPassword, err := c.grpc.HandleGetListLoginPassword(c.context,
		&grpc.GetListLoginPasswordRequest{AccessToken: &grpc.Token{Token: token.AccessToken,
			UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, err
	}

	var plaintext string
	secretKey := encryption.AesKeySecureRandom([]byte(password))

	titleText := []string{"ID", "NAME", "DESCRIPTION", "DATA", "CREATED AT", "UPDATED AT"}
	titleCard := []string{"ID", "NAME", "DESCRIPTION", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC",
		"END DATE", "CREATED AT", "UPDATED AT"}
	titleLoginPassword := []string{"ID", "NAME", "DESCRIPTION", "LOGIN", "PASSWORD", "CREATED AT", "UPDATED AT"}
	dataTblText = append(dataTblText, titleText)
	dataTblCard = append(dataTblCard, titleCard)
	dataTblLoginPassword = append(dataTblLoginPassword, titleLoginPassword)
	dataTblTextPointer := &dataTblText
	dataTblCardPointer := &dataTblCard
	dataTblLoginPasswordPointer := &dataTblLoginPassword

	for _, node := range nodesText.Node {
		plaintext, err = encryption.Decrypt(string(node.Data), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, err
		}
		index := table.GetIndex(dataTblText, table.ColId, strconv.Itoa(int(node.Id)))
		if index == 0 { // entity_id does not exist, add record
			table.AppendText(node, dataTblTextPointer, plaintext)
		} else { // entity_id exists, update tags
			table.UpdateText(node, dataTblTextPointer, index)
		}
	}

	for _, node := range nodesCard.Node {
		plaintext, err = encryption.Decrypt(string(node.Data), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, err
		}

		var card model.Card
		err = json.Unmarshal([]byte(plaintext), &card)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, err
		}
		index := table.GetIndex(dataTblCard, table.ColId, strconv.Itoa(int(node.Id)))
		if index == 0 { // entity_id does not exist, add record
			table.AppendCard(node, dataTblCardPointer, card)
		} else { // entity_id exists, update tags
			table.UpdateCard(node, dataTblCardPointer, index)
		}
	}

	for _, node := range nodesLoginPassword.Node {
		plaintext, err = encryption.Decrypt(string(node.Data), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, err
		}

		var loginPassword model.LoginPassword
		err = json.Unmarshal([]byte(plaintext), &loginPassword)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, err
		}
		index := table.GetIndex(dataTblLoginPassword, table.ColId, strconv.Itoa(int(node.Id)))
		if index == 0 { // entity_id does not exist, add record
			table.AppendLoginPassword(node, dataTblLoginPasswordPointer, loginPassword)
		} else { // entity_id exists, update tags
			table.UpdateLoginPassword(node, dataTblLoginPasswordPointer, index)
		}
	}

	table.DeleteColId(dataTblTextPointer)
	table.DeleteColId(dataTblCardPointer)
	table.DeleteColId(dataTblLoginPasswordPointer)
	logrus.Debug(dataTblText)
	logrus.Debug(dataTblCard)
	logrus.Debug(dataTblLoginPassword)

	return dataTblText, dataTblCard, dataTblLoginPassword, nil
}
