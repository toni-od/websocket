package routes

import (
	"html/template"
	"net/http"
)

func Register(tmpl *template.Template, staticDir string) {
	http.HandleFunc("GET /", Index(tmpl))
	http.HandleFunc("GET /static/", Static(staticDir))
}
