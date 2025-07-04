package dto

import (
	"time"
)

type CreateUserDTO struct {
	Nickname string    `json:"nickname"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"Birthday"`
}
