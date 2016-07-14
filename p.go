package main

import (
	"net/http"

	ws "golang.org/x/net/websocket"
)

type msg struct {
	Txt string `json:"txt"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.Handle("/ws", ws.Handler(func(client *ws.Conn) {
		var m msg

		if err := ws.JSON.Receive(client, &m); err != nil {
			panic(err)
		}

		ws.JSON.Send(client, msg{Txt: m.Txt + "!"})
	}))

	http.ListenAndServe(":3456", nil)
}
