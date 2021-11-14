package main

import "fmt"

func sayHi(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hi " + name)
	}
}

func main() {
	fmt.Println("Demo 2")
	// goroutine 并发调研 sayHi
	go sayHi("Frankie")	
	// 控制台并没有打印 sayHi 的输出
	// 因为主函数 main 对应的 goroutine 已结束，其它的 goroutine 也结束, 实际 sayHi 没有执行到
	fmt.Println("I'm leaving.")
}
