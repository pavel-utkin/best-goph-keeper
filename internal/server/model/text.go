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
	Data      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateTextRequest struct {
	UserID      int64
	Name        string
	Description string
	Type        string
	Data        []byte
	AccessToken string
}

type CreateTextResponse struct {
	Text Text
}

// ----------------------------------------
type GetNodeTextRequest struct {
	UserID      int64
	Key         string
	Value       string
	AccessToken string
}

type GetNodeTextResponse struct {
	Key   string
	Value string
	Data  Text
}

// ----------------------------------------
type GetListTextRequest struct {
	UserID      int64
	AccessToken string
}

type GetListTextResponse struct {
	Text []Text
}

func GetText(text *Text) *grpc.Text {
	created, _ := service.ConvertTimeToTimestamp(text.CreatedAt)
	updated, _ := service.ConvertTimeToTimestamp(text.UpdatedAt)
	deleted, _ := service.ConvertTimeToTimestamp(text.DeletedAt)
	return &grpc.Text{
		UserId:    text.UserID,
		Data:      text.Data,
		CreatedAt: created,
		UpdatedAt: updated,
		DeletedAt: deleted,
	}
}

func GetListText(text []Text) []*grpc.Text {
	items := make([]*grpc.Text, len(text))
	for i := range text {
		created, _ := service.ConvertTimeToTimestamp(text[i].CreatedAt)
		updated, _ := service.ConvertTimeToTimestamp(text[i].UpdatedAt)
		items[i] = &grpc.Text{Id: text[i].ID, Key: text[i].Key, Data: text[i].Data, Value: text[i].Value, CreatedAt: created, UpdatedAt: updated}
	}
	return items
}
