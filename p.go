package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	amount   = 100
	fileName = "x.json"
	url      = "http://jsonplaceholder.typicode.com/todos/"
)

type todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func log(verb, url string) {
	fmt.Printf("%s -> %s\n", verb, url)
}

func g() <-chan []todo {
	var todos []todo
	var mutex sync.Mutex

	chT := make(chan todo)
	chTs := make(chan []todo, 1)

	for i := 1; i <= amount; i++ {
		go func(j int) {
			var t todo

			u := url + strconv.Itoa(j)

			log("GET", u)

			r, _ := http.Get(u)
			b, _ := ioutil.ReadAll(r.Body)

			defer r.Body.Close()

			json.Unmarshal(b, &t)

			chT <- t
		}(i)

		select {
		case t := <-chT:
			mutex.Lock()
			todos = append(todos, t)
			mutex.Unlock()
		}
	}

	chTs <- todos

	return chTs
}

func p(ts []todo) {
	for _, t := range ts {
		var nT todo

		log("POST", url)

		bT, _ := json.Marshal(t)
		strT := string(bT)

		r, _ := http.Post(url, "application/json", bytes.NewBufferString(strT))
		b, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		json.Unmarshal(b, &nT)

		bNT, _ := json.Marshal(nT)

		ioutil.WriteFile(strconv.Itoa(nT.ID)+".json", bNT, 0644)
	}
}

func main() {
	start := time.Now()

	select {
	case ts := <-g():
		bTs, _ := json.Marshal(ts)
		ioutil.WriteFile(fileName, bTs, 0644)

		p(ts)

		fmt.Println(time.Since(start))
	}
}
