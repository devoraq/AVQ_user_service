package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	GrpcConfig     *GrpcConfig     `env:",init"`
	DatabaseConfig *DatabaseConfig `env:",init"`
}

type GrpcConfig struct {
	Addr    string `env:"GRPC_ADDR"`
	Reflect bool   `env:"REFLECT" envDefault:"true"`
}

type DatabaseConfig struct {
	ConnURL          string        `env:"CONN_URL"`
	MaxOpenConns     int           `env:"MAX_OPEN_CONNS" envDefault:"5"`
	MaxOpenIdleConns int           `env:"MAX_OPEN_IDLE_CONNS" envDefault:"5"`
	ConnMaxIdleTime  time.Duration `env:"CONN_MAX_IDLE_TIME" envDefault:"5m"`
	ConnMaxLifeTime  time.Duration `env:"CONN_MAX_LIFE_TIME" envDefault:"1h"`
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
		panic(err)
	}

	return &cfg
}
