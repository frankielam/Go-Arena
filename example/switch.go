package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1: 
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3: fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: 
		fmt.Println("Today is the weekend")
	default:
		fmt.Println("Today is a weekday")
	}
	
	switch {
	case time.Now().Hour() < 10:
		fmt.Println("Good morning")
	case time.Now().Hour() > 19:
		fmt.Println("Good evening")
	default:
		fmt.Println("Good day")
	}

	typeDetecter := func(i interface{}){
		switch t := i.(type) {
		case int:
			fmt.Println("It's an int")
		case bool:
			fmt.Println("It's a bool")
		default:
			fmt.Printf("It's a %T\n", t)		
		}
	}

	typeDetecter(true)
	typeDetecter(23)
	typeDetecter("hello")
}
