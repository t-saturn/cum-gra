package handlers

import (
	"html/template"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/login.html",
	))

	tmpl.ExecuteTemplate(w, "base", nil)
}
