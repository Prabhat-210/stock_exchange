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
		Service: cfg.Environment,
		Level:   cfg.Level,
	})

	log.Info().Interface("config", cfg).Msg("Service started, loaded configs :")

	ctx := logger.WithContext(context.Background(), log)

	app, err := bootstrap.Initialize(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to initialize application")
		os.Exit(1)
	}

	defer app.Close()

	log.Info().Msg("application initialized successfully")
}
