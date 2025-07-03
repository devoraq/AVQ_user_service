package dao

import "database/sql"

type SecurityDataDAO struct {
	UserID             sql.NullString `db:"user_id" goqu:"omitempty"`
	Login              sql.NullString `db:"login" goqu:"omitempty"`
	Email              sql.NullString `db:"email" goqu:"omitempty"`
	PasswordHash       sql.NullString `db:"password_hash" goqu:"omitempty"`
	LockoutUntil       sql.NullTime   `db:"lockout_until" goqu:"omitempty"`
	ErrorLoginAttempts sql.NullInt16  `db:"error_login_attempts" goqu:"omitempty"`
	IsDeleted          sql.NullBool   `db:"is_deleted" goqu:"omitempty"`
	DeletedAt          sql.NullTime   `db:"deleted_at" goqu:"omitempty"`
}
