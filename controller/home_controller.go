package controller

import (
	"html/template"
	"net/http"
)

var indexTmpl = template.Must(template.ParseFiles("view/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	indexTmpl.Execute(w, nil)
}
