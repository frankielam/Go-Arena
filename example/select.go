package main

import "fmt"
import "time"

// https://gobyexample.com/select
func main() {
	
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func(c chan<- string) {
		time.Sleep(2 * time.Second)
		c <- "one"
	}(c1)
	go func(c chan<- string) {
		time.Sleep(time.Second)
		c <- "two"
		//time.Sleep(time.Second)
	}(c2)
	for i := 0; i < 2; i++ {
		// 每次选择一个有值返回的chan
		select {
		case m1 := <-c1:
			fmt.Println("received", m1)
		case m2 := <-c2:
			fmt.Println("received", m2)
		}
	}
}
