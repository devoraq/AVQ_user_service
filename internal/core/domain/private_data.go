package domain

import (
	"time"
)

type PrivateData struct {
	UserID      string
	FirstName   string
	LastName    string
	MiddleName  string
	DateOfBirth time.Time
	Gender      string
}
