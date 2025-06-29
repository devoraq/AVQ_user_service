package domain

import (
	"time"
)

type PrivateData struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	MiddleName  string    `db:"middle_name"`
	DateOfBirth time.Time `db:"date_of_birth"`
	Gender      string    `db:"gender"`
	DeletedAt   time.Time `db:"deleted_at"`
	IsDeleted   bool      `db:"is_deleted"`
}
