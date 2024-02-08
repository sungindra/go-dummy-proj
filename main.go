package main

import (
	"dummy/controller"
	"dummy/controller/api"
	"dummy/database"
	"dummy/handler"
	"dummy/repository"
	"log"
	"net/http"
)

func main() {
	var err error
	db, err := database.SetupDatabase()
	if err != nil {
		panic(err)
	}
	repo := repository.New(db)

	// create controllers
	homeController := controller.NewHome()
	modelController := controller.NewModel(repo)
	apiController := api.NewAPI(repo)

	// route and start server
	r := handler.HandleRouting(homeController, modelController, apiController)
	log.Fatal(http.ListenAndServe(":8080", r))
}
