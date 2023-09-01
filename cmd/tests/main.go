package main

import (
	"best-goph-keeper/internal/client/config"
	"best-goph-keeper/internal/client/model"
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/service/randomizer"
	gophkeeper "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	log := logrus.New()
	config := config.NewConfig()
	log.SetLevel(config.DebugLevel)

	conn, err := grpc.Dial(config.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	accessToken := model.Token{}
	client := gophkeeper.NewGophkeeperClient(conn)

	resp, err := client.HandlePing(context.Background(), &gophkeeper.PingRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(resp.Message)

	username := randomizer.RandStringRunes(10)
	password := "Пароль-1"

	password, err = encryption.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.HandleRegistration(context.Background(), &gophkeeper.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	authenticatedUser, err := client.HandleAuthentication(context.Background(), &gophkeeper.AuthenticationRequest{Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	created, _ := service.ConvertTimestampToTime(authenticatedUser.AccessToken.CreatedAt)
	endDate, _ := service.ConvertTimestampToTime(authenticatedUser.AccessToken.EndDateAt)

	accessToken = model.Token{AccessToken: authenticatedUser.AccessToken.Token, UserID: authenticatedUser.AccessToken.UserId,
		CreatedAt: created, EndDateAt: endDate}
	log.Info(accessToken)

	randName := randomizer.RandStringRunes(10)
	randDescription := randomizer.RandStringRunes(10)
	plaintext := "Hi my sweetly friends!!!!!!!TeST ВСЕМПРИВЕТ!"

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptText, err := encryption.Encrypt(plaintext, secretKey)
	if err != nil {
		log.Fatal(err)
	}
	createdText, err := client.HandleCreateText(context.Background(),
		&gophkeeper.CreateTextRequest{Name: randName, Description: randDescription, Text: []byte(encryptText), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText.Text)

	getNodeText, err := client.HandleGetNodeText(context.Background(), &gophkeeper.GetNodeTextRequest{Name: randName, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	plaintext, err = encryption.Decrypt(string(getNodeText.Text.Text), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(plaintext)

	randName = randomizer.RandStringRunes(10)
	randDescription = randomizer.RandStringRunes(10)
	createdText2, err := client.HandleCreateText(context.Background(),
		&gophkeeper.CreateTextRequest{Name: randName, Description: randDescription, Text: []byte(encryptText), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText2.Text)

	getListText, err := client.HandleGetListText(context.Background(), &gophkeeper.GetListTextRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListText)
	//-----------------------------------------------------------------
	randName = randomizer.RandStringRunes(10)
	card := model.Card{Name: randName, PaymentSystem: randName, Number: randName, Holder: randName, EndData: time.Now(), CVC: 13579}
	jsonCard, err := json.Marshal(card)
	if err != nil {
		log.Fatal(err)
	}

	secretKey = encryption.AesKeySecureRandom([]byte(password))
	encryptCard, err := encryption.Encrypt(string(jsonCard), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	createdCard, err := client.HandleCreateCard(context.Background(),
		&gophkeeper.CreateCardRequest{Name: randName, Data: []byte(encryptCard), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdCard.Data)

	getNodeCard, err := client.HandleGetNodeCard(context.Background(), &gophkeeper.GetNodeCardRequest{Name: randName, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	plaintext, err = encryption.Decrypt(string(getNodeCard.Data.Data), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(plaintext)

	randName = randomizer.RandStringRunes(10)
	card = model.Card{Name: randName, PaymentSystem: randName, Number: randName, Holder: randName, EndData: time.Now(), CVC: 13579}
	jsonCard, err = json.Marshal(card)
	if err != nil {
		log.Fatal(err)
	}

	secretKey = encryption.AesKeySecureRandom([]byte(password))
	encryptCard, err = encryption.Encrypt(string(jsonCard), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	createdCard2, err := client.HandleCreateCard(context.Background(),
		&gophkeeper.CreateCardRequest{Name: randName, Data: []byte(encryptCard), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdCard2.Data)

	getListCard, err := client.HandleGetListCard(context.Background(), &gophkeeper.GetListCardRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListCard)
}
