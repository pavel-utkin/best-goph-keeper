package main

import (
	"best-goph-keeper/internal/client/api"
	"best-goph-keeper/internal/client/config"
	"best-goph-keeper/internal/client/gui"
	gophkeeper "best-goph-keeper/internal/server/proto"
	"context"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//---------------------------------------------------------------------- client application init
	log := logrus.New()
	ctx := context.Background()
	config := config.NewConfig(log)
	log.SetLevel(config.DebugLevel)
	conn, err := grpc.Dial(config.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	grpc := gophkeeper.NewGophkeeperClient(conn)
	client := api.NewClient(ctx, log, grpc)
	ping, err := client.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(ping)
	//---------------------------------------------------------------------- fyne.app init
	application := app.New()
	application.Settings().SetTheme(theme.LightTheme())
	gui.InitGUI(log, application, client)
}
