package metadata

import (
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/model"
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
