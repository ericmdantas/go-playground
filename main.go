package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		time.Sleep(1001 * time.Millisecond)
		fmt.Println("1")
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		time.Sleep(1002 * time.Millisecond)
		fmt.Println("2")
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		time.Sleep(1003 * time.Millisecond)
		fmt.Println("3")
		wg.Done()
	}(&wg)

	wg.Wait()

	fmt.Println("done")
}
