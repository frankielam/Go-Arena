package main

import (
	"fmt"
	"time"
)

func enqueue(c chan string, id int, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	c <- fmt.Sprintf("msg %d", id)
}

func main() {
	c1 := make(chan string, 1)
	go enqueue(c1, 1, 2)
	select {
	case m1 := <-c1:
		fmt.Println(m1)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go enqueue(c2, 2, 2)
	select {
	case m2 := <-c2:
		fmt.Println(m2)
	case <-time.After(3  * time.Second):
		fmt.Println("timeout 2")
	}
	

}
