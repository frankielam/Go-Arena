package main

import "fmt"

func main() {
	msg := make(chan string, 3)
	done := make(chan bool)

	go func(){
		for {
			m, more := <- msg
			if more {
				fmt.Println("received", m, ".")
			} else {
				fmt.Println("received all msgs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 2; i++ {
		msg <- fmt.Sprintf("%d", i)
		fmt.Println("sent", i)
	}
	fmt.Println("sent all")
	close(msg)

	<-done
}
