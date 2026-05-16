package postgres

import (
	"context"
	"fmt"
	"time"
	"userAuth/internal/platform/config"
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
func NewPool(ctx context.Context, pgConfig config.PostgresConfig) (*pgxpool.Pool, error) {
	log := logger.FromContext(ctx)
	cfg, err := pgxpool.ParseConfig(DSN(pgConfig))
	if err != nil {
		return nil, fmt.Errorf("parse postgres config: %w", err)
	}

	cfg.MaxConns = int32(pgConfig.MaxConns)
	cfg.MinConns = int32(pgConfig.MinConns)
	cfg.MaxConnLifetime = time.Duration(pgConfig.MaxConnLifetimeMinute) * time.Minute
	cfg.MaxConnIdleTime = time.Duration(pgConfig.MaxConnIdleTimeMinute) * time.Minute
	cfg.HealthCheckPeriod = time.Duration(pgConfig.HealthCheckPeriodMinute) * time.Minute
	cfg.ConnConfig.ConnectTimeout = time.Duration(pgConfig.ConnectTimeoutSec) * time.Second

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

func DSN(p config.PostgresConfig) string { //data source name
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.Database,
		p.SSLMode,
	)
}
