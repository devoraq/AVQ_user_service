package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/DENFNC/awq_user_service/internal/core/domain"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type UserDAO interface {
	Insert(
		ctx context.Context,
		user *domain.User,
	) (string, error)
}

type UserRepository struct {
	*slog.Logger
	*sql.DB
	*goqu.DialectWrapper
	UserDAO
}

func NewUserRepository(
	db *sql.DB,
	userDAO UserDAO,
) *UserRepository {
	dialect := goqu.Dialect("postgres")

	return &UserRepository{
		DB:             db,
		DialectWrapper: &dialect,
		UserDAO:        userDAO,
	}
}

func (repo *UserRepository) Save(
	ctx context.Context,
	agg *domain.UserAggregate,
) (string, error) {
	const op = "repository.UserRepository.CreateUser"

	log := repo.Logger.With("op", op)

	if _, err := repo.UserDAO.Insert(ctx, agg.User); err != nil {
		return "", err
	}

	return "", nil
}
