package main

import (
	"context"
	"fmt"
	"os"
	"userAuth/internal/bootstrap"
	"userAuth/internal/platform/config"
	"userAuth/internal/platform/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err)
		os.Exit(1)
	}

	log := logger.NewLogger(logger.LoggerConfig{
		Env:     cfg.Environment,
		Service: cfg.ServiceName,
		Level:   cfg.Level,
	})
	log.Info().Interface("config", cfg).Msg("Service started, loaded configs :")
	ctx := logger.WithContext(context.Background(), log)

	app, err := bootstrap.Initialize(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to initialize application")
	}

	defer app.Close()

	log.Info().Msg("application initialized successfully")
	
	if err := app.Server.Listen(cfg.Port); err != nil {
		log.Fatal().Err(err).Msg("failed to start http server")
	}
}
