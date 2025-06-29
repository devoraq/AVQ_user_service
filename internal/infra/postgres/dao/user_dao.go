package dao

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/DENFNC/awq_user_service/internal/core/domain"
	"github.com/doug-martin/goqu/v9"
)

type UserDAO struct {
	*slog.Logger
	*sql.DB
	*goqu.DialectWrapper
}

func NewUserDAO(
	log *slog.Logger,
	db *sql.DB,
	dialect *goqu.DialectWrapper,
) *UserDAO {
	return &UserDAO{
		Logger:         log,
		DB:             db,
		DialectWrapper: dialect,
	}
}

func (dao *UserDAO) Insert(
	ctx context.Context,
	user *domain.User,
) (string, error) {
	const op = "dao.UserDAO.Create"

	log := dao.Logger.With("op", op)

	stmt, args, err := dao.DialectWrapper.
		Insert("user").
		Rows(&user).
		Prepared(true).
		ToSQL()
	if err != nil {
		log.Error(
			"Error in sql query generation",
			slog.String("err", err.Error()),
		)
		return "", err
	}

	if _, err := dao.DB.ExecContext(ctx, stmt, args...); err != nil {
		log.Error(
			"Unable to save user",
			slog.String("err", err.Error()),
		)
	}

	return "", nil
}
