package main

import (
	"context"
	"userAuth/internal/bootstrap"
	"userAuth/internal/platform/logger"
)

func main() {

	log := logger.NewLogger(logger.LoggerConfig{

		Service: "auth-service",
		Level:   "info",
	})

	ctx := logger.WithContext(context.Background(), log)

	log.Info().Msg("starting auth service..")

	app, err := bootstrap.Initialize(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to initialize application")
	}
	defer app.Close()

	log.Info().Msg("application initialized successfully")
}
