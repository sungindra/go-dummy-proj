package handler

import (
	"dummy/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouteHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", controller.Index)
	r.Route("/model", func(r chi.Router) {
		r.Get("/", controller.ListModel)
		r.Get("/{id}", controller.GetModel)
	})
	return r
}
