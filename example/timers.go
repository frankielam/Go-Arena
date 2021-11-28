package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(1 * time.Second)	

	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()	
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer 2 stoped")
	}

	time.Sleep(2 * time.Second)

}