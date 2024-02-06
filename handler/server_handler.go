package handler

import (
	"database/sql"
	"log"
	"net/http"
)

var DatabaseConnection *sql.DB

func StartServer() {
	var err error
	DatabaseConnection, err = setupDatabase()
	if err != nil {
		panic(err)
	}

	defer DatabaseConnection.Close()

	r := handleRouting()
	log.Fatal(http.ListenAndServe(":8080", r))
}
