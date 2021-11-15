package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(queue chan int) {
	idx := 0
	for {
		queue <- idx
		idx++
		time.Sleep(100 * time.Millisecond)
	}
}

// 消费者
func consumer(queue chan int) {
	for {
		product := <-queue	//channel 中无数据时会阻塞 
		fmt.Println(product)	
	}
}

func main() {
	queue := make(chan int)
	go producer(queue)	//异步调用生产者
	consumer(queue)		//同步调用消费者
}
