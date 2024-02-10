package database

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type config struct {
	//
	PostgresHost         string `env:"POSTGRES_HOST"`
	PostgresPort         int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser         string `env:"POSTGRES_USER"`
	PostgresPassword     string `env:"POSTGRES_PASSWORD"`
	PostgresDatabaseName string `env:"POSTGRES_DB"`
	//
	MaxLifeTime  int `env:"DATABASE_MAX_LIFE_TIME"`
	MaxIdleTime  int `env:"DATABASE_MAX_IDLE_TIME"`
	MaxOpenConns int `env:"DATABASE_MAX_OPEN_CONNECTIONS"`
	MaxIdleConns int `env:"DATABASE_MAX_IDLE_CONNECTIONS"`
}

func loadEnv() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("error when parsing env variables: %s", err.Error())
		return nil, err
	}

	return &cfg, nil
}
