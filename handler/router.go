package handler

import (
	"dummy/controller"
	"dummy/controller/api"

	"github.com/go-chi/chi/v5"
)

func handleRouting() *chi.Mux {
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

	return r
}
