package card

import (
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/server/model"
	"best-goph-keeper/internal/server/storage/errors"
	"best-goph-keeper/internal/server/storage/vars"
	"database/sql"
	"time"
)

type Card struct {
	db *database.DB
}

func New(db *database.DB) *Card {
	return &Card{
		db: db,
	}
}

func (c *Card) CreateCard(cardRequest *model.CreateCardRequest) (*model.Card, error) {
	card := &model.Card{}
	if err := c.db.Pool.QueryRow(
		"INSERT INTO card (user_id, data, created_at, updated_at) VALUES ($1, $2, $3, $4) "+
			"RETURNING card_id, data",
		cardRequest.UserID,
		cardRequest.Data,
		time.Now(),
		time.Now(),
	).Scan(&card.ID, &card.Data); err != nil {
		return nil, err
	}
	return card, nil
}

func (c *Card) GetIdCard(value string, userID int64) (int64, error) {
	var cardID int64
	err := c.db.Pool.QueryRow("SELECT card.card_id FROM metadata "+
		"inner join card on metadata.entity_id = card.card_id "+
		"inner join users on card.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4 and card.deleted_at IS NULL",
		string(vars.Name), value, userID, string(vars.Card)).
		Scan(&cardID)
	if err != nil {
		if err == sql.ErrNoRows {
			return cardID, errors.ErrRecordNotFound
		} else {
			return cardID, err
		}
	}
	return cardID, nil
}

func (c *Card) KeyExists(cardRequest *model.CreateCardRequest) (bool, error) {
	var exists bool
	row := c.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM metadata "+
		"inner join card on metadata.entity_id = card.card_id "+
		"inner join users on card.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4 and card.deleted_at IS NULL)",
		string(vars.Name), cardRequest.Name, cardRequest.UserID, string(vars.Card))
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (c *Card) GetNodeCard(cardRequest *model.GetNodeCardRequest) (*model.Card, error) {
	card := &model.Card{}
	err := c.db.Pool.QueryRow("SELECT card.data FROM metadata "+
		"inner join card on metadata.entity_id = card.card_id "+
		"inner join users on card.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4 and card.deleted_at IS NULL",
		string(vars.Name), cardRequest.Value, cardRequest.UserID, string(vars.Card)).Scan(
		&card.Data,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return card, nil
}

func (c *Card) GetListCard(userId int64) ([]model.Card, error) {
	listCard := []model.Card{}

	rows, err := c.db.Pool.Query("SELECT metadata.entity_id, metadata.key, card.data, metadata.value, card.created_at, "+
		"card.updated_at FROM metadata "+
		"inner join card on metadata.entity_id = card.card_id "+
		"inner join users on card.user_id  = users.user_id "+
		"where users.user_id = $1 and metadata.type = $2 and card.deleted_at IS NULL",
		userId, string(vars.Card))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		card := model.Card{}
		err = rows.Scan(&card.ID, &card.Key, &card.Data, &card.Value, &card.CreatedAt, &card.UpdatedAt)
		if err != nil {
			return nil, err
		}
		listCard = append(listCard, card)
	}
	return listCard, nil
}

func (c *Card) DeleteCard(entityId int64) error {
	var id int64
	if err := c.db.Pool.QueryRow("UPDATE card SET deleted_at = $1 WHERE card_id = $2 RETURNING card_id",
		time.Now(),
		entityId,
	).Scan(&id); err != nil {
		return err
	}
	return nil
}

func (lp *Card) UpdateCard(textID int64, data []byte) error {
	var id int64
	if err := lp.db.Pool.QueryRow("UPDATE card SET data = $1, updated_at = $2 WHERE card_id = $3 RETURNING card_id",
		data,
		time.Now(),
		textID,
	).Scan(&id); err != nil {
		return err
	}
	return nil
}
