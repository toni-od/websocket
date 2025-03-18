package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"websocket/src/routes"
)

const (
	STATIC = "www"
)

var (
	port = 9000
	tmpl *template.Template
	err  error
)

func init() {
	tmpl, err = template.ParseGlob("./www/html/*")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	address := fmt.Sprintf(":%d", port)

	routes.Register(tmpl, STATIC)

	fmt.Printf("\n\n\tServer running on %s\n\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
