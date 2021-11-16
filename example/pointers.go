package main

import "fmt"

func zeroval(n int) {
	n = 0
} 

func zeroptr(n *int) {
	*n = 0
}

func main() {
	n := 8
	fmt.Println(n)
	fmt.Println(&n)
	zeroval(n)
	fmt.Println(n)
	zeroptr(&n)
	fmt.Println(n)
	fmt.Println(&n)
}
