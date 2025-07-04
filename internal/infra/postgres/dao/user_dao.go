package dao

import (
	"database/sql"
)

type User struct {
	UID               sql.NullString `db:"id" goqu:"omitempty" map:"id"`
	Nickname          sql.NullString `db:"nickname" goqu:"omitempty" map:"nickname"`
	UserRole          sql.NullString `db:"user_role" goqu:"omitempty" map:"user_role"`
	SystemStatus      sql.NullString `db:"system_status" goqu:"omitempty" map:"system_status"`
	CreatedAt         sql.NullTime   `db:"created_at" goqu:"omitempty" map:"created_at"`
	UpdatedAt         sql.NullTime   `db:"updated_at" goqu:"omitempty" map:"updated_at"`
	LastLoginAt       sql.NullTime   `db:"last_login_at" goqu:"omitempty" map:"last_login_at"`
	LastActivityAt    sql.NullTime   `db:"last_activity_at" goqu:"omitempty" map:"last_activity_at"`
	PasswordChangedAt sql.NullTime   `db:"password_changed_at" goqu:"omitempty" map:"password_changed_at"`
	DeletedAt         sql.NullTime   `db:"deleted_at" goqu:"omitempty" map:"deleted_at"`
}
