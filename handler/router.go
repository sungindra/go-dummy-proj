package handler

import (
	"dummy/controller"
	"dummy/controller/api"

	"github.com/go-chi/chi/v5"
)

func HandleRouting(
	homeController controller.Home,
	modelController controller.Model,
	apiController api.API,
) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", homeController.Index)

	r.Route("/cms", func(r chi.Router) {
		r.Route("/model", func(r chi.Router) {
			r.Get("/", modelController.ListModel)
			r.Get("/{id}", modelController.GetModel)
		})
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/model", func(r chi.Router) {
			r.Get("/", apiController.GetModels)
		})
	})

	return r
}
