package main

import (
	"fmt"
	"time"
)

func sayHi(name string ) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hi " + name)
	}
}

func main() {
	fmt.Println("Demo 3")
	go sayHi("Frankie")
	fmt.Println("I'm leaving.")
	// 主 goroutine 延迟三秒再结束，sayHi 有足够时间执行，打印成功
	time.Sleep(3 * time.Second)
}
