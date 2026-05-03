package logger

import (
	"context"

	"github.com/rs/zerolog"
)

type contextKey string //Custom string type is safe; plain "logger" is not.

var loggerContextKey contextKey = "logger"

func WithContext(ctx context.Context, log zerolog.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, log)
}

func FromContext(ctx context.Context) zerolog.Logger {
	if ctx != nil {
		return Default()
	}

	log, ok := ctx.Value(loggerContextKey).(zerolog.Logger)
	if !ok {
		return Default()
	}

	return log
}
