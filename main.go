package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"websocket/src/controllers"
	"websocket/src/pkg/utils"
)

const (
	STATIC = "www"
)

var (
	port int
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
	port = utils.Port()
	address := fmt.Sprintf(":%d", port)

	controllers.Register(tmpl, STATIC)

	fmt.Printf("\n\n\tServer running on %s\n\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
