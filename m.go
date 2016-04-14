package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   int    `json:"body"`
}

const (
	url1 = "http://jsonplaceholder.typicode.com/posts/1"
	url2 = "http://jsonplaceholder.typicode.com/posts/2"
	url3 = "http://jsonplaceholder.typicode.com/posts/3"
)

func goGetIt(ps chan post, url string) {
	var p post

	pl, _ := http.Get(url)
	b, _ := ioutil.ReadAll(pl.Body)

	defer pl.Body.Close()

	json.Unmarshal(b, &p)

	ps <- p
}

func logger(pc1, pc2, pc3 chan post) {
	select {
	case <-pc1:
		fmt.Println(".")
	case <-pc2:
		fmt.Println("..")
	case <-pc3:
		fmt.Println("...")
	}
}

func main() {
	pc1 := make(chan post)
	pc2 := make(chan post)
	pc3 := make(chan post)

	go goGetIt(pc1, url1)
	go goGetIt(pc2, url2)
	go goGetIt(pc3, url3)

	logger(pc1, pc2, pc3)
}
