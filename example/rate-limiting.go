package main

import (
	"fmt"
	"time"
)

func main() {
	xx := time.Tick(1 * time.Second)

	for i := 0; i < 5; i++ {
		fmt.Println(i, time.Now())
		<-xx
	}

	limiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		limiter <- time.Now()
	}
	go func() {
		for now := range time.Tick(500 * time.Millisecond) {
			limiter <- now
		}
	}()

	reqs := make(chan int, 5)
	for i := 0; i < 5; i++ {
		reqs <- i
	}
	close(reqs)

	for req := range reqs {
		<-limiter
		fmt.Println(req, time.Now())
	}

	
}