package main

import "fmt"

func s(cb func(s string)) string {
	cb("yo")

	return "yo"
}

func main() {
	s(func(info string) {
		fmt.Println(info)
	})
}
