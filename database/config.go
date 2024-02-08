package database

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type config struct {
	//
	postgresHost         string `env:"POSTGRES_HOST"`
	postgresPort         int    `env:"POSTGRES_PORT" envDefault:"5432"`
	postgresUser         string `env:"POSTGRESS_USER"`
	postgresPassword     string `env:"POSTGRESS_PASSWORD"`
	postgresDatabaseName string `env:"POSTGRES_DB"`
	//
	maxLifeTime  int `env:"DATABASE_MAX_LIFE_TIME"`
	maxIdleTime  int `env:"DATABASE_MAX_IDLE_TIME"`
	maxOpenConns int `env:"DATABASE_MAX_OPEN_CONNECTIONS"`
	maxIdleConns int `env:"DATABASE_MAX_IDLE_CONNECTIONS"`
}

func loadEnv() (*config, error) {
	// ignore error in case if `.env` does not exist.
	// In that case, no env var will be loaded.
	_ = godotenv.Load()

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
