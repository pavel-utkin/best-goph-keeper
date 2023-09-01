package server

import (
	grpcHandler "best-goph-keeper/internal/server/api/handlers"
	"best-goph-keeper/internal/server/config"
	grpcGophkeeper "best-goph-keeper/internal/server/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

// StartService - starts the gophkeeper server
func StartService(grpcHandler *grpcHandler.Handler, config *config.ConfigServer, log *logrus.Logger) {
	log.Infof("Start gophkeeper application %s ", config.GRPC)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", config.GRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcGophkeeper.RegisterGophkeeperServer(grpcServer, grpcHandler)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("gprc server: %v", err)
	}
}
