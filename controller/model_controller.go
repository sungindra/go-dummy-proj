package controller

import (
	"dummy/repository"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var listViewTmpl = template.Must(template.ParseFiles("view/model/listview.html"))
var errorTmpl = template.Must(template.ParseFiles("view/error.html"))

type Model interface {
	ListModel(w http.ResponseWriter, r *http.Request)
	GetModel(w http.ResponseWriter, r *http.Request)
}

type modelController struct {
	repo *repository.Repository
}

func NewModel(repo *repository.Repository) Model {
	return &modelController{
		repo: repo,
	}
}

func (c *modelController) ListModel(w http.ResponseWriter, r *http.Request) {
	models, err := c.repo.GetModels()
	if err != nil {
		log.Print("error when list model: " + err.Error())
		errorTmpl.Execute(w, nil)
		return
	}

	listViewTmpl.Execute(w, models)
}

func (c *modelController) GetModel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	model, err := c.repo.GetModel(id)
	if err != nil {
		log.Print("error when get model: " + err.Error())
		errorTmpl.Execute(w, nil)
		return
	}

	w.Write([]byte("Model : " + model.Attribute1))
}
