package main

import (
	"fmt"
	"time"
)

func async1(ch chan uint32) {
	time.Sleep(1 * time.Millisecond)
	ch <- 1
}

func async2(ch chan uint32) {
	time.Sleep(2 * time.Millisecond)
	ch <- 2
}

func async3(ch chan uint32) {
	ch <- 3
}

func main() {
	ch1 := make(chan uint32)
	ch2 := make(chan uint32)
	ch3 := make(chan uint32)
	
	go async1(ch1)
	go async2(ch2)
	go async3(ch3)
	
	fmt.Println("waiting")
	
	select {
		case x := <- ch1:
			fmt.Println(x)
		case x := <- ch2:
			fmt.Println(x)
		case x := <- ch3:
			fmt.Println(x)
	}
	
	fmt.Println("done")
}