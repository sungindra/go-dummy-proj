package controller

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("view/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}
