package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(str string){
		fmt.Println("str", str)
	}("inner")
	time.Sleep(time.Second)

	fmt.Println("Done")
}
