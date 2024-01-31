package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to upgrade to WebSocket")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade errors:", err)
		return
	}
	log.Println("WebSocket connection established")
	defer conn.Close()

	go periodicSend(conn)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Received: %s\n", message)

		// Echo the received message back to the client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func periodicSend(conn *websocket.Conn) {
	log.Println("Starting periodic send")
	for {
		time.Sleep(5 * time.Second)
		message := "Hello from server: " + time.Now().Format(time.RFC3339)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
