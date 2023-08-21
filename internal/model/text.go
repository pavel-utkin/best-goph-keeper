package model

import (
	grpc "best-goph-keeper/internal/api/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Text struct {
	ID        int64
	UserID    int64
	Text      []byte
	CreatedAt timestamp.Timestamp
	UpdatedAt timestamp.Timestamp
	DeletedAt timestamp.Timestamp
}

type CreateTextRequest struct {
	UserID      int64
	Key         string
	Value       string
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
	return &grpc.Text{
		UserId:    data.UserID,
		Text:      data.Text,
		CreatedAt: &data.CreatedAt,
		UpdatedAt: &data.UpdatedAt,
		DeletedAt: &data.DeletedAt,
	}
}

func GetListData(data []Text) []*grpc.Text {
	items := make([]*grpc.Text, len(data))
	for i := range data {
		items[i] = &grpc.Text{Text: data[i].Text}
	}
	return items
}
