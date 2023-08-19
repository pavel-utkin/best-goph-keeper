package model

import "time"

type Text struct {
	ID         int
	UserID     int
	MetadataID int
	Text       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type CreateTextRequest struct {
	UserID int
	Text   string
}

type GetListTextRequest struct {
	UserID int
}

type GetNodeTextRequest struct {
	CartName string
}
