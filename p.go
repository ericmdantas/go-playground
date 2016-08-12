package main

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func main() {
	c := 0
	ticker := time.NewTicker(1 * time.Millisecond)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(func(conn *websocket.Conn) {
		var p pos

		for {
			select {
			case <-ticker.C:
				c = c + 1

				p.X = c
				p.Y = c

				bP, _ := json.Marshal(p)

				conn.Write(bP)
			}
		}
	}))

	http.ListenAndServe(":7777", nil)
}
