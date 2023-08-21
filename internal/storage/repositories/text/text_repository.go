package text

import (
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/model"
	"best-goph-keeper/internal/storage/errors"
	"database/sql"
	"time"
)

type Text struct {
	db *database.DB
}

func New(db *database.DB) *Text {
	return &Text{
		db: db,
	}
}

func (t *Text) CreateText(textRequest *model.CreateTextRequest) (*model.Text, error) {
	text := &model.Text{}
	if err := t.db.Pool.QueryRow(
		"INSERT INTO text (user_id, metadata_id, text, updated_at, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING text_id, text",
		textRequest.UserID,
		1, //metadata_id - надо изменить
		textRequest.Text,
		time.Now(),
		time.Now(),
	).Scan(&text.ID, &text.Text); err != nil {
		return nil, err
	}
	return text, nil
}

func (t *Text) GetNodeText(textRequest *model.GetNodeTextRequest) (*model.Text, error) {
	text := &model.Text{}
	err := t.db.Pool.QueryRow("SELECT text_id, text FROM text WHERE text_id=$1",
		textRequest.TextId).Scan(
		&text.ID,
		&text.Text,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return text, nil
}
