package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct{}

func NewWSManager() *Manager {
	return &Manager{}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Println("Websocket connection")

	// upgrade http conncetion to ws
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
}
