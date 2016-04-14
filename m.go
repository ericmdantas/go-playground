package main

import (
	"fmt"
	"time"
)

type ball struct {
	hits int
}

func play(name string, start chan *ball) {
	for {
		b := <-start

		b.hits++
		fmt.Println(name, b.hits)

		time.Sleep(time.Millisecond * 100)

		start <- b
	}
}

func main() {
	s := make(chan *ball)

	go play("a", s)
	go play("b", s)

	s <- new(ball)

	time.Sleep(1 * time.Second)

	<-s
}
