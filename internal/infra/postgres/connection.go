package postgres

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
)

type Option func(*databaseOption)

type Database struct {
	log     *slog.Logger
	Dialect *goqu.DialectWrapper
	option  *databaseOption
}

func NewDatabase(
	log *slog.Logger,
	opts ...Option,
) *Database {
	var dbOption databaseOption
	for _, opt := range opts {
		opt(&dbOption)
	}

	dialect := goqu.Dialect("postgres")

	return &Database{
		log:     log,
		Dialect: &dialect,
		option:  &dbOption,
	}
}

type databaseOption struct {
	maxOpenConns    int
	maxIdleConns    int
	connMaxIdleTime time.Duration
	connMaxLifetime time.Duration
}

func WithMaxOpenConns(maxConns int) Option {
	return func(do *databaseOption) {
		do.maxOpenConns = maxConns
	}
}

func WithMaxIdleConns(maxIdleConns int) Option {
	return func(do *databaseOption) {
		do.maxIdleConns = maxIdleConns
	}
}

func WithConnMaxIdleTime(maxIdleTime time.Duration) Option {
	return func(do *databaseOption) {
		do.connMaxIdleTime = maxIdleTime
	}
}

func WithConnMaxLifetime(maxLifeTime time.Duration) Option {
	return func(do *databaseOption) {
		do.connMaxLifetime = maxLifeTime
	}
}

func (db *Database) ReceiveDialect() *goqu.DialectWrapper {
	return db.Dialect
}

func (db *Database) OpenDB(
	dataSourceName string,
	opts ...Option,
) (*sql.DB, error) {
	const op = "postgres.Database.OpenDB"

	log := db.log.With("op", op)

	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Error(
			"The database connection could not be opened",
			slog.String("err", err.Error()),
		)
		return nil, err
	}

	database.SetMaxOpenConns(db.option.maxOpenConns)
	database.SetMaxIdleConns(db.option.maxIdleConns)
	database.SetConnMaxIdleTime(db.option.connMaxIdleTime)
	database.SetConnMaxLifetime(db.option.connMaxLifetime)

	if err := database.Ping(); err != nil {
		log.Error(
			"The database is unavailable",
			slog.String("err", err.Error()),
		)
		return nil, err
	}

	return database, nil
}
