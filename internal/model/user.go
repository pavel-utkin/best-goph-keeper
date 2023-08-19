package model

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	DeletedAt time.Time
}

type LoginRequest struct {
	Username string
	Password string
}

type RegistrationRequest struct {
	Username string
	Password string
}
