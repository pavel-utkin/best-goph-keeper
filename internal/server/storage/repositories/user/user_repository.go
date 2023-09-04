package user

import (
	"best-goph-keeper/internal/server/database"
	"best-goph-keeper/internal/server/model"
	"best-goph-keeper/internal/server/storage/errors"
	"database/sql"
	"time"
)

type Constructor interface {
	Registration(data *model.UserRequest) error
	Authentication(user *model.UserRequest) (bool, error)
	UserExists(username string) (bool, error)
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
	err := u.db.Pool.QueryRow("SELECT user_id, username FROM users WHERE username=$1 and password=$2 and deleted_at IS NULL",
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

func (u *User) UserExists(username string) (bool, error) {
	var exists bool
	row := u.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM users where username = $1)", username)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (u *User) UserList() ([]model.GetAllUsers, error) {
	users := []model.GetAllUsers{}
	rows, err := u.db.Pool.Query("SELECT username, deleted_at FROM users")
	if err != nil {
		if err == sql.ErrNoRows {
			return users, errors.ErrRecordNotFound
		} else {
			return users, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		user := model.GetAllUsers{}

		err = rows.Scan(&user.Username, &user.DeletedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) Block(username string) (int64, error) {
	var id int64
	if err := u.db.Pool.QueryRow("UPDATE users SET deleted_at = $1 "+
		"where username = $2 RETURNING user_id",
		time.Now(),
		username,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (u *User) Unblock(username string) (int64, error) {
	var id int64
	if err := u.db.Pool.QueryRow("UPDATE users SET deleted_at = null "+
		"where username = $1 RETURNING user_id",
		username,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (u *User) GetUserID(username string) (int64, error) {
	var id int64
	if err := u.db.Pool.QueryRow("SELECT user_id FROM users where username = $1",
		username,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
