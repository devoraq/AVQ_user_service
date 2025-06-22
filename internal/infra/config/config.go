package config

import (
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	GrpcConfig *GrpcConfig `env:",init"`
}

type GrpcConfig struct {
	Addr    string `env:"GRPC_ADDR"`
	Reflect bool   `env:"REFLECT" envDefault:"true"`
}

func NewConfig(
	log *slog.Logger,
	path string,
) *Config {
	const op = "config.NewConfig"

	log = log.With("op", op)

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			log.Error(
				"File not found",
				slog.String("path", path),
			)
		} else {
			log.Error(
				"Error checking file",
				slog.String("err", err.Error()),
			)
		}
		panic(err)
	}

	if err := godotenv.Load(path); err != nil {
		log.Error(
			"Error reading file",
			slog.String("err", err.Error()),
		)
		panic(err)
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Error(
			"Error parsing variables into structure",
			slog.String("err", err.Error()),
		)
	}

	return &cfg
}
