package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/DENFNC/awq_user_service/internal/adapters/mapper"
	"github.com/DENFNC/awq_user_service/internal/core/domain"
	"github.com/DENFNC/awq_user_service/internal/utils/dbutils"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type UserRepository struct {
	*slog.Logger
	*sql.DB
	*goqu.DialectWrapper
}

func NewUserRepository(
	log *slog.Logger,
	db *sql.DB,
) *UserRepository {
	dialect := goqu.Dialect("postgres")

	return &UserRepository{
		Logger:         log,
		DB:             db,
		DialectWrapper: &dialect,
	}
}

func (repo *UserRepository) Save(
	ctx context.Context,
	agg *domain.UserAggregate,
) (string, error) {
	const op = "repository.UserRepository.Save"

	log := repo.Logger.With("op", op)

	daoUser := mapper.UserToDAO(agg.User)
	daoSecData := mapper.SecurityDataToDAO(agg.SecurityData)
	daoPrivData := mapper.PrivateDataToDAO(agg.PrivateData)

	err := dbutils.WithTransaction(ctx, repo.DB, func(tx *sql.Tx) error {
		if err := insertData(ctx, tx, repo.DialectWrapper, "users", daoUser); err != nil {
			return err
		}
		if err := insertData(ctx, tx, repo.DialectWrapper, "security_data", daoSecData); err != nil {
			return err
		}
		if err := insertData(ctx, tx, repo.DialectWrapper, "private_data", daoPrivData); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error(
			"Transaction failed",
			slog.String("err", err.Error()),
		)
		return "", err
	}

	return "", nil
}

func insertData[T any](
	ctx context.Context,
	tx *sql.Tx,
	wrapper *goqu.DialectWrapper,
	table string,
	data *T,
) error {
	stmt, args, err := wrapper.Insert(table).Rows(data).Prepared(true).ToSQL()
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, stmt, args...); err != nil {
		return err
	}

	return nil
}
