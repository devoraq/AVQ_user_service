package domain

import (
	"time"

	"github.com/google/uuid"
)

type AccountData struct {
	UserID    uuid.UUID
	AvatarURL string
	BannerURL string
	Bio       string
	Status    string
	Socials   map[string]any
	IsDeleted bool
	DeletedAt time.Time
}
