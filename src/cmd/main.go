package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"websocket/src/app/controllers"
	"websocket/src/pkg/utils"
)

var (
	port                     int
	tmpl                     *template.Template
	err                      error
	templatePath, staticPath string
)

func init() {
	cwd, err := os.Getwd()
	staticPath = path.Join(cwd, "www")
	templatesPath := path.Join(cwd, "www", "html", "*")
	tmpl, err = template.ParseGlob(templatesPath)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	port = utils.Port()
	address := fmt.Sprintf(":%d", port)

	controllers.Register(tmpl, staticPath)

	fmt.Printf("\n\n\tServer running on http://localhost%s\n\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
