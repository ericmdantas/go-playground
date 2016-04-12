package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const url = "http://jsonplaceholder.typicode.com/posts/1"

func main() {
	transport := http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, time.Duration(1*time.Second))
		},
	}

	client := http.Client{
		Transport: &transport,
	}

	resp, err := client.Get(url)

	var m map[string]interface{}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(b, &m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m)
}
