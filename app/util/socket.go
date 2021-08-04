package util

import (
	"app/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//Varibales for websockets
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected sucessfully...")
	reader(ws)
}

func reader(conn *websocket.Conn) {
	v := models.Messages{}
	for {
		conn.ReadJSON(&v)
		models.CreateMsg(&v)
		fmt.Println(v)
	}
}
