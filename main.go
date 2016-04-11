package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("yo!")
	case <-time.After(time.Nanosecond):
		fmt.Println("not yo!")
	}
}
