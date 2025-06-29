package v1

import "time"

type UserCreateDTO struct {
	Nickname string
	Password string
	Email    string
	Birthday time.Time
}
