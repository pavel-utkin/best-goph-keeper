package model

import (
	grpc "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/service"
	"encoding/json"
	"time"
)

type Entity struct {
	ID        int64
	UserID    int64
	Data      []byte
	Metadata  MetadataEntity
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type MetadataEntity struct {
	Name        string
	Description string
	Type        string
}

type CreateEntityRequest struct {
	UserID      int64
	Data        []byte
	Metadata    MetadataEntity
	AccessToken string
}

func GetListEntity(data []Entity) ([]*grpc.Entity, error) {
	items := make([]*grpc.Entity, len(data))
	for i := range data {
		jsonMetadata, err := json.Marshal(data[i].Metadata)
		if err != nil {
			return items, err
		}
		created := service.ConvertTimeToTimestamp(data[i].CreatedAt)
		updated := service.ConvertTimeToTimestamp(data[i].UpdatedAt)
		items[i] = &grpc.Entity{Id: data[i].ID, UserId: data[i].UserID, Data: data[i].Data,
			Metadata: string(jsonMetadata), CreatedAt: created, UpdatedAt: updated}
	}
	return items, nil
}
