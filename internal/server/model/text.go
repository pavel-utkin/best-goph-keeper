package model

import (
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"time"
)

type Text struct {
	ID        int64
	UserID    int64
	Key       string
	Value     string
	Text      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateTextRequest struct {
	UserID      int64
	Name        string
	Description string
	Type        string
	Text        []byte
	AccessToken string
}

type CreateTextResponse struct {
	Text Text
}

type GetNodeTextRequest struct {
	UserID      int64
	Key         string
	Value       string
	AccessToken string
}

type GetNodeTextResponse struct {
	Key   string
	Value string
	Text  Text
}

type GetListTextRequest struct {
	UserID      int64
	AccessToken string
}

type GetListTextResponse struct {
	Text []Text
}

func GetTextData(data *Text) *grpc.Text {

	created, _ := service.ConvertTimeToTimestamp(data.CreatedAt)
	updated, _ := service.ConvertTimeToTimestamp(data.UpdatedAt)
	deleted, _ := service.ConvertTimeToTimestamp(data.DeletedAt)

	return &grpc.Text{
		UserId:    data.UserID,
		Text:      data.Text,
		CreatedAt: created,
		UpdatedAt: updated,
		DeletedAt: deleted,
	}
}

func GetListData(data []Text) []*grpc.Text {
	items := make([]*grpc.Text, len(data))
	for i := range data {
		created, _ := service.ConvertTimeToTimestamp(data[i].CreatedAt)
		updated, _ := service.ConvertTimeToTimestamp(data[i].UpdatedAt)
		items[i] = &grpc.Text{Id: data[i].ID, Key: data[i].Key, Text: data[i].Text, Value: data[i].Value, CreatedAt: created, UpdatedAt: updated}
	}
	return items
}
