package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

const port = ":1234"

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

	fmt.Println(port)

	http.ListenAndServe(port, nil)
}
