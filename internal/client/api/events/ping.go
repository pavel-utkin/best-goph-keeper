package events

import grpc "best-goph-keeper/internal/server/proto"

func (c Event) Ping() (string, error) {
	c.logger.Info("ping")

	msg, err := c.grpc.Ping(c.context, &grpc.PingRequest{})
	if err != nil {
		c.logger.Error(err)
		return "", err
	}

	return msg.Message, nil
}
