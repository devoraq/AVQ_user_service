package app

import (
	"database/sql"
	"log/slog"

	user "github.com/DENFNC/awq_user_service/internal/adapters/grpc/v1"
	grpcapp "github.com/DENFNC/awq_user_service/internal/app/grpc"
	service "github.com/DENFNC/awq_user_service/internal/core/services/v1"
	"github.com/DENFNC/awq_user_service/internal/infra/config"
	"github.com/DENFNC/awq_user_service/internal/infra/postgres"
	"github.com/DENFNC/awq_user_service/internal/infra/postgres/repository"
)

type App struct {
	App *grpcapp.App
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	db := initDatabase(log, cfg)

	userRepo := repository.NewUserRepository(log, db)
	userSrv := service.NewUserService(log, userRepo)
	userHandle := user.NewUser(userSrv)

	return &App{
		App: grpcapp.NewApp(
			log,
			cfg.GrpcConfig.Addr,
			cfg.GrpcConfig.Reflect,
			userHandle,
		),
	}
}

func initDatabase(
	log *slog.Logger,
	cfg *config.Config,
) *sql.DB {
	db := postgres.NewDatabase(
		log,
		postgres.WithMaxOpenConns(cfg.DatabaseConfig.MaxOpenConns),
		postgres.WithMaxIdleConns(cfg.DatabaseConfig.MaxOpenIdleConns),
	)

	psql, err := db.OpenDB(cfg.DatabaseConfig.ConnURL)
	if err != nil {
		panic(err)
	}

	return psql
}
