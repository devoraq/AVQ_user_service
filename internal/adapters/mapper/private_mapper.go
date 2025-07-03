package mapper

import (
	"github.com/DENFNC/awq_user_service/internal/core/domain"
	"github.com/DENFNC/awq_user_service/internal/infra/postgres/dao"
)

func PrivateDataToDAO(data *domain.PrivateData) *dao.PrivateDataDAO {
	return &dao.PrivateDataDAO{
		UserID:      toNullString(data.UserID),
		FirstName:   toNullString(data.FirstName),
		LastName:    toNullString(data.LastName),
		MiddleName:  toNullString(data.LastName),
		DateOfBirth: toNullTimestamp(data.DateOfBirth),
		Gender:      toNullString(data.Gender),
		DeletedAt:   toNullTimestamp(data.DeletedAt),
		IsDeleted:   toNullBool(data.IsDeleted),
	}
}
