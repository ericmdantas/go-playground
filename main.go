package main

import (
	"crypto/md5"
	"fmt"
	"sync"
	"time"
)

type user struct {
	name string
	age  uint
}

func (u user) hashIt() string {
	nu := []byte(fmt.Sprintf("%v", u))
	return fmt.Sprintf("%x", md5.Sum(nu))
}

func launch(hash string, u user, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Second)
	fmt.Printf("Launching %v [%s]\n", u, hash)
}

func main() {
	var wg sync.WaitGroup
	users := make(map[string]user)

	u0 := user{name: "x0", age: 10}
	u1 := user{name: "x1", age: 18}
	u2 := user{name: "x2", age: 19}
	u3 := user{name: "x3", age: 20}
	u4 := user{name: "x4", age: 21}
	u5 := user{name: "x5", age: 22}

	users[u0.hashIt()] = u0
	users[u1.hashIt()] = u1
	users[u2.hashIt()] = u2
	users[u3.hashIt()] = u3
	users[u4.hashIt()] = u4
	users[u5.hashIt()] = u5

	wg.Add(len(users))

	for k, v := range users {
		go launch(k, v, &wg)
	}

	wg.Wait()
}
