package domain

import "time"

type User struct {
	UID               string    `db:"id"`
	Nickname          string    `db:"nickname"`
	UserRole          string    `db:"user_role"`
	SystemStatus      string    `db:"system_status"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
	LastLoginAt       time.Time `db:"last_login_at"`
	LastActivityAt    time.Time `db:"last_activity_at"`
	PasswordChangedAt time.Time `db:"password_changed_at"`
	DeletedAt         time.Time `db:"deleted_at"`
	IsDeleted         bool      `db:"is_deleted"`
}

type UserAggregate struct {
	*User
	*PrivateData
	*SecurityData
}
