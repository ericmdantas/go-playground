package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func received(c *websocket.Conn) <-chan string {
	strCh := make(chan string, 1)

	var m string

	websocket.Message.Receive(c, &m)

	strCh <- m

	return strCh
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(func(c *websocket.Conn) {
		c.Write([]byte("yo!"))

		for {
			select {
			case msg := <-received(c):
				c.Write([]byte(msg + "!"))
			}
		}
	}))

	http.ListenAndServe(":1234", nil)
}
