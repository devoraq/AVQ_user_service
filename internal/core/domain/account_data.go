package domain

import (
	"github.com/google/uuid"
)

type AccountData struct {
	UserID    uuid.UUID
	AvatarURL string
	BannerURL string
	Bio       string
	Status    string
	Socials   map[string]any
}
