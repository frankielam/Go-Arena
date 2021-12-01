package main

import "fmt"

func mayPanic() {
	panic("oops...")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered. error: \n", r)
		}
	}()

	mayPanic()

	fmt.Println("after mayPanic")
}