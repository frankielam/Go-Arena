package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	messages <- "1"
	messages <- "2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
