package postgres

import (
	"context"
	"fmt"
	"time"
	"userAuth/internal/platform/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

/*
----------------------NOTE-------------------------------------------
Instead of opening a new DB connection for every query, pgxpool.Pool:
> opens a set of connections
> keeps them alive
> reuses them across requests
So your app talks to the pool.
---------------------------------------------------------------------
*/
func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	log := logger.FromContext(ctx)
	cfg, err := pgxpool.ParseConfig(DSN())
	if err != nil {
		return nil, fmt.Errorf("parse postgres config: %w", err)
	}

	//TODO: need to set
	cfg.MaxConns = 25
	cfg.MinConns = 5
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.MaxConnIdleTime = 10 * time.Minute
	cfg.HealthCheckPeriod = 1 * time.Minute
	cfg.ConnConfig.ConnectTimeout = 5 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Error().Msg("Failed to create postgres pool")
		return nil, fmt.Errorf("unable to connect with pool:%w ", err)
	}

	log.Info().Msg("postgres pool created")

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	log.Info().Msg("postgres ping successful")
	return pool, nil
}

func DSN() string { //data source name
	return "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
}
