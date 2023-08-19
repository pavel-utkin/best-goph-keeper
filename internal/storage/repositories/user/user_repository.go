package user

import (
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/model"
	"time"
)

type Constructor interface {
	Registration(data *model.UserRequest) error
	Authentication(user *model.UserRequest) (bool, error)
	UserExists(user *model.UserRequest) (bool, error)
}

type User struct {
	db *database.DB
}

func New(db *database.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Registration(user *model.UserRequest) (int, error) {
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

func (u *User) Authentication(user *model.UserRequest) (bool, error) {
	var authentication bool
	row := u.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM users where username = $1 AND password = $2)", user.Username, user.Password)
	if err := row.Scan(&authentication); err != nil {
		return authentication, err
	}
	return authentication, nil
}

func (u *User) UserExists(user *model.UserRequest) (bool, error) {
	var exists bool
	row := u.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM users where username = $1)", user.Username)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}
