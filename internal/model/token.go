package model

import "github.com/golang/protobuf/ptypes/timestamp"

type Token struct {
	AccessToken string
	UserID      int64
	CreatedAt   timestamp.Timestamp
	EndDateAt   timestamp.Timestamp
}
