package user

import (
	"time"

	v1 "github.com/DENFNC/awq_user_service/api/gen/go/user/v1"
)

type CreateUserDTO struct {
	Nickname string    `json:"nickname"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"Birthday"`
}

func CreateUserRequestToDTO(req *v1.CreateUserRequest) CreateUserDTO {
	return CreateUserDTO{
		Nickname: req.Nickname,
		Password: req.Password,
		Email:    req.Email,
		Birthday: req.Birthday.AsTime(),
	}
}
