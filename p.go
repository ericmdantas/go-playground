package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type info struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Amount int    `json:"amount"`
	Hash   string `json:"hash"`
}

const port = ":1234"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(func(c *websocket.Conn) {
		for {
			var i info

			h := md5.New()

			websocket.JSON.Receive(c, &i)

			bI, _ := json.Marshal(i)

			h.Write(bI)
			i.Hash = hex.EncodeToString(h.Sum(nil))

			websocket.JSON.Send(c, i)
		}
	}))

	fmt.Println(port)

	http.ListenAndServe(port, nil)
}
