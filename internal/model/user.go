package model

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	DeletedAt time.Time
}

type UserRequest struct {
	Username string
	Password string
}
