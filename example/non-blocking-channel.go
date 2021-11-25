package main

import "fmt"

func main() {
	msg := make(chan string)

	select {
	case m := <-msg : 
		fmt.Println("received", m)
	default:
		fmt.Println("no data received")
	}

	select {
	case msg <- "hello": 
		fmt.Println("data sent")
	default: 
		fmt.Println("no data sent")
	}


}
