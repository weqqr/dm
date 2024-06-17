package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Connection struct {
	pool *pgxpool.Pool
}

type Config struct {
	DSN string `toml:"dsn"`
}

func Connect(ctx context.Context, config Config) (*Connection, error) {
	poolConfig, err := pgxpool.ParseConfig(config.DSN)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return &Connection{
		pool: pool,
	}, nil
}

func (c *Connection) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return c.pool.Query(ctx, sql, args...)
}

type Database interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}
