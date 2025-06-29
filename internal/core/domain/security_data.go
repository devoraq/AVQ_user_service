package domain

import "time"

type SecurityData struct {
	ID                 string    `db:"id"`
	UserID             string    `db:"user_id"`
	Login              string    `db:"login"`
	Email              string    `db:"email"`
	PasswordHash       string    `db:"password_hash"`
	LockoutUntil       time.Time `db:"lockout_until"`
	ErrorLoginAttempts int       `db:"error_login_attempts"`
	IsDeleted          bool      `db:"is_deleted"`
	DeletedAt          time.Time `db:"deleted_at"`
}
