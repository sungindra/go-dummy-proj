package handler

import (
	"dummy/database"
	"dummy/repository"
	"log"
	"net/http"
)

func StartServer() {
	var err error
	repository.Repository.DB, err = database.SetupDatabase()
	if err != nil {
		panic(err)
	}

	r := handleRouting()
	log.Fatal(http.ListenAndServe(":8080", r))
}
