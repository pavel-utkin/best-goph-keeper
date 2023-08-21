package main

import (
	config2 "best-goph-keeper/internal/api/agent/config"
	grpcClient "best-goph-keeper/internal/api/proto"
	"best-goph-keeper/internal/model"
	"best-goph-keeper/internal/service/encryption"
	"best-goph-keeper/internal/service/randomizer"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log := logrus.New()
	config := config2.NewConfigClient(log)
	log.SetLevel(config.DebugLevel)

	conn, err := grpc.Dial(config.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := grpcClient.NewGophkeeperClient(conn)
	resp, err := client.HandlePing(context.Background(), &grpcClient.PingRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(resp.Message)

	username := randomizer.RandStringRunes(10)
	password := "MyBestPassword-123"

	password, err = encryption.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	registeredUser, err := client.HandleRegistration(context.Background(), &grpcClient.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	authenticatedUser, err := client.HandleAuthentication(context.Background(), &grpcClient.AuthenticationRequest{Username: registeredUser.User.Username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	user := model.User{ID: authenticatedUser.User.UserId, Username: authenticatedUser.User.Username}
	log.Info(user)

	randName := randomizer.RandStringRunes(10)
	plaintext := randomizer.RandStringRunes(10)

	secretKey := encryption.AesKeySecureRandom([]byte(password))

	encryptText := encryption.Encrypt(plaintext, secretKey)
	if err != nil {
		log.Fatal(err)
	}
	createdText, err := client.HandleCreateText(context.Background(),
		&grpcClient.CreateTextRequest{Key: "Name", Value: randName, Text: []byte(encryptText), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText.Text)

	getNodeText, err := client.HandleGetNodeText(context.Background(), &grpcClient.GetNodeTextRequest{Key: "Name", Value: randName, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	plaintext = encryption.Decrypt(string(getNodeText.Text.Text), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(plaintext)

	createdText2, err := client.HandleCreateText(context.Background(),
		&grpcClient.CreateTextRequest{Key: "Name2", Value: randName, Text: []byte(encryptText), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText2.Text)

	getListText, err := client.HandleGetListText(context.Background(), &grpcClient.GetListTextRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListText)
}
