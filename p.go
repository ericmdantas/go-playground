package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type mouse struct {
	OffsetX int `json:"offsetX"`
	OffsetY int `json:"offsetY"`
	Amount  int `json:"amount"`
}

const port = ":1234"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(func(c *websocket.Conn) {
		for {
			var m mouse
			websocket.JSON.Receive(c, &m)
			websocket.JSON.Send(c, m)
		}
	}))

	fmt.Println(port)

	http.ListenAndServe(port, nil)
}
