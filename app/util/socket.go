package util

import (
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

//Struct for messages
type msg struct {
	rep string
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
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

		v := msg{}
		m := v.rep
		err = conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		//fmt.Printf("Got message: %#v\n", m)
		fmt.Println(m)
		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
	}
}
