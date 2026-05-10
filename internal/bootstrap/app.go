package bootstrap

import (
	"context"
	"fmt"
	"userAuth/internal/platform/config"
	"userAuth/internal/platform/logger"
	"userAuth/internal/platform/postgres"

	"github.com/jackc/pgx/v5/pgxpool"

	"userAuth/internal/platform/migrations"
)

type App struct {
	DB     *pgxpool.Pool
}

func Initialize(ctx context.Context, cfg *config.Config) (*App, error) {
	log := logger.FromContext(ctx)

	if err := migrations.Migrate(postgres.DSN(cfg.Postgres), cfg.Postgres.MigrationVersion); err != nil {
		log.Error().Msg("Failed to run migrate scripts")
		return nil, err
	}
	log.Info().Msg(fmt.Sprintf("Migrations scripts run successfully with version: %d", cfg.Postgres.MigrationVersion))

	db, err := postgres.NewPool(ctx, cfg.Postgres)
	if err != nil {
		log.Error().Msg("Failed to intialize postgres")
		return nil, err
	}

	log.Info().Msg("postgres initialize")
	return &App{
		DB: db,
	}, nil
}

func (a *App) Close() {
	if a.DB != nil {
		a.DB.Close()
	}
}
