package controller

import (
	"html/template"
	"net/http"
)

var indexTmpl = template.Must(template.ParseFiles("view/index.html"))

type Home interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type homeController struct{}

func NewHome() Home {
	return &homeController{}
}

func (c *homeController) Index(w http.ResponseWriter, r *http.Request) {
	indexTmpl.Execute(w, nil)
}
