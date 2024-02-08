package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func SetupDatabase() (*sql.DB, error) {
	cfg, err := loadEnv()
	if err != nil {
		return nil, err
	}

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.postgresHost,
		cfg.postgresPort,
		cfg.postgresUser,
		cfg.postgresPassword,
		cfg.postgresDatabaseName,
	)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Duration(cfg.maxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.maxIdleTime) * time.Second)
	db.SetMaxOpenConns(cfg.maxOpenConns)
	db.SetMaxIdleConns(cfg.maxIdleConns)

	return db, nil
}
