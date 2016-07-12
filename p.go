package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	amount = 5000
	url    = "http://jsonplaceholder.typicode.com/photos/"
)

type photo struct {
	AlbumID     int    `json:"albumId"`
	ID          int    `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	ThumbailURL string `json:"thumbnailUrl"`
}

func main() {
	var photos []photo
	var mutex sync.Mutex

	start := time.Now()
	chP := make(chan photo)

	for i := 1; i <= amount; i++ {
		go func(j int) {
			var p photo

			u := url + strconv.Itoa(j)

			fmt.Printf("-> %s\n", u)

			r, _ := http.Get(u)
			b, _ := ioutil.ReadAll(r.Body)

			json.Unmarshal(b, &p)

			r.Body.Close()

			chP <- p
		}(i)

		select {
		case p := <-chP:
			mutex.Lock()
			photos = append(photos, p)
			mutex.Unlock()
		}
	}

	bTs, _ := json.Marshal(photos)
	ioutil.WriteFile("p.json", bTs, 0644)

	fmt.Println(time.Since(start))
}
