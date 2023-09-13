package token

import (
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/storage/layouts"
	"best-goph-keeper/internal/server/database"
	"best-goph-keeper/internal/server/model"
	custom_errors "best-goph-keeper/internal/server/storage/errors"
	"database/sql"
	"errors"
	"time"
)

const lengthToken = 32

type TokenRepository interface {
	Create(user *model.User) (string, error)
}

type Token struct {
	db *database.DB
}

func New(db *database.DB) *Token {
	return &Token{
		db: db,
	}
}

func (t *Token) Create(userID int64, lifetime time.Duration) (*model.Token, error) {
	token := &model.Token{}
	accessToken := encryption.GenerateAccessToken(lengthToken)
	currentTime := time.Now()

	if err := t.db.Pool.QueryRow(
		"INSERT INTO access_token (access_token, user_id, created_at, end_date_at) VALUES ($1, $2, $3, $4) "+
			"RETURNING access_token, user_id, created_at, end_date_at",
		accessToken,
		userID,
		currentTime,
		currentTime.Add(lifetime),
	).Scan(&token.AccessToken, &token.UserID, &token.CreatedAt, &token.EndDateAt); err != nil {
		return nil, err
	}
	return token, nil
}

func (t *Token) Validate(endDate time.Time) bool {
	now := time.Now().Format(layouts.LayoutDateAndTime.ToString())
	end := endDate.Format(layouts.LayoutDateAndTime.ToString())
	valid := end > now
	return valid
}

func (t *Token) GetEndDateToken(accessToken string) (time.Time, error) {
	var end time.Time
	if err := t.db.Pool.QueryRow("SELECT end_date_at FROM access_token where access_token = $1",
		accessToken,
	).Scan(&end); err != nil {
		return end, err
	}
	return end, nil
}

func (t *Token) Block(accessToken string) (string, error) {
	var token string
	if err := t.db.Pool.QueryRow("UPDATE access_token SET end_date_at = $1 "+
		"where access_token = $2 RETURNING access_token",
		time.Now(),
		accessToken,
	).Scan(&token); err != nil {
		return token, err
	}
	return token, nil
}

func (t *Token) BlockAllTokenUser(userID int64) (string, error) {
	var token string
	if err := t.db.Pool.QueryRow("UPDATE access_token SET end_date_at = $1 "+
		"where user_id = $2 RETURNING access_token",
		time.Now(),
		userID,
	).Scan(&token); err != nil {
		return token, err
	}
	return token, nil
}

func (t *Token) GetList(userID int64) ([]model.Token, error) {
	tokens := []model.Token{}
	rows, err := t.db.Pool.Query("SELECT access_token, user_id, created_at, end_date_at FROM access_token "+
		"where user_id = $1",
		userID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tokens, custom_errors.ErrRecordNotFound
		} else {
			return tokens, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		token := model.Token{}
		err = rows.Scan(&token.AccessToken, &token.UserID, &token.CreatedAt, &token.EndDateAt)
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
