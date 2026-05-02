package bootstrap

import (
	"context"
	"userAuth/internal/platform/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB *pgxpool.Pool
}

func Initialize(ctx context.Context) (*App, error) {
	db, err := postgres.NewPool(ctx)
	if err != nil {
		return nil, err
	}
	
	return &App{
		DB: db,
	}, nil
}

func (a *App) Close() {
	if a.DB != nil {
		a.DB.Close()
	}
}
