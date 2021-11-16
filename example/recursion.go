package main

import "fmt"

func recur(n int) int {
	if n == 1 {
		return 1
	}
	return n + recur(n-1)
}

func main() {
	fmt.Println(recur(5))

	var f func(n int) int
	f = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * f(n-1)
	}

	fmt.Println(f(4))
}
