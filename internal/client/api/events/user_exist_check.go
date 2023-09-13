package events

import grpc "best-goph-keeper/internal/server/proto"

// UserExist - check if user exist in db
func (c Event) UserExist(username string) (bool, error) {
	c.logger.Info("user exist check")

	user, err := c.grpc.UserExist(c.context, &grpc.UserExistRequest{Username: username})
	if err != nil {
		c.logger.Error(err)
		return user.Exist, err
	}

	return user.Exist, nil
}
