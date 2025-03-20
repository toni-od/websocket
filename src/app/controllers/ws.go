package controllers

import (
	"fmt"
	"net/http"

	ws "websocket/src/app/websocket"
)

func ServeWS(w http.ResponseWriter, r *http.Request) {
	wsManager := ws.NewWSManager()
	fmt.Println(wsManager)

	fmt.Fprintf(w, "%s", "Hello")
}
