package domain

import "time"

type User struct {
	UID               string
	UserRole          string
	CurrentStatus     string
	SystemStatus      string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastLoginAt       time.Time
	PasswordChangedAt time.Time
	DeletedAt         time.Time
	IsDeleted         bool
}
