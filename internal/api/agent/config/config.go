package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type ConfigClient struct {
	GRPC       string       `env:"GRPC"`
	DebugLevel logrus.Level `env:"DEBUG_LEVEL" envDefault:"debug"`
}

// NewConfigClient - creates a new instance with the configuration for the client
func NewConfigClient(log *logrus.Logger) *ConfigClient {
	// Set default values
	configClient := ConfigClient{
		GRPC: "localhost:8080",
	}

	flag.StringVar(&configClient.GRPC, "g", configClient.GRPC, "Server address")
	flag.Parse()
	err := env.Parse(&configClient)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(configClient)

	return &configClient
}
