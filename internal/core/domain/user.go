package domain

import "time"

type User struct {
	UID               string    `json:"id" map:"id"`
	Nickname          string    `json:"nickname" map:"nickname"`
	UserRole          string    `json:"user_role" map:"user_role"`
	SystemStatus      string    `json:"system_status" map:"system_status"`
	CreatedAt         time.Time `json:"created_at" map:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" map:"updated_at"`
	LastLoginAt       time.Time `json:"last_login_at" map:"last_login_at"`
	LastActivityAt    time.Time `json:"last_activity_at" map:"last_activity_at"`
	PasswordChangedAt time.Time `json:"password_changed_at" map:"password_changed_at"`
	DeletedAt         time.Time `json:"deleted_at" map:"deleted_at"`
}
type UserAggregate struct {
	*User
	*PrivateData
	*SecurityData
}
