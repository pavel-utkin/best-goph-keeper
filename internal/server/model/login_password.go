package model

import (
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"time"
)

type LoginPassword struct {
	ID        int64
	UserID    int64
	Key       string
	Value     string
	Data      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateLoginPasswordRequest struct {
	UserID      int64
	Name        string
	Description string
	Type        string
	Data        []byte
	AccessToken string
}

type GetNodeLoginPasswordRequest struct {
	UserID      int64
	Key         string
	Value       string
	AccessToken string
}

func GetLoginPassword(data *LoginPassword) *grpc.LoginPassword {
	created, _ := service.ConvertTimeToTimestamp(data.CreatedAt)
	updated, _ := service.ConvertTimeToTimestamp(data.UpdatedAt)
	deleted, _ := service.ConvertTimeToTimestamp(data.DeletedAt)
	return &grpc.LoginPassword{
		UserId:    data.UserID,
		Data:      data.Data,
		CreatedAt: created,
		UpdatedAt: updated,
		DeletedAt: deleted,
	}
}

func GetListLoginPassword(data []LoginPassword) []*grpc.LoginPassword {
	items := make([]*grpc.LoginPassword, len(data))
	for i := range data {
		created, _ := service.ConvertTimeToTimestamp(data[i].CreatedAt)
		updated, _ := service.ConvertTimeToTimestamp(data[i].UpdatedAt)
		items[i] = &grpc.LoginPassword{Id: data[i].ID, Key: data[i].Key, Data: data[i].Data, Value: data[i].Value, CreatedAt: created, UpdatedAt: updated}
	}
	return items
}
