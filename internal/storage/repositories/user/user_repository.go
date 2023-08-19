package user

import (
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/model"
	"time"
)

type Constructor interface {
	Registration(data *model.RegistrationRequest) error
}

type User struct {
	db *database.DB
}

func New(db *database.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Registration(user *model.RegistrationRequest) (int, error) {
	var id int
	if err := u.db.Pool.QueryRow(
		"INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3) RETURNING user_id",
		user.Username,
		user.Password,
		time.Now(),
	).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}
