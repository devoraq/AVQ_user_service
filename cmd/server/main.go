package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/DENFNC/awq_user_service/internal/app"
	"github.com/DENFNC/awq_user_service/internal/infra/config"
	"github.com/DENFNC/awq_user_service/internal/infra/db/postgres"
)

func main() {
	logger := initLogger()

	cfg := config.NewConfig(logger, "./.env.example")

	db := postgres.NewDatabase(
		logger,
		postgres.WithMaxOpenConns(cfg.DatabaseConfig.MaxOpenConns),
		postgres.WithMaxIdleConns(cfg.DatabaseConfig.MaxOpenIdleConns),
	)

	db.OpenDB(cfg.DatabaseConfig.ConnURL)

	application := app.New(logger, cfg)
	go application.App.Start()

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh

	logger.Info(
		"Calling program termination",
		slog.String("signal", sig.String()),
	)

	application.App.Stop()
}

func initLogger() *slog.Logger {
	logger := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{},
		),
	)
	slog.SetDefault(logger)

	return logger
}
