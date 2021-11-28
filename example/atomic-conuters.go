package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var counter uint64

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("conter =", counter)
	fmt.Println("conter =", atomic.LoadUint64(&counter))
}
