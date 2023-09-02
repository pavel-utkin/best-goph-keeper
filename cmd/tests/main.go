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

	//----------------------------------------------------------------- text

	randName1 := randomizer.RandStringRunes(10)
	randDescription1 := randomizer.RandStringRunes(10)
	plaintext := "Hi my sweetly friends!!!!!!!TeST ВСЕМПРИВЕТ!"

	secretKey := encryption.AesKeySecureRandom([]byte(password))
	encryptText, err := encryption.Encrypt(plaintext, secretKey)
	if err != nil {
		log.Fatal(err)
	}
	createdText, err := client.HandleCreateText(context.Background(),
		&gophkeeper.CreateTextRequest{Name: randName1, Description: randDescription1, Text: []byte(encryptText), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText.Text)

	getNodeText, err := client.HandleGetNodeText(context.Background(), &gophkeeper.GetNodeTextRequest{Name: randName1, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	plaintext, err = encryption.Decrypt(string(getNodeText.Text.Data), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(plaintext)

	randName2 := randomizer.RandStringRunes(10)
	randDescription2 := randomizer.RandStringRunes(10)
	createdText2, err := client.HandleCreateText(context.Background(),
		&gophkeeper.CreateTextRequest{Name: randName2, Description: randDescription2, Text: []byte(encryptText), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdText2.Text)

	getListText, err := client.HandleGetListText(context.Background(), &gophkeeper.GetListTextRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListText)

	deleteText, err := client.HandleDeleteText(context.Background(), &gophkeeper.DeleteTextRequest{Name: randName2, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(deleteText)

	updateText, err := client.HandleUpdateText(context.Background(), &gophkeeper.UpdateTextRequest{Name: randName1, Data: []byte("update text"), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(updateText)

	getListText, err = client.HandleGetListText(context.Background(), &gophkeeper.GetListTextRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListText)

	//----------------------------------------------------------------- card
	randName3 := randomizer.RandStringRunes(10)
	randDescription3 := randomizer.RandStringRunes(10)
	card := model.Card{Name: randName3, Description: randDescription3, PaymentSystem: randName3, Number: randName3, Holder: randName3, EndData: time.Now(), CVC: 13579}
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
		&gophkeeper.CreateCardRequest{Name: randName3, Description: randDescription3, Data: []byte(encryptCard), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdCard.Data)

	getNodeCard, err := client.HandleGetNodeCard(context.Background(), &gophkeeper.GetNodeCardRequest{Name: randName3, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	plaintext, err = encryption.Decrypt(string(getNodeCard.Data.Data), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(plaintext)

	randName4 := randomizer.RandStringRunes(10)
	randDescription4 := randomizer.RandStringRunes(10)
	card = model.Card{Name: randName4, Description: randDescription4, PaymentSystem: randName4, Number: randName4, Holder: randName4, EndData: time.Now(), CVC: 13579}
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
		&gophkeeper.CreateCardRequest{Name: randName4, Data: []byte(encryptCard), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdCard2.Data)

	getListCard, err := client.HandleGetListCard(context.Background(), &gophkeeper.GetListCardRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListCard)

	deleteCard, err := client.HandleDeleteCard(context.Background(), &gophkeeper.DeleteCardRequest{Name: randName4, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(deleteCard)

	updateCard, err := client.HandleUpdateCard(context.Background(), &gophkeeper.UpdateCardRequest{Name: randName3, Data: []byte("update card"), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(updateCard)

	getListCard, err = client.HandleGetListCard(context.Background(), &gophkeeper.GetListCardRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListCard)
	//----------------------------------------------------------------- login password
	randName5 := randomizer.RandStringRunes(10)
	randDescription5 := randomizer.RandStringRunes(10)
	loginPassword := model.LoginPassword{Name: randName5, Description: randDescription5, Login: "Login", Password: "Password"}
	jsonLoginPassword, err := json.Marshal(loginPassword)
	if err != nil {
		log.Fatal(err)
	}

	secretKey = encryption.AesKeySecureRandom([]byte(password))
	encryptLoginPassword, err := encryption.Encrypt(string(jsonLoginPassword), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	createdLoginPassword, err := client.HandleCreateLoginPassword(context.Background(),
		&gophkeeper.CreateLoginPasswordRequest{Name: randName5, Description: randName5, Data: []byte(encryptLoginPassword), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdLoginPassword.Data)

	getNodeLoginPassword, err := client.HandleGetNodeLoginPassword(context.Background(), &gophkeeper.GetNodeLoginPasswordRequest{Name: randName5, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	plaintext, err = encryption.Decrypt(string(getNodeLoginPassword.Data.Data), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(plaintext)

	randName6 := randomizer.RandStringRunes(10)
	randDescription6 := randomizer.RandStringRunes(10)
	loginPassword = model.LoginPassword{Name: randName6, Description: randDescription6, Login: "Login", Password: "Password"}
	jsonLoginPassword, err = json.Marshal(loginPassword)
	if err != nil {
		log.Fatal(err)
	}

	secretKey = encryption.AesKeySecureRandom([]byte(password))
	encryptLoginPassword, err = encryption.Encrypt(string(jsonLoginPassword), secretKey)
	if err != nil {
		log.Fatal(err)
	}
	createdLoginPassword2, err := client.HandleCreateLoginPassword(context.Background(),
		&gophkeeper.CreateLoginPasswordRequest{Name: randName6, Description: randName6, Data: []byte(encryptLoginPassword), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(createdLoginPassword2.Data)

	getListLoginPassword, err := client.HandleGetListLoginPassword(context.Background(), &gophkeeper.GetListLoginPasswordRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListLoginPassword)

	deleteLoginPassword, err := client.HandleDeleteLoginPassword(context.Background(), &gophkeeper.DeleteLoginPasswordRequest{Name: randName6, AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(deleteLoginPassword)

	updateLoginPassword, err := client.HandleUpdateLoginPassword(context.Background(), &gophkeeper.UpdateLoginPasswordRequest{Name: randName5, Data: []byte("update login password"), AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(updateLoginPassword)

	getListLoginPassword, err = client.HandleGetListLoginPassword(context.Background(), &gophkeeper.GetListLoginPasswordRequest{AccessToken: authenticatedUser.AccessToken})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(getListLoginPassword)
}
