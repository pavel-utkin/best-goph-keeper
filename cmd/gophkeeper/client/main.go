package client

import (
	config2 "best-goph-keeper/internal/api/agent/config"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	config := config2.NewConfigClient(log)
	log.SetLevel(config.DebugLevel)
}
