package controller

import (
	"dummy/repository"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var listViewTmpl = template.Must(template.ParseFiles("view/model/listview.html"))

func ListModel(w http.ResponseWriter, r *http.Request) {
	models, err := repository.GetModels()

	if err != nil {
		log.Fatal(err)
	}

	listViewTmpl.Execute(w, models)
}

func GetModel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	model, err := repository.GetModel(id)

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("Model : " + model.Attribute1))
}
