package main 

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)	// close a non-empty channel

	for msg := range queue {
		fmt.Println(msg)
	}
}