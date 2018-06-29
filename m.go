package main

import (
	"fmt"
	"sync"
)

func async1(wg *sync.WaitGroup) {
	fmt.Println(1)	
	wg.Done()
}

func async2(wg *sync.WaitGroup) {
	fmt.Println(2)	
	wg.Done()
}

func async3(wg *sync.WaitGroup) {
	fmt.Println(3)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	
	wg.Add(3)
	
	go async1(&wg)
	go async2(&wg)
	go async3(&wg)
	
	fmt.Println("waiting")
	
	wg.Wait()
	
	fmt.Println("done")
}