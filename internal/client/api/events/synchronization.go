package events

import (
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/service/table"
	"best-goph-keeper/internal/client/storage/labels"
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"best-goph-keeper/internal/server/storage/vars"
	"encoding/json"
)

func (c Event) Synchronization(password string, token model.Token) ([][]string, [][]string, [][]string, [][]string, error) {
	c.logger.Info("synchronization")

	dataTblText := [][]string{}
	dataTblCard := [][]string{}
	dataTblLoginPassword := [][]string{}
	dataTblBinary := [][]string{}

	created := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDate := service.ConvertTimeToTimestamp(token.EndDateAt)
	//-----------------------------------------------
	var plaintext string
	secretKey := encryption.AesKeySecureRandom([]byte(password))

	titleText := []string{labels.NameItem, labels.DescriptionItem, labels.DataItem, labels.CreatedAtItem, labels.UpdatedAtItem}
	titleCard := []string{labels.NameItem, labels.DescriptionItem, labels.PaymentSystemItem, labels.NumberItem, labels.HolderItem,
		labels.CVCItem, labels.EndDateItem, labels.CreatedAtItem, labels.UpdatedAtItem}
	titleLoginPassword := []string{labels.NameItem, labels.DescriptionItem, labels.LoginItem, labels.PasswordItem,
		labels.CreatedAtItem, labels.UpdatedAtItem}
	titleBinary := []string{labels.NameItem, labels.CreatedAtItem}

	dataTblText = append(dataTblText, titleText)
	dataTblCard = append(dataTblCard, titleCard)
	dataTblLoginPassword = append(dataTblLoginPassword, titleLoginPassword)
	dataTblBinary = append(dataTblBinary, titleBinary)

	dataTblTextPointer := &dataTblText
	dataTblCardPointer := &dataTblCard
	dataTblLoginPasswordPointer := &dataTblLoginPassword
	dataTblBinaryPointer := &dataTblBinary

	//-----------------------------------------------
	nodesTextEntity, err := c.grpc.EntityGetList(c.context,
		&grpc.GetListEntityRequest{Type: vars.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken,
				UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}

	for _, node := range nodesTextEntity.Node {
		plaintext, err = encryption.Decrypt(string(node.Data), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		err = table.AppendTextEntity(node, dataTblTextPointer, plaintext)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}

	//-----------------------------------------------
	nodesCardEntity, err := c.grpc.EntityGetList(c.context,
		&grpc.GetListEntityRequest{Type: vars.Card.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken,
				UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}

	for _, node := range nodesCardEntity.Node {
		plaintext, err = encryption.Decrypt(string(node.Data), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}

		var card model.Card
		err = json.Unmarshal([]byte(plaintext), &card)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		err = table.AppendCardEntity(node, dataTblCardPointer, card)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}
	//-----------------------------------------------
	nodesLoginPasswordEntity, err := c.grpc.EntityGetList(c.context,
		&grpc.GetListEntityRequest{Type: vars.LoginPassword.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken,
				UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}

	for _, node := range nodesLoginPasswordEntity.Node {
		plaintext, err = encryption.Decrypt(string(node.Data), secretKey)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}

		var loginPassword model.LoginPassword
		err = json.Unmarshal([]byte(plaintext), &loginPassword)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
		err = table.AppendLoginPasswordEntity(node, dataTblLoginPasswordPointer, loginPassword)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}
	//-----------------------------------------------
	nodesBinary, err := c.grpc.FileGetList(c.context,
		&grpc.GetListBinaryRequest{AccessToken: &grpc.Token{Token: token.AccessToken,
			UserId: token.UserID, CreatedAt: created, EndDateAt: endDate}})
	if err != nil {
		c.logger.Error(err)
		return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
	}

	for _, node := range nodesBinary.Node {
		err = table.AppendBinary(node, dataTblBinaryPointer)
		if err != nil {
			c.logger.Error(err)
			return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, err
		}
	}
	//-----------------------------------------------
	return dataTblText, dataTblCard, dataTblLoginPassword, dataTblBinary, nil
}
