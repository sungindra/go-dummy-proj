package handler

import (
	"dummy/controller"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouteHandler() {
	r := chi.NewRouter()
	r.Get("/", controller.Index)
	r.Route("/model", func(r chi.Router) {
		r.Get("/", controller.ListModel)
		r.Get("/{id}", controller.GetModel)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
