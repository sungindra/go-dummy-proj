package handler

import (
	"dummy/controller"
	"dummy/controller/api"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouteHandler() {
	r := chi.NewRouter()

	r.Get("/", controller.Index)

	r.Route("/cms", func(r chi.Router) {
		r.Route("/model", func(r chi.Router) {
			r.Get("/", controller.ListModel)
			r.Get("/{id}", controller.GetModel)
		})
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/model", func(r chi.Router) {
			r.Get("/", api.GetModels)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
