package controllers

import (
	"html/template"
	"net/http"
)

func Register(tmpl *template.Template, staticDir string) {
	http.HandleFunc("/static/", Static(staticDir))
	
	http.HandleFunc("/", Index(tmpl))
	http.HandleFunc("/ws", ServeWS)
}
