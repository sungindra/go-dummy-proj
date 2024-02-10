package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func SetupDatabase() (*sql.DB, error) {
	cfg, err := loadEnv()
	if err != nil {
		return nil, err
	}

	log.Printf("cfg = %#v", cfg)

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=public",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabaseName,
	)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Printf("database open connection error: " + err.Error())
		return nil, err
	}

	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleTime) * time.Second)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return db, nil
}
