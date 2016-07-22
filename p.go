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

func doEet(c *websocket.Conn) <-chan info {
	ch := make(chan info, 1)

	var i info

	h := md5.New()

	websocket.JSON.Receive(c, &i)

	bI, _ := json.Marshal(i)

	h.Write(bI)
	i.Hash = hex.EncodeToString(h.Sum(nil))

	ch <- i

	return ch
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(func(c *websocket.Conn) {
		for {
			select {
			case i := <-doEet(c):
				websocket.JSON.Send(c, i)
			}
		}
	}))

	fmt.Println(port)

	http.ListenAndServe(port, nil)
}
