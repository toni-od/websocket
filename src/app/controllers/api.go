package controllers

import (
	"fmt"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintf(w, "%s", "Hello")
}
