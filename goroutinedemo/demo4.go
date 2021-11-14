package main

import (
	"fmt"
	"sync"
)
	
var wg sync.WaitGroup	// 使用 sync.WaitGroup 

func sayHi(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hi " + name)
	}
	wg.Done()	// wait group 计数器 -1
}

func main() {
	fmt.Println("Demo 4")
	wg.Add(1)	// wait group 计数器 +1
	go sayHi("Frankie")
	wg.Wait()	// wait group 计数器不为 0 时会一直阻塞
	fmt.Println("I'm leaving.")
}
