package domain

import "time"

type User struct {
	UID               string    `json:"id"`
	Nickname          string    `json:"nickname"`
	UserRole          string    `json:"user_role"`
	SystemStatus      string    `json:"system_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	LastLoginAt       time.Time `json:"last_login_at"`
	LastActivityAt    time.Time `json:"last_activity_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	DeletedAt         time.Time `json:"deleted_at"`
	IsDeleted         bool      `json:"is_deleted"`
}

type UserAggregate struct {
	*User
	*PrivateData
	*SecurityData
}
