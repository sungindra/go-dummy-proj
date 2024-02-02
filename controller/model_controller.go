package controller

import (
	"dummy/repository"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ListModel(w http.ResponseWriter, r *http.Request) {
	models := repository.GetModels()
	tmpl := template.Must(template.ParseFiles("view/model/listview.html"))
	tmpl.Execute(w, models)
}

func GetModel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	model := repository.GetModel(id)
	w.Write([]byte("Model : " + model.Attribute1))
}