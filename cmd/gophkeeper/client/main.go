package main

import (
	config2 "best-goph-keeper/internal/api/agent/config"
	grpcClient "best-goph-keeper/internal/api/proto"
	"best-goph-keeper/internal/model"
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
	password := "passworD-123"

	registeredUser, err := client.HandleRegistration(context.Background(), &grpcClient.RegistrationRequest{Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	authenticatedUser, err := client.HandleAuthentication(context.Background(), &grpcClient.AuthenticationRequest{Username: registeredUser.Username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	user := model.User{ID: authenticatedUser.UserId, Username: authenticatedUser.Username}
	log.Info(user)

	randText := randomizer.RandStringRunes(10)

	createdText, err := client.HandleCreateText(context.Background(), &grpcClient.CreateTextRequest{UserId: user.ID, Text: randText})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText)
	getNodeText, err := client.HandleGetNodeText(context.Background(), &grpcClient.GetNodeTextRequest{TextId: createdText.TextId})
	if err != nil {
		log.Fatal(err)
	}
	text := model.Text{ID: getNodeText.TextId, Text: getNodeText.Text}
	log.Info(text)
}
