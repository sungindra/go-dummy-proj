package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dummyuser"
	password = "12345678"
	dbname   = "dummydb"

	maxLifeTime  = 30 * time.Second
	maxIdleTime  = 5 * time.Second
	maxOpenConns = 5
	maxIdleConns = 1
)

func SetupDatabase() (*sql.DB, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(maxLifeTime)
	db.SetConnMaxIdleTime(maxIdleTime)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	return db, nil
}
