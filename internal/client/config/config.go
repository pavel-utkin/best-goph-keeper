package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
	"log"
)

type ConfigClient struct {
	GRPC       string       `env:"GRPC" envDefault:"localhost:8080"`
	DebugLevel logrus.Level `env:"DEBUG_LEVEL" envDefault:"debug"`
}

// NewConfigClient - creates a new instance with the configuration for the client
func NewConfig() *ConfigClient {
	configClient := ConfigClient{}
	flag.StringVar(&configClient.GRPC, "g", configClient.GRPC, "Server address")
	flag.Parse()
	err := env.Parse(&configClient)
	if err != nil {
		log.Fatal(err)
	}

	return &configClient
}
