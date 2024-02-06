package handler

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = ""
	dbname   = "dummydb"

	idleTimeInSeconds = 30
	idleConns         = 5
)

func setupDatabase() (*sql.DB, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(idleTimeInSeconds * time.Second)
	db.SetMaxIdleConns(idleConns)

	return db, nil
}
