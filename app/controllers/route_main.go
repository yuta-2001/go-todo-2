package controllers

import (
	"net/http"
	"html/template"
)

func top(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("app/views/templates/top.html"))
	tpl.Execute(w, nil)
}
