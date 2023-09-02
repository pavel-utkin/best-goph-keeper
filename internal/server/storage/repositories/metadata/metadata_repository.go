package metadata

import (
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/server/model"
	"time"
)

type Metadata struct {
	db *database.DB
}

func New(db *database.DB) *Metadata {
	return &Metadata{
		db: db,
	}
}

func (m *Metadata) CreateMetadata(metadataRequest *model.CreateMetadataRequest) (*model.Metadata, error) {
	metadata := &model.Metadata{}
	if err := m.db.Pool.QueryRow(
		"INSERT INTO metadata (entity_id, key, value, type, created_at) "+
			"VALUES ($1, $2, $3, $4, $5) RETURNING metadata_id, key, value",
		metadataRequest.EntityId,
		metadataRequest.Key,
		metadataRequest.Value,
		metadataRequest.Type,
		time.Now(),
	).Scan(&metadata.ID, &metadata.Key, &metadata.Value); err != nil {
		return nil, err
	}
	return metadata, nil
}

func (m *Metadata) DeleteMetadata(metadataRequest model.DeleteMetadataRequest) error {
	metadata := &model.Metadata{}
	layout := "01/02/2006 15:04:05"
	if err := m.db.Pool.QueryRow("UPDATE metadata SET deleted_at = $1 WHERE entity_id = $2 and type = $3 RETURNING entity_id",
		time.Now().Format(layout),
		metadataRequest.EntityId,
		metadataRequest.Type,
	).Scan(&metadata.EntityId); err != nil {
		return err
	}
	return nil
}
