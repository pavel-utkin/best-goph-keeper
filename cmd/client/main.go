package main

import (
	"best-goph-keeper/internal/client/api/events"
	"best-goph-keeper/internal/client/config"
	"best-goph-keeper/internal/client/gui"
	gophkeeper "best-goph-keeper/internal/server/proto"
	"context"
	"fmt"
	"fyne.io/fyne/v2/app"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// InterceptorLogger adapts logrus logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l logrus.FieldLogger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make(map[string]any, len(fields)/2)
		i := logging.Fields(fields).Iterator()
		if i.Next() {
			k, v := i.At()
			f[k] = v
		}
		l := l.WithFields(f)

		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg)
		case logging.LevelInfo:
			l.Info(msg)
		case logging.LevelWarn:
			l.Warn(msg)
		case logging.LevelError:
			l.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

// @Title Password Manager best-goph-keeper
// @Description GophKeeper is a client-server system that allows the user to safely and securely store logins, passwords, binary data and other private information.
// @Version 1.0

// @Contact.email pavel@utkin-pro.ru

func main() {
	logger := logrus.New()
	ctx := context.Background()
	clientConfig := config.NewConfig()

	conn, err := grpc.Dial(
		clientConfig.GRPC,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			logging.UnaryClientInterceptor(InterceptorLogger(logger)),
		),
		grpc.WithChainStreamInterceptor(
			logging.StreamClientInterceptor(InterceptorLogger(logger)),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	gophkeeperClient := gophkeeper.NewGophkeeperClient(conn)
	client := events.NewEvent(ctx, clientConfig, logger, gophkeeperClient)
	_, err = client.Ping()
	if err != nil {
		log.Fatal(err)
	}
	application := app.New()
	gui.InitGUI(logger, application, client)
}
