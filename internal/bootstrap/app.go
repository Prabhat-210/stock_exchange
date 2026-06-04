package bootstrap

import (
	"context"
	handler "userAuth/internal/adapters/inbound/http/handlers"
	server "userAuth/internal/adapters/inbound/http/httpserver"
	"userAuth/internal/adapters/inbound/http/routes"
	"userAuth/internal/adapters/outbound/postgress/repository"
	authService "userAuth/internal/core/service/auth"
	"userAuth/internal/platform/config"
	"userAuth/internal/platform/logger"
	"userAuth/internal/platform/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"userAuth/internal/platform/migrations"
)

type App struct {
	DB     *pgxpool.Pool
	Server *fiber.App
}

func Initialize(ctx context.Context, cfg *config.Config) (*App, error) {
	log := logger.FromContext(ctx)

	if err := migrations.Migrate(ctx, postgres.DSN(cfg.Postgres), cfg.Postgres.MigrationVersion); err != nil {
		log.Error().Msg("Failed to run migrate scripts")
		return nil, err
	}
	log.Info().Msgf("Migrations scripts run successfully with version: %d", cfg.Postgres.MigrationVersion)

	db, err := postgres.NewPool(ctx, cfg.Postgres)
	if err != nil {
		log.Error().Msg("Failed to intialize postgres")
		return nil, err
	}
	log.Info().Msg("Postgres initialized successfully")

	//todo:repository
	userRepo := repository.NewUserRepository(db)
	authSvc := authService.NewAuthService(log, userRepo)

	authHandle := handler.NewAuthHandler(authSvc, log)

	server := server.NewServer()
	routes.RegisterAuthRoutes(server, authHandle)

	log.Info().Str("Application PORT", cfg.Port).Msg("Application wired successfully, starting server")
	return &App{
		DB:     db,
		Server: server,
	}, nil
}

func (a *App) Close() {
	if a.DB != nil {
		a.DB.Close()
	}
}
