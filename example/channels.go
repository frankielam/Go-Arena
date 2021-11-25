package main

import "fmt"

func main() {
	messages := make(chan string)
	// messages := make(chan string 2)  // with buffered 
	
	go func() {
		messages <- "hello"
	}()

	msg := <-messages
	fmt.Println(msg)
}
