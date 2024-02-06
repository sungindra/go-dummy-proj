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
)

var DatabaseConnection *sql.DB

func setupDatabase() {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DatabaseConnection, err = sql.Open("postgres", psqlConn)
	if err != nil {
		panic(err)
	}

	DatabaseConnection.SetConnMaxIdleTime(30 * time.Second)
	DatabaseConnection.SetMaxIdleConns(5)

}
