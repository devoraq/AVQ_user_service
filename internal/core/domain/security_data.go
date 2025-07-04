package domain

import "time"

type SecurityData struct {
	UserID             string
	Login              string
	Email              string
	PasswordHash       string
	LockoutUntil       time.Time
	ErrorLoginAttempts int
}
