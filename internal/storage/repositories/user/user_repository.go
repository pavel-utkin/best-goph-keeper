package user

import (
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/model"
	"best-goph-keeper/internal/storage/errors"
	"database/sql"
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

func (u *User) Registration(user *model.UserRequest) (*model.User, error) {
	registeredUser := &model.User{}
	if err := u.db.Pool.QueryRow(
		"INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3) RETURNING user_id, username",
		user.Username,
		user.Password,
		time.Now(),
	).Scan(&registeredUser.ID, &registeredUser.Username); err != nil {
		return nil, err
	}
	return registeredUser, nil
}

func (u *User) Authentication(userRequest *model.UserRequest) (*model.User, error) {
	authenticatedUser := &model.User{}
	err := u.db.Pool.QueryRow("SELECT user_id, username FROM users WHERE username=$1 and password=$2",
		userRequest.Username, userRequest.Password).Scan(
		&authenticatedUser.ID,
		&authenticatedUser.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrWrongUsernameOrPassword
		} else {
			return nil, err
		}
	}
	return authenticatedUser, nil
}

func (u *User) UserExists(user *model.UserRequest) (bool, error) {
	var exists bool
	row := u.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM users where username = $1)", user.Username)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}
