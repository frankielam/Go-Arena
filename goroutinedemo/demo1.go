package main

import "fmt"

func sayHi(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hi " + name)
	}
}

func main() {
	fmt.Println("Demo 1")
	// 同步调用方法
	sayHi("Frankie")
	fmt.Println("I'm leaving.")
}
