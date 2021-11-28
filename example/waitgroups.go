package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(i int) {
	fmt.Println("worker", i, "start")
	time.Sleep(time.Second)
	fmt.Println("worker", i, "end")
}

func worker2(i int, wg *sync.WaitGroup) {
	fmt.Println("worker2", i, "start")
	time.Sleep(time.Second)
	fmt.Println("worker2", i, "end")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i 
		go func() {
			defer wg.Done()		// 延迟执行
			worker(i)
		}()
	}

	wg.Wait()

	fmt.Println("round 2")
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go worker2(j, &wg)
	}

	wg.Wait()
}