package database

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type config struct {
	//
	postgresHost         string
	postgresPort         int
	postgresUser         string
	postgresPassword     string
	postgresDatabaseName string
	//
	maxLifeTime  int
	maxIdleTime  int
	maxOpenConns int
	maxIdleConns int
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
