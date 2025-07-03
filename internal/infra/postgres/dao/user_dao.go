package dao

import (
	"database/sql"
)

type UserDAO struct {
	UID               sql.NullString `db:"id" goqu:"omitempty"`
	Nickname          sql.NullString `db:"nickname" goqu:"omitempty"`
	UserRole          sql.NullString `db:"user_role" goqu:"omitempty"`
	SystemStatus      sql.NullString `db:"system_status" goqu:"omitempty"`
	CreatedAt         sql.NullTime   `db:"created_at" goqu:"omitempty"`
	UpdatedAt         sql.NullTime   `db:"updated_at" goqu:"omitempty"`
	LastLoginAt       sql.NullTime   `db:"last_login_at" goqu:"omitempty"`
	LastActivityAt    sql.NullTime   `db:"last_activity_at" goqu:"omitempty"`
	PasswordChangedAt sql.NullTime   `db:"password_changed_at" goqu:"omitempty"`
	DeletedAt         sql.NullTime   `db:"deleted_at" goqu:"omitempty"`
	IsDeleted         sql.NullBool   `db:"is_deleted" goqu:"omitempty"`
}
