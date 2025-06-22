package app

import (
	"log/slog"

	grpcapp "github.com/DENFNC/awq_user_service/internal/app/grpc"
	"github.com/DENFNC/awq_user_service/internal/infra/config"
)

type App struct {
	App *grpcapp.App
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	return &App{
		App: grpcapp.NewApp(
			log,
			cfg.GrpcConfig.Addr,
			cfg.GrpcConfig.Reflect,
		),
	}
}
