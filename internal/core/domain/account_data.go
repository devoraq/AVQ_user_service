package domain

import (
	"time"

	"github.com/google/uuid"
)

type AccountData struct {
	ID        uuid.UUID      `db:"id"`
	UserID    uuid.UUID      `db:"user_id"`
	AvatarURL string         `db:"avatar_url"`
	BannerURL string         `db:"banner_url"`
	Bio       string         `db:"bio"`
	Status    string         `db:"status"`
	Socials   map[string]any `db:"socials"`
	IsDeleted bool           `db:"is_deleted"`
	DeletedAt time.Time      `db:"deleted_at"`
}
