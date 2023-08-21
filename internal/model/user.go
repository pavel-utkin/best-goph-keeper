package model

import (
	grpc "best-goph-keeper/internal/api/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt timestamp.Timestamp
	UpdatedAt timestamp.Timestamp
	DeletedAt timestamp.Timestamp
}

type UserRequest struct {
	Username string
	Password string
}

func GetUserData(data *User) *grpc.User {
	return &grpc.User{
		UserId:    data.ID,
		Username:  data.Username,
		CreatedAt: &data.CreatedAt,
		UpdatedAt: &data.UpdatedAt,
		DeletedAt: &data.DeletedAt,
	}
}
