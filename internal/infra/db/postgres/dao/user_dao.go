package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserDAO struct {
	UID               uuid.UUID
	UserRole          string
	CurrentStatus     string
	SystemStatus      string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastLoginAt       sql.NullTime
	PasswordChangedAt sql.NullTime
	DeletedAt         sql.NullTime
	IsDeleted         bool
}
