package mapper

import (
	"github.com/DENFNC/awq_user_service/internal/core/domain"
	"github.com/DENFNC/awq_user_service/internal/infra/postgres/dao"
)

func SecurityDataToDAO(data *domain.SecurityData) *dao.SecurityDataDAO {
	return &dao.SecurityDataDAO{
		UserID:             toNullString(data.UserID),
		Login:              toNullString(data.Login),
		Email:              toNullString(data.Email),
		PasswordHash:       toNullString(data.PasswordHash),
		LockoutUntil:       toNullTimestamp(data.LockoutUntil),
		ErrorLoginAttempts: toNullInt16(int16(data.ErrorLoginAttempts)),
		IsDeleted:          toNullBool(data.IsDeleted),
		DeletedAt:          toNullTimestamp(data.DeletedAt),
	}
}
