package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

	start := time.Now()

	for i := 1; i <= amount; i++ {
		var p photo

		u := url + strconv.Itoa(i)

		fmt.Printf("-> %s\n", u)

		r, _ := http.Get(u)
		b, _ := ioutil.ReadAll(r.Body)

		json.Unmarshal(b, &p)

		photos = append(photos, p)

		r.Body.Close()
	}

	fmt.Println(time.Since(start))
}
