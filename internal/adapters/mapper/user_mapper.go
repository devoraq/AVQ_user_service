package mapper

import (
	"github.com/DENFNC/awq_user_service/internal/core/domain"
	"github.com/DENFNC/awq_user_service/internal/infra/postgres/dao"
)

func UserToDomain(user *dao.UserDAO) *domain.User {
	return &domain.User{
		// UID:               user.UID,
		// Nickname:          user.Nickname,
		// UserRole:          user.UserRole,
		// SystemStatus:      user.SystemStatus,
		// CreatedAt:         user.CreatedAt,
		// UpdatedAt:         user.UpdatedAt,
		// LastLoginAt:       user.LastLoginAt,
		// LastActivityAt:    user.LastActivityAt,
		// PasswordChangedAt: user.PasswordChangedAt,
		// DeletedAt:         user.DeletedAt,
		// IsDeleted:         user.IsDeleted,
	}
}

func UserToDAO(user *domain.User) *dao.UserDAO {
	return &dao.UserDAO{
		UID:               toNullString(user.UID),
		Nickname:          toNullString(user.Nickname),
		UserRole:          toNullString(user.UserRole),
		SystemStatus:      toNullString(user.SystemStatus),
		CreatedAt:         toNullTimestamp(user.CreatedAt),
		UpdatedAt:         toNullTimestamp(user.UpdatedAt),
		LastLoginAt:       toNullTimestamp(user.LastLoginAt),
		LastActivityAt:    toNullTimestamp(user.LastActivityAt),
		PasswordChangedAt: toNullTimestamp(user.PasswordChangedAt),
		DeletedAt:         toNullTimestamp(user.DeletedAt),
		IsDeleted:         toNullBool(user.IsDeleted),
	}
}
