package text

import (
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/server/model"
	"best-goph-keeper/internal/server/storage/errors"
	"best-goph-keeper/internal/server/storage/vars"
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
		//"INSERT INTO text (user_id, text, created_at) VALUES ($1, $2, $3) "+
		"INSERT INTO text (user_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4) "+
			"RETURNING text_id, text",
		textRequest.UserID,
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
	err := t.db.Pool.QueryRow("SELECT text.text FROM metadata "+
		"inner join text on metadata.entity_id = text.text_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3",
		string(vars.Name), textRequest.Value, textRequest.UserID).Scan(
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

func (t *Text) GetListText(userId int64) ([]model.Text, error) {
	ListText := []model.Text{}

	rows, err := t.db.Pool.Query("SELECT metadata.entity_id, metadata.key, text.text, metadata.value, text.created_at, "+
		"text.updated_at FROM metadata "+
		//rows, err := t.db.Pool.Query("SELECT text.text FROM metadata "+
		"inner join text on metadata.entity_id = text.text_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where users.user_id = $1", userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		text := model.Text{}
		err = rows.Scan(&text.ID, &text.Key, &text.Text, &text.Value, &text.CreatedAt, &text.UpdatedAt)
		if err != nil {
			return nil, err
		}
		ListText = append(ListText, text)
	}
	return ListText, nil
}

func (t *Text) KeyExists(textRequest *model.CreateTextRequest) (bool, error) {
	var exists bool
	row := t.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM metadata "+
		"inner join text on metadata.entity_id = text.text_id "+
		"inner join users on text.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3)", string(vars.Name), textRequest.Name, textRequest.UserID)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}
