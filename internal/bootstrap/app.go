package bootstrap

import (
	"context"
	"userAuth/internal/platform/logger"
	"userAuth/internal/platform/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB *pgxpool.Pool
}

func Initialize(ctx context.Context) (*App, error) {
	log := logger.FromContext(ctx)
	db, err := postgres.NewPool(ctx)
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
