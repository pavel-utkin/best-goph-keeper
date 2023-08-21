package model

import "time"

type Card struct {
	ID                int64
	UserID            int64
	MetadataID        int64
	CardPaymentSystem string
	CardNumber        string
	CardHolder        string
	CardEndDate       time.Time
	CVC               int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}
