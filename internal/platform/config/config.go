package config

import "github.com/caarlos0/env"

type Config struct {
	ServiceName string `env:"SERVICE_NAME" envDefault:"auth-service"`
	Environment string `env:"ENVIRONMENT" envDefault:"DEV"`
	Postgres    PostgresConfig
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Database string `env:"POSTGRES_DB,required"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`

	MaxConns                int `env:"POSTGRES_MAX_CONNS" envDefault:"25"`
	MinConns                int `env:"POSTGRES_MIN_CONNS" envDefault:"5"`
	MaxConnLifetimeMinute   int `env:"POSTGRES_MAX_CONN_LIFETIME_MIN" envDefault:"30"`
	MaxConnIdleTimeMinute   int `env:"POSTGRES_MAX_CONN_IDLE_TIME_MIN" envDefault:"10"`
	HealthCheckPeriodMinute int `env:"POSTGRES_HEALTH_CHECK_PERIOD_MIN" envDefault:"1"`
	ConnectTimeoutSec       int `env:"POSTGRES_CONNECT_TIMEOUT_SEC" envDefault:"5"`

	MigrationVersion uint `env:"MIGRATION_VERSION" envDefault:"5"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	//parse nested structs
	if err := env.Parse(&cfg.Postgres); err != nil {
		return nil, err
	}
	return cfg, nil
}
