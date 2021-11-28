package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu sync.Mutex
	counter map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[name]++
}

func main() {
	var wg sync.WaitGroup
	c := Container{
		counter : map[string]int{"a":0, "b":0, "c":0},
	}

	doIncr := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(5)
	go doIncr("a", 1000)
	go doIncr("b", 1000)
	go doIncr("c", 1000)
	go doIncr("c", 1000)
	go doIncr("a", 1000)

	wg.Wait()
	fmt.Println(c.counter)
}