package main

import (
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(func(c *websocket.Conn) {
		c.Write([]byte("yo!"))

		time.Sleep(1 * time.Second)

		for {
			var m string
			websocket.Message.Receive(c, &m)
			c.Write([]byte(m + "!"))
		}
	}))

	http.ListenAndServe(":1234", nil)
}
