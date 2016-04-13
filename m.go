package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const url = "http://jsonplaceholder.typicode.com/posts/1"

func getGet() map[string]interface{} {
	resp, err := http.Get(url)

	var m map[string]interface{}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(b, &m)

	if err != nil {
		fmt.Println(err)
	}

	return m
}

func main() {
	d := make(chan map[string]interface{})

	go func() {
		d <- getGet()
	}()

	select {
	case m := <-d:
		fmt.Println("get returned")
		fmt.Println(m)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	fmt.Println("yo!")
}
