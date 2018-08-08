package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Member struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var upgrader = websocket.Upgrader{} // use default options

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read: ", err)
			break
		}
		fmt.Println("recv: ", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("write: ", err)
			break
		}
	}
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		member := Member{
			ID:   0,
			Name: "Sean",
		}
		err = conn.WriteJSON(member)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/json", jsonHandler)
	http.ListenAndServe(":8080", nil)
}
