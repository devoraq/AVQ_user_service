package postgres

import (
	"database/sql"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

type Option func(*databaseOption)

type Database struct {
	*slog.Logger
	*databaseOption
}

func NewDatabase(
	log *slog.Logger,
	opts ...Option,
) *Database {
	var dbOption databaseOption
	for _, opt := range opts {
		opt(&dbOption)
	}

	return &Database{
		Logger:         log,
		databaseOption: &dbOption,
	}
}

type databaseOption struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

func WithMaxOpenConns(maxConns int) Option {
	return func(do *databaseOption) {
		do.MaxOpenConns = maxConns
	}
}

func WithMaxIdleConns(maxIdleConns int) Option {
	return func(do *databaseOption) {
		do.MaxIdleConns = maxIdleConns
	}
}

func WithConnMaxIdleTime(maxIdleTime time.Duration) Option {
	return func(do *databaseOption) {
		do.ConnMaxIdleTime = maxIdleTime
	}
}

func WithConnMaxLifetime(maxLifeTime time.Duration) Option {
	return func(do *databaseOption) {
		do.ConnMaxLifetime = maxLifeTime
	}
}

func (db *Database) OpenDB(
	dataSourceName string,
	opts ...Option,
) (*sql.DB, error) {
	const op = "postgres.Database.OpenDB"

	log := db.Logger.With("op", op)

	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Error(
			"The database connection could not be opened",
			slog.String("err", err.Error()),
		)
		return nil, err
	}

	database.SetMaxOpenConns(db.MaxOpenConns)
	database.SetMaxIdleConns(db.MaxIdleConns)
	database.SetConnMaxIdleTime(db.ConnMaxIdleTime)
	database.SetConnMaxLifetime(db.ConnMaxLifetime)

	if err := database.Ping(); err != nil {
		log.Error(
			"The database is unavailable",
			slog.String("err", err.Error()),
		)
		return nil, err
	}

	return database, nil
}
