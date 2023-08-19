package model

import "time"

type Card struct {
	ID                int
	UserID            int
	MetadataID        int
	CardPaymentSystem string
	CardNumber        string
	CardHolder        string
	CardEndDate       time.Time
	CVC               int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}
