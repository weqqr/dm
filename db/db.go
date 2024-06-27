package db

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// Используется для регистрации драйвера.
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// Используется для инициализации pgx драйвера.
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Connection представляет подключение к базе данных через пул соединений pgxpool.
type Connection struct {
	// pool представляет пул соединений к базе данных.
	pool *pgxpool.Pool
}

// Config содержит конфигурацию подключения к базе данных.
type Config struct {
	// DSN строка подключения к базе данных.
	DSN string `toml:"dsn"`
}

// Connect устанавливает соединение с базой данных и возвращает Connection.
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

// Migrate выполняет миграции базы данных на основе конфигурации.
func Migrate(config Config) error {
	db, err := sql.Open("pgx", config.DSN)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		"postgres",
		driver,
	)

	if err != nil {
		return err
	}

	if err = mig.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

// Query выполняет SQL-запрос и возвращает результат.
func (c *Connection) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return c.pool.Query(ctx, sql, args...)
}

// QueryRow выполняет SQL-запрос и возвращает одну строку результата.
func (c *Connection) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return c.pool.QueryRow(ctx, sql, args...)
}

// Database представляет интерфейс для работы с базой данных.
type Database interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}
