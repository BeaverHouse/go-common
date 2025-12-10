package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/BeaverHouse/go-common/env"
	"github.com/jackc/pgx/v5/pgxpool"
)

// InitFromEnv initializes a PostgreSQL connection pool from environment variables.
// Panics if connection fails.
func InitFromEnv() *pgxpool.Pool {
	postgresConfig := PostgresConfig{
		Host:     env.GetEnv("POSTGRES_HOST", "localhost"),
		Port:     env.GetIntEnv("POSTGRES_PORT", 5432),
		User:     env.GetEnv("POSTGRES_USER", "postgres"),
		Password: env.GetEnv("POSTGRES_PASSWORD", "postgres"),
		DBName:   env.GetEnv("POSTGRES_DB", "postgres"),
		SSLMode:  env.GetEnv("POSTGRES_SSLMODE", "disable"),
	}

	pool, err := newPool(postgresConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return pool
}

func newPool(cfg PostgresConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
