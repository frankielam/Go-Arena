package main

import (
	"fmt"
)

func ping(q chan<- string) {
	q <- "hello"
}

func pong(q <-chan string) {
	m := <-q
	fmt.Println(m)
}

func main() {
	q := make(chan string)
	go ping(q)
	pong(q)
}
