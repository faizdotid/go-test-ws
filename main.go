package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %v", err)
		return
	}
	defer conn.Close()
	for {
		mt, buf, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
		}
		log.Printf("recv: %s", buf)
		err = conn.WriteMessage(mt, buf)
		if err != nil {
			log.Println("write:", err)
			break

		}

	}
}

func main() {

	http.HandleFunc("/ws", wsHandler)
	log.Println("Server started at :9000")
	http.ListenAndServe(":9000", nil)

}
