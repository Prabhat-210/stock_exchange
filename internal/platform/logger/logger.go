package logger

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var defaultLogger zerolog.Logger

type LoggerConfig struct {
	Service string
	Env     string
	Level   string
}

func NewLogger(cfg LoggerConfig) zerolog.Logger {
	writer := getWriter(cfg.Env)

	log := zerolog.New(writer).
		Level(parseLevel(cfg.Level)).
		With().
		Timestamp().
		Caller().
		Str("service", cfg.Service).
		Str("env", cfg.Env).
		Logger()

	defaultLogger = log
	return log
}

func Default() zerolog.Logger {
	return defaultLogger //will use when no request context exists.
}

func getWriter(env string) io.Writer {
	switch strings.ToLower(env) {

	case "local", "dev":
		return zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	default:
		return os.Stdout //no need to set TimeFormat In prod, logs are JSON and zerolog handles timestamp automatically.
	}
}

func parseLevel(level string) zerolog.Level {

	switch strings.ToLower(level) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}

}
