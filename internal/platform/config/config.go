package config

import "github.com/caarlos0/env/v11"

type Config struct {
	ServiceName string         `env:"SERVICE_NAME" envDefault:"auth-service"`
	Port        string         `env:"PORT" envDefault:":8080"`
	Environment string         `env:"ENVIRONMENT" envDefault:"DEV"`
	Level       string         `env:"LEVEL" envDefault:"info"`
	Postgres    PostgresConfig `envPrefix:"POSTGRES_"`
}

type PostgresConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	Database string `env:"DB,required"`
	SSLMode  string `env:"SSL_MODE" envDefault:"disable"`

	MaxConns                int `env:"MAX_CONNS" envDefault:"25"`
	MinConns                int `env:"MIN_CONNS" envDefault:"5"`
	MaxConnLifetimeMinute   int `env:"MAX_CONN_LIFETIME_MIN" envDefault:"30"`
	MaxConnIdleTimeMinute   int `env:"MAX_CONN_IDLE_TIME_MIN" envDefault:"10"`
	HealthCheckPeriodMinute int `env:"HEALTH_CHECK_PERIOD_MIN" envDefault:"1"`
	ConnectTimeoutSec       int `env:"CONNECT_TIMEOUT_SEC" envDefault:"5"`

	MigrationVersion uint `env:"MIGRATION_VERSION" envDefault:"5"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
